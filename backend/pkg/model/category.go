package model

import "time"

type Category struct {
	Id          int64
	Title       string
	Description string
	CreateDate  time.Time
	UpdateDate  time.Time
}
