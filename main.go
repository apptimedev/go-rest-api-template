package main

import (
	"github.com/foruscommunity/go-rest-api-example/pkg/routes"
	"github.com/foruscommunity/go-rest-api-example/pkg/utils/env"
	"log"
	"net/http"
)

func main() {
	name := env.Get("NAME")
	version := env.Get("VERSION")
	port := env.Get("PORT")

	log.Println(name)
	log.Println("Version " + version)
	log.Println("Serving on port " + port)
	log.Println("http://localhost:" + port)

	http.Handle("/", routes.Handlers())

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
