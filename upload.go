package main

import (
	entityfileuploader "github.com/joegasewicz/entity-file-uploader"
	"log"
)

func NewFilesManager(table, allowedTypes string) *entityfileuploader.FileManager {
	fileUpload := entityfileuploader.FileUpload{
		UploadDir:   "uploads",
		MaxFileSize: 10,
		FileTypes:   []string{"png", "jpeg", "jpg"}, // TODO
		URL:         "http://localhost:4444",
	}
	fileUploader, err := fileUpload.Init(table)
	if err != nil {
		log.Printf("Error creating file uploarder: %e", err)
	}
	return fileUploader
}
