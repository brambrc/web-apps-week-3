package model

import (
	"quoteapp/db"
	"gorm.io/gorm"
)

type Users struct {
	gorm *gorm.DB
}


func NewUsersModel(g *gorm.DB) *Users {
	return &Users{gorm: g}
}

func (q *Users) Create(users any) error {
	return q.gorm.Create(users).Error
}

func (q *Users) FindAll() ([]db.Users, error) {
	result := []db.Users{}
	err := q.gorm.Find(&result).Error
	return result, err
}

func (q *Users) FindByID(id string) (db.Users, error) {
	result := db.Users{}
	err := q.gorm.Where("id = ?", id).First(&result).Error
	return result, err
}

func (q *Users) Update(id string, users any) error {
	return q.gorm.Where("id = ?", id).Updates(users).Error
}

func (q *Users) Delete(id string) error {
	return q.gorm.Where("id = ?", id).Delete(&db.Users{}).Error
}

func (q *Users) FindByEmail(email string) (db.Users, error) {
	result := db.Users{}
	err := q.gorm.Where("email = ?", email).First(&result).Error
	return result, err
}