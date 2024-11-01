package models

// User model for application users
type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	RoleID   int    `json:"role_id" gorm:"not null"` // Foreign key to Role
	Role     Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}
