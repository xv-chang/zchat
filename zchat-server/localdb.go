package main

import (
	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var LocalDb *gorm.DB

func init() {

	LocalDb, _ = gorm.Open(sqlite.Open("zchat.db"), &gorm.Config{})
	// 迁移 schema
	LocalDb.AutoMigrate(&ChatMessage{})

}
