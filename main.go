package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nach9/go-todolist/pkg/activities"
	"github.com/nach9/go-todolist/pkg/common/db"
	"github.com/nach9/go-todolist/pkg/todos"
	"github.com/spf13/viper"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	gin.SetMode(gin.ReleaseMode)

	appPort := "3030"

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

	srv := &http.Server{
		Addr:    ":" + appPort,
		Handler: r,
	}

	log.Println("Listening and serving HTTP on PORT", appPort)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
