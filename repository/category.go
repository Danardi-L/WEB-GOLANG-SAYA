package repository

import (
	"go-web-toko/config"
	"go-web-toko/models"
	"time"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func Create(category models.Category) (models.Category, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	result, err := config.DB.Exec(`INSERT INTO categories 
	(name,created_at,updated_at) 
		VALUES
	(?,?,?)`, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return models.Category{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Category{}, err
	}

	category.Id = id
	return category, nil
}

func GetOne(id int64) (models.Category, error) {
	var category models.Category
	err := config.DB.QueryRow("SELECT * FROM categories WHERE id = ?", id).
		Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	return category, err
}
