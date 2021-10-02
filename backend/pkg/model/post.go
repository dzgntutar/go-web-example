package model

import "time"

//post object
type Post struct {
	Id       int64
	Title    string
	Body     string
	CreateAt time.Time
	UpdateAt time.Time
	IsShare  bool
	IsActive bool
}
