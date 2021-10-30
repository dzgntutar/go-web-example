package model

import (
	"time"
)

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

type PostDto struct {
	Id         int64
	Title      string
	Body       string
	CreateDate time.Time
	UpdateDate time.Time
	CategoryId int64 `json:"CategoryId,string,omitempty"`
}

func (postDto PostDto) ToPost() Post {
	return Post{
		Title:      postDto.Title,
		Body:       postDto.Body,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		IsShare:    false,
		IsActive:   true,
		CategoryId: postDto.CategoryId,
	}
}
func (post Post) ToPostDto() PostDto {
	return PostDto{
		Id:         post.Id,
		Title:      post.Title,
		Body:       post.Body,
		CreateDate: post.CreateDate,
		UpdateDate: post.UpdateDate,
	}
}
