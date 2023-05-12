package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nach9/go-todolist/pkg/activities"
	"github.com/nach9/go-todolist/pkg/common/db"
	"github.com/nach9/go-todolist/pkg/todos"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	gin.SetMode(gin.ReleaseMode)

	appPort := "3030"

	dbName := os.Getenv("MYSQL_DBNAME")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	r := gin.Default()
	db := db.Init(dbUrl)

	r.Use(cors.Default())

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
