package services

import (
	"go-web-toko/models"
	"go-web-toko/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	data, err := repository.GetAllCategory()
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.SecureJSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CategoryCreate(c *gin.Context) (*models.Category, error) {
	var form models.CategoryForm
	err := c.BindJSON(&form)
	if err != nil {
		return nil, err
	}
	category := models.Category{
		Name: form.Name,
	}

	data, err := repository.Create(category)
	return &data, err
}

func CategoryEdit(w http.ResponseWriter, r *http.Request) {

}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {

}
