package pkg

import (
	"time"

	"gorm.io/gorm"
)

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model or you may build your own model without it
//    type User struct {
//      gorm.Model
//    }

type Model struct {
	ID        string         `json:"id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
