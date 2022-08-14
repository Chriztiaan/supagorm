package model

import (
	"supagorm/pkg"

	"gorm.io/gorm"
)

type Post struct {
	pkg.Model
	Title    string `json:"title"`
	Content  string `json:"content" gorm:"type:string; size:256; not null"`
	Comments []Comment
	AuthorId string `json:"authorId"`
	Author   Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Comment struct {
	pkg.Model
	Content string `json:"content" gorm:"type:string; size:256; not null"`
	PostID  string `json:"postId"`
	Post    Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Author struct {
	pkg.Model
	Name   string `json:"name" gorm:"type:string; size:256; not null"`
	UserId string `json:"userId"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Because Supabase already provides a Users table, we want to latch onto the existing table which is in the "auth" schema
// We don't want to migrate this table, any fields defined here are for the TS model
// Any
type User struct {
	ID    string `json:"id" gorm:"primarykey;type:uuid"`
	Email string `json:"email" gorm:"-:all"`
}

func (u User) TableName() string {
	return "auth.users"
}

func LoadModel(db *gorm.DB) {
	// Ensure tables exist (Users table's schema is managed by Supabase)
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Post{}, &Comment{})

	// Ensure models have TS representation
	pkg.GenerateTypeScriptModel(Post{})
	pkg.GenerateTypeScriptModel(Comment{})
	pkg.GenerateTypeScriptModel(Author{})
	pkg.GenerateTypeScriptModel(User{})

	// The start of access control.
	db.Exec("ALTER TABLE authors ENABLE ROW LEVEL SECURITY;")
}
