package provider

import "fmt"

const (
	aisiaProductDocsURL = "https://aisia.fr/docs"
	aisiaAPIDocsURL     = "https://api.aisia.fr/docs"
)

// docLinksForEndpoint — liens markdown vers la doc produit et l'OpenAPI (registry Terraform).
func docLinksForEndpoint(path string) string {
	if path == "" {
		return fmt.Sprintf(
			"\n\n- **Documentation produit** : [aisia.fr/docs](%s)\n- **Référence API OpenAPI** : [api.aisia.fr/docs](%s)",
			aisiaProductDocsURL, aisiaAPIDocsURL,
		)
	}
	return fmt.Sprintf(
		"\n\n- **Endpoint** : `%s`\n- **Documentation produit** : [aisia.fr/docs](%s)\n- **Référence API OpenAPI** : [api.aisia.fr/docs](%s)",
		path, aisiaProductDocsURL, aisiaAPIDocsURL,
	)
}