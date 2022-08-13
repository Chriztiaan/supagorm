package main

import (
	"supagorm/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://postgres:[YOUR-PASSWORD]@db.[HOST].supabase.co:5432/postgres"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	model.LoadModel(db)
}
