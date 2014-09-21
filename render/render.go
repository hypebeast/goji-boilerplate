package render

import (
	"encoding/json"
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

func RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	var result []byte
	var err error

	result, err = json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// json rendered fine, write out the result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

func GetBaseTemplates() []string {
	templates := []string{"views/base/foot.html", "views/base/head.html", "views/base/navbar.html", "views/base.html"}
	return templates
}
