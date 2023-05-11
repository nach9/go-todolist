package db

import (
	"log"

	entityActivity "github.com/nach9/go-todolist/pkg/activities/entity"
	entityTodo "github.com/nach9/go-todolist/pkg/todos/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&entityActivity.Activity{},
		&entityTodo.Todo{},
	)

	return db
}
