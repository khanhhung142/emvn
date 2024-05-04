package main

import (
	"emvn/cmd/server"
	_ "emvn/docs"
)

//	@title			EMVN API
//	@version		1.0
//	@description	Hung.Phan apply for junior Golang developer position at EMVN

// @securityDefinitions.apikey BearerAuth
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	server.StartServer()
}
