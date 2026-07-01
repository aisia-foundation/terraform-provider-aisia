#!/usr/bin/env bash
# =============================================================================
# publish-hcp.sh — Publie une version du provider PRIVÉ sur le registry HCP
# (app.terraform.io) via l'API v2. Un provider privé NE s'ingère PAS par tag git
# (contrairement aux modules) : chaque version doit être enregistrée par l'API.
#
# Flux (doc HashiCorp « Publishing Private Providers ») :
#   1. POST create version  → renvoie les liens d'upload shasums + shasums.sig
#   2. PUT  upload SHA256SUMS                (lien pré-signé, sans header auth)
#   3. PUT  upload SHA256SUMS.sig
#   4. pour chaque plateforme : POST create platform → lien d'upload binaire
#   5. PUT  upload le .zip de la plateforme
#
# Requiert : TF_API_TOKEN (SECRET GitHub — jamais en clair), GPG_FINGERPRINT,
#            VERSION (sans 'v'), artefacts GoReleaser dans dist/. Outils : curl, jq.
# Idempotent-friendly : si la version existe déjà (409), on continue les plateformes.
# Si TF_API_TOKEN est absent → skip propre (exit 0) pour ne pas casser la release.
# =============================================================================
set -euo pipefail

ORG="${HCP_ORG:-AISIA}"
NS="${HCP_NAMESPACE:-AISIA}"
NAME="${HCP_PROVIDER_NAME:-aisia}"
API="https://app.terraform.io/api/v2"
DIST="${DIST_DIR:-dist}"
VERSION="${VERSION:?VERSION requis (ex. 6.9.63, sans 'v')}"

if [[ -z "${TF_API_TOKEN:-}" ]]; then
  echo "::warning::TF_API_TOKEN absent → publication HCP du provider SKIPPÉE."
  echo "Ajoute un secret GitHub Actions 'TF_API_TOKEN' (HCP team API token) pour activer."
  exit 0
fi

# key-id HCP = ID GPG (16 hex de fin du fingerprint), doit matcher la clé publique
# enregistrée dans HCP (Account Settings > GPG keys).
KEYID="${GPG_KEY_ID:-}"
if [[ -z "$KEYID" && -n "${GPG_FINGERPRINT:-}" ]]; then
  KEYID="${GPG_FINGERPRINT: -16}"
fi
: "${KEYID:?GPG_KEY_ID ou GPG_FINGERPRINT requis}"

SHASUMS="$DIST/terraform-provider-${NAME}_${VERSION}_SHA256SUMS"
SHASIG="${SHASUMS}.sig"
[[ -f "$SHASUMS" ]] || { echo "::error::introuvable: $SHASUMS"; exit 1; }
[[ -f "$SHASIG"  ]] || { echo "::error::introuvable: $SHASIG"; exit 1; }

AUTH=(-H "Authorization: Bearer $TF_API_TOKEN" -H "Content-Type: application/vnd.api+json")

echo ">> [1/5] Création de la version $VERSION (key-id $KEYID)"
ver_payload=$(printf '{"data":{"type":"registry-provider-versions","attributes":{"version":"%s","key-id":"%s","protocols":["6.0"]}}}' "$VERSION" "$KEYID")
resp=$(curl -s -w '\n%{http_code}' "${AUTH[@]}" -X POST \
  "$API/organizations/$ORG/registry-providers/private/$NS/$NAME/versions" -d "$ver_payload")
code=$(tail -n1 <<<"$resp"); body=$(sed '$d' <<<"$resp")
if [[ "$code" == "201" ]]; then
  shasums_url=$(jq -r '.data.links."shasums-upload"' <<<"$body")
  shasig_url=$(jq -r '.data.links."shasums-sig-upload"' <<<"$body")
  echo ">> [2/5] Upload SHA256SUMS"; curl -sf -X PUT -T "$SHASUMS" "$shasums_url"
  echo ">> [3/5] Upload SHA256SUMS.sig"; curl -sf -X PUT -T "$SHASIG" "$shasig_url"
elif [[ "$code" == "422" || "$code" == "409" ]]; then
  echo "::warning::version $VERSION déjà présente (HTTP $code) → on (ré)upload les plateformes manquantes"
else
  echo "::error::create version a échoué (HTTP $code): $body"; exit 1
fi

echo ">> [4-5/5] Upload des plateformes"
while read -r shasum filename; do
  [[ "$filename" == *_SHA256SUMS* ]] && continue
  [[ "$filename" == *.zip ]] || continue
  base="${filename%.zip}"; rest="${base#terraform-provider-${NAME}_${VERSION}_}"
  os="${rest%%_*}"; arch="${rest#*_}"
  echo "   - ${os}/${arch}  ($filename)"
  plat_payload=$(printf '{"data":{"type":"registry-provider-version-platforms","attributes":{"os":"%s","arch":"%s","shasum":"%s","filename":"%s"}}}' "$os" "$arch" "$shasum" "$filename")
  presp=$(curl -s -w '\n%{http_code}' "${AUTH[@]}" -X POST \
    "$API/organizations/$ORG/registry-providers/private/$NS/$NAME/versions/$VERSION/platforms" -d "$plat_payload")
  pcode=$(tail -n1 <<<"$presp"); pbody=$(sed '$d' <<<"$presp")
  if [[ "$pcode" != "201" ]]; then
    echo "::warning::plateforme $os/$arch non créée (HTTP $pcode) — peut déjà exister, skip"; continue
  fi
  bin_url=$(jq -r '.data.links."provider-binary-upload"' <<<"$pbody")
  curl -sf -X PUT -T "$DIST/$filename" "$bin_url"
done < "$SHASUMS"

echo ">> ✅ Provider ${NAME} ${VERSION} publié sur HCP (org ${ORG})."
