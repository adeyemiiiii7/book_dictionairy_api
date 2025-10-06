package models

import (
	"time"
	"gorm.io/gorm"
)

// Book represents a book in the dictionary
type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;size:255"`
	Author    string         `json:"author" gorm:"not null;size:255"`
	Quantity  int            `json:"quantity" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName specifies the table name for GORM
func (Book) TableName() string {
	return "books"
}
