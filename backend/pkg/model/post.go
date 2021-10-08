package model

import "time"

//post object
type Post struct {
	Id         int64
	Title      string
	Body       string
	CreateDate time.Time
	UpdateDate time.Time
	IsShare    bool
	IsActive   bool
	CategoryId int64
}
