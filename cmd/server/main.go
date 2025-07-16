package main

import (
	"log"
	"net/http"

	appHttp "ojeg/delivery/http"
	"ojeg/delivery/http/registry"
	"ojeg/pkg/bootstrap"
)

func main() {
	// Initialize DB
	database := bootstrap.InitDB()

	// Register handlers
	handlers := registry.NewHandlerRegistry(database)

	// Setup router
	router := appHttp.NewRouter(handlers, database)

	log.Println("🚀 Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
