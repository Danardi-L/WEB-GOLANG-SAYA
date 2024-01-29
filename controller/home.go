package controller

import (
	"fmt"
	"go-web-toko/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "Main website",
		"categories": categories,
	})
}
