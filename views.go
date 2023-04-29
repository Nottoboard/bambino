package main

import (
	"encoding/json"
	entityfileuploader "github.com/joegasewicz/entity-file-uploader"
	"github.com/joegasewicz/gomek"
	"net/http"
)

type FileView struct{}
type UserView struct{}
type HealthView struct{}

// file
func (f *FileView) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var data = struct {
		Name string
	}{
		Name: "test1",
	}
	gomek.JSON(w, data, http.StatusOK)
}

func (f *FileView) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var options OptionSchema
	optionsStr, err := gomek.GetParams(r, "options")
	if err != nil || optionsStr == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return // TODO return json
	}
	err = json.Unmarshal([]byte(optionsStr[0]), &options)
	for _, fileName := range options.Files {
		_, err := entityfileuploader.GetFileName(r, fileName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return // return json
		}
	}
}

func (f *FileView) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (f *FileView) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}

// health
func (h *HealthView) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	templateData := make(gomek.Data)
	templateData["health"] = map[string]string{
		"status": "OK",
	}
	gomek.JSON(w, templateData, http.StatusOK)
}
func (h *HealthView) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *HealthView) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *HealthView) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}

// user
func (u *UserView) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *UserView) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *UserView) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *UserView) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
