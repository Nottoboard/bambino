package main

import (
	"github.com/joegasewicz/gomek"
	"net/http/httptest"
	"testing"
)

func TestFileView_Post(t *testing.T) {
	request := httptest.NewRequest("GET", "/files", nil)
	respRecorder := httptest.NewRecorder()
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
		},
	}
	app := gomek.New(c)
	app.Route("/files").Methods("POST").Resource(&FileView{})
}
