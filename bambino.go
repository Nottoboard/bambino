package main

import (
	"github.com/joegasewicz/bambino/views"
	"github.com/joegasewicz/gomek"
	"log"
)

var FileUpload = NewFileUpload()

func main() {
	// config
	config := NewConfig()
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
		},
	}
	// app
	app := gomek.New(c)
	// views
	app.Route("/health").Methods("GET").Resource(&views.Health{})
	app.Route("/files").Methods("GET", "POST").Resource(&views.File{})
	// middleware
	app.Use(gomek.Logging)
	app.Use(gomek.CORS)
	app.SetHost(config.HOST)
	app.Listen(config.PORT)
	if err := app.Start(); err != nil {
		log.Printf("Error running Gomek: %e", err)
	}
}
