package main

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	Name       string `gorm:"not null"`
	FileName   string
	Data       string
	EntityName string
}
