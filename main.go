package main

import (
	"go-web-toko/config"
	"go-web-toko/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 1.homepagenya
	r.GET("/", controller.Welcome)

	// // 2. category
	// http.HandleFunc("/category", categorycontroller.indexCategory)
	// r.POST("/category/create", services.CategoryCreate)
	// r.PATCH("/category/edit", services.CategoryEdit)
	// r.DELETE("/category/delete", services.CategoryDelete)

	log.Println("Server started on: http://localhost:8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
