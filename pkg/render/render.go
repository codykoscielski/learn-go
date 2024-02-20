package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate Render template using HTML
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}