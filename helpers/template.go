package helpers

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templates []string, name string, data interface{}) error {
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}

	return nil
}

func GetBaseTemplates() []string {
	templates := []string{"views/base/foot.html", "views/base/head.html", "views/base/navbar.html", "views/base.html"}
	return templates
}
