package model

import "time"

type Category struct {
	Id          int64
	Title       string
	Description string
	CreateDate  time.Time
	UpdateDate  time.Time
}

type CategoryDto struct {
	Id          int64
	Title       string
	Description string
}

func (c CategoryDto) ToCategory() Category {
	return Category{
		Id:          c.Id,
		Title:       c.Title,
		Description: c.Description,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
	}
}
