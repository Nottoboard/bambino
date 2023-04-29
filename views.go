package main

import (
	"encoding/json"
	entityfileuploader "github.com/joegasewicz/entity-file-uploader"
	"github.com/joegasewicz/gomek"
	"net/http"
)

type File struct{}
type User struct{}
type Health struct{}

// file
func (f *File) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var data = struct {
		Name string
	}{
		Name: "test1",
	}
	gomek.JSON(w, data, http.StatusOK)
}

func (f *File) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
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

func (f *File) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (f *File) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}

// health
func (h *Health) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	templateData := make(gomek.Data)
	templateData["health"] = map[string]string{
		"status": "OK",
	}
	gomek.JSON(w, templateData, http.StatusOK)
}
func (h *Health) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *Health) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *Health) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}

// user
func (u *User) Get(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *User) Post(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *User) Put(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (u *User) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}
