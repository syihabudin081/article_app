package models

import "github.com/google/uuid"

type Article struct {
	BaseModel
	Title    string    `json:"title" gorm:"not null"`
	Category string    `json:"category" gorm:"not null"`
	Content  string    `json:"content" gorm:"not null"`
	AuthorID uuid.UUID `json:"author_id" gorm:"not null"` // Foreign key to User
	Author   User      `json:"author,omitempty" gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:SET NULL;"`
	Status   string    `json:"status" gorm:"not null;default:'draft'"` // Options: draft, published, archived
}
