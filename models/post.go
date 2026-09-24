package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"uniqueIndex:idx_author_title"`
	Content   string    `json:"content" gorm:"not null"`
	AuthorID  uint      `json:"author_id" gorm:"uniqueIndex:idx_author_title"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}