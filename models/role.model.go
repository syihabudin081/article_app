// Role model for user roles with unique role names
package models

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;unique"` // Examples: "admin", "author", "user"
}

// Define available roles as constants for easy reference
const (
	AdminRoleName  = "admin"
	AuthorRoleName = "author"
	UserRoleName   = "user"
	GuestRoleName  = "guest"
)
