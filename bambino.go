package main

import (
	"github.com/joegasewicz/gomek"
	"log"
)

var (
	AppConfig = NewConfig()
	DB        = CreateDB()
)

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
	app.Route("/health").Methods("GET").Resource(&HealthView{})
	app.Route("/files").Methods("GET", "POST").Resource(&FileView{})
	// middleware
	app.Use(gomek.Logging)
	app.Use(gomek.CORS)
	app.SetHost(config.HOST)
	app.Listen(config.PORT)
	if err := app.Start(); err != nil {
		log.Printf("Error running Gomek: %e", err)
	}
}
