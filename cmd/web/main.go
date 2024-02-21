package main

import (
	"fmt"
	"github.com/codykoscielski/learn-go/pkg/config"
	"github.com/codykoscielski/learn-go/pkg/handlers"
	"github.com/codykoscielski/learn-go/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":6969"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Start app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
