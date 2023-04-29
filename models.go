package main

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Data       string
	EntityName string
}
