package main

import (
	"fmt"
	"gingonic-test/database"
	"gingonic-test/job"
	"gingonic-test/product"
	"gingonic-test/router/admin"
	v1API "gingonic-test/router/api/v1"
	"gingonic-test/router/gorutine"
	"gingonic-test/router/scheduler"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func initDatabase() {
	var err error
	database.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	//Migrate the schema
	database.Db.AutoMigrate(&product.Product{})
	fmt.Println("Database Migrated")
}

func setRoutes() *gin.Engine {
	router := gin.Default()
	v1API.SetApiRouter(router)
	admin.SetAdminRouter(router)
	gorutine.SetGoRoutineRouter(router)
	scheduler.SetSchedulerRouter(router)
	return router
}

func setServerConfig(router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}

func main() {

	router := setRoutes()
	initDatabase()

	server := setServerConfig(router)

	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@every 5s", job.Scheduler{})

	server.ListenAndServe()

}
