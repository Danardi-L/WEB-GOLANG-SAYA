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

func CategoryList(c *gin.Context) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		fmt.Println(err.Error())
	}

	list_produk := "<ul>"

	for _, category := range categories {
		fmt.Println(category)
		list_produk += "<li>" + category.Name + "</li>"
	}

	list_produk += "</ul>"

	c.HTML(http.StatusOK, "category.html", gin.H{
		"title":      "Category",
		"categories": categories,
		"Items":      list_produk,
	})
}

func ProductList(c *gin.Context) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "product.html", gin.H{
		"title":      "Product",
		"categories": categories,
	})
}

func About(c *gin.Context) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title":      "About",
		"categories": categories,
	})
}
