// terraform-provider-aisia — point d'entrée du plugin.
//
// Provider Terraform pour gérer les ressources AISIA (organisations, clés
// providers par org, …) déclarativement via l'API AISIA (api.aisia.fr).
package main

import (
	"context"
	"flag"
	"log"

	"github.com/AISIA-fr/terraform-provider-aisia/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Renseigné au build (GoReleaser : -ldflags "-X main.version=...").
var version = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// Adresse registry : <namespace>/<type> → AISIA-fr/aisia.
		Address: "registry.terraform.io/AISIA-fr/aisia",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
