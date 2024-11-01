package models

import "github.com/google/uuid"

type Comment struct {
	BaseModel
	ArticleID uuid.UUID `json:"article_id" gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with CASCADE delete
	Content   string    `json:"content" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	Article   Article   `json:"article,omitempty" gorm:"foreignKey:ArticleID;references:ID"`
}
