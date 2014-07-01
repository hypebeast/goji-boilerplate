package controllers

import (
	"github.com/hypebeast/goji-boilerplate/helpers"

	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	templates := helpers.GetBaseTemplates()
	templates = append(templates, "views/index.html")
	err := helpers.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}