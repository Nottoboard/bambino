package main

import (
	"context"
	"github.com/joegasewicz/gomek"
	"log"
	"net/http"
)

var (
	AppConfig = NewConfig()
	DB        = CreateDB()
)

func main() {
	var err error
	// config
	config := NewConfig()
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
		},
	}
	// migrations
	err = DB.AutoMigrate(
		&FileModel{},
	)
	if err != nil {
		log.Panic(err)
	}
	// app
	app := gomek.New(c)
	// views
	app.Route("/health").Methods("GET").Resource(&HealthView{})
	app.Route("/files").Methods("GET", "POST").Resource(&FileView{})
	// Serve static content
	uploadFiles := http.FileServer(http.Dir("uploads"))
	app.Handle("/uploads/", http.StripPrefix("/uploads/", uploadFiles))
	// middleware
	app.Use(gomek.Logging)
	app.Use(gomek.CORS)
	app.Use(gomek.Authorize(WHITE_LIST, func(r *http.Request) (bool, context.Context) {
		return true, context.Background() // TODO
	}))
	app.SetHost(config.HOST)
	app.Listen(config.PORT)
	if err := app.Start(); err != nil {
		log.Printf("Error running Gomek: %e", err)
		panic(err)
	}
}
