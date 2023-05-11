package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nach9/go-todolist/pkg/activities"
	"github.com/nach9/go-todolist/pkg/common/db"
	"github.com/nach9/go-todolist/pkg/todos"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	dbName := viper.Get("MYSQL_DBNAME").(string)
	dbUser := viper.Get("MYSQL_USER").(string)
	dbPass := viper.Get("MYSQL_PASSWORD").(string)
	dbHost := viper.Get("MYSQL_HOST").(string)
	dbPort := viper.Get("MYSQL_PORT").(string)

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	r := gin.Default()
	db := db.Init(dbUrl)

	activities.RegisterRoutes(r, db)
	todos.RegisterRoutes(r, db)

	r.Run(":3030")
}
