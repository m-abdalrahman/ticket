package models

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	db *gorm.DB
}

func NewModel(db *gorm.DB) *Model {
	return &Model{db: db}
}
