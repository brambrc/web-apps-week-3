package db

import "gorm.io/gorm"

type Quotes struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
	gorm.Model
}
