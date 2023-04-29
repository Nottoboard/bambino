package main

import (
	"encoding/json"
	entityfileuploader "github.com/joegasewicz/entity-file-uploader"
	"github.com/joegasewicz/gomek"
	"log"
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
	if err != nil {
		log.Println(err.Error())
		gomek.JSON(w, nil, http.StatusBadRequest)
		return
	}
	for _, optionsfileName := range options.Files {
		fileName, err := entityfileuploader.GetFileName(r, optionsfileName)
		if err != nil {
			log.Println(err.Error())
			gomek.JSON(w, nil, http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(options.Data)
		if err != nil {
			log.Println(err.Error())
			gomek.JSON(w, nil, http.StatusBadRequest)
			return
		}
		fileModel := FileModel{
			Name:       optionsfileName,
			Data:       string(data),
			EntityName: options.EntityName,
		}
		result := DB.Create(&fileModel)
		if result.Error != nil {
			log.Printf("error saving file %s\n", result.Error.Error())
			gomek.JSON(w, nil, http.StatusInternalServerError)
			return
		}
		if result.RowsAffected == 0 {
			log.Printf("unable to save file with name: %s", fileName)
		}
		_, err = FileUploader.Upload(w, r, fileModel.ID, optionsfileName)
		if err != nil {
			log.Println(err)
			log.Printf("unable to store file on server with name: %s", fileName)
			gomek.JSON(w, nil, http.StatusInternalServerError)
			return
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
