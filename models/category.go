package models

import "time"

type Category struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}
