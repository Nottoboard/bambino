package main

import entityfileuploader "github.com/joegasewicz/entity-file-uploader"

var FileUpload = entityfileuploader.FileUpload{
	UploadDir:   "uploads",
	MaxFileSize: 10,
	FileTypes:   []string{"png", "jpeg", "jpg"},
	URL:         "http://localhost:",
}
