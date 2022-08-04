package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"leave.gg/pkg/employee"

	colorable "github.com/mattn/go-colorable"
)

func main() {
	// Fixing windows colours
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()

	// Init gin and gorm
	router := gin.Default()
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	// Load controllers
	employee.NewEmployeeController(router, db)
	prepareRouter(router)

	// Serve
	log.Fatal(router.Run(":8664"))

}

func prepareRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Ping"})
	})
}
