package main

import (
	"github.com/hypebeast/goji-boilerplate/routes"

	"github.com/zenazn/goji"
	"net/http"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	// Add routes
	routes.Include()

	// Run Goji
	goji.Serve()
}
