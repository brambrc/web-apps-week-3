package model

import (
	"quoteapp/db"
	"gorm.io/gorm"
)

type Quote struct {
	gorm *gorm.DB
}

func NewQuoteModel(g *gorm.DB) *Quote {
	return &Quote{gorm: g}
}

func (q *Quote) Create(qoute any) error {
	return q.gorm.Create(qoute).Error
}

func (q *Quote) Get() (db.Quotes, error) {
	result := db.Quotes{}
	err := q.gorm.Raw("select * from quotes order by random() limit 1").Scan(&result).Error
	return result, err
}

func (q *Quote) Count() (int64, error) {
	var result int64
	err := q.gorm.Model(&db.Quotes{}).Count(&result).Error
	return result, err
}
