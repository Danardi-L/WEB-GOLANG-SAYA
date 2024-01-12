package main

import (
	"go-web-toko/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
