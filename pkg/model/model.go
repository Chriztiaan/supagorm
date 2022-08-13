package model

import (
	"supagorm/pkg"

	"gorm.io/gorm"
)

type Post struct {
	pkg.Model
	Title    string    `json:"title"`
	Content  string    `json:"content" gorm:"type:string; size:256; not null"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	pkg.Model
	Content string `json:"content" gorm:"type:string; size:256; not null"`
	PostID  string `json:"postId"`
	Post    Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func LoadModel(db *gorm.DB) {
	db.AutoMigrate(&Post{}, &Comment{})

	// Ensures Post model, and by extension Comment model, have TS representation
	pkg.GenerateTypeScriptModel(Post{})
}
