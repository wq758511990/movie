package models

import "github.com/jinzhu/gorm"

type Movie struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	ImgUrl       string `json:"imgUrl"`
	Introduction string `json:"introduction""`
	Rating       string `json:"rating" gorm:"size:10" `
	Country      string `json:"country" gorm:"size:20"`
	Actors       string `json:"actors" gorm:"size:100"`
	Director     string `json:"director" gorm:"size:50"`
	CategoryId   string `json:"categoryId"`
}

type Category struct {
	gorm.Model
	Name             string `json:"name""`
	ParentCategoryID int64  `json:"parentCategoryId"`
}
