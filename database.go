package main

import (
	pg_conf "github.com/joegasewicz/pg-conf"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	//

	pg_conf.Update()
}
