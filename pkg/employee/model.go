package employee

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name" gorm:"type:string; size:20; unique; not null" binding:"required"`
	// ManagerID uint "This is how you set a relationship"
	// Manager   Manager
}
