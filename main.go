package main

import "github.com.br/NicolasDeveloper/catalog-api/api"

func main() {

	api := api.NewApp()
	api.Initialize().StartupDatabase().ConfigEndpoints().Run()
}
