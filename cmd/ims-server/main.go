package main

import (
	"codelamp-ims/pkg/gateways/storage/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//dsn := "agusotto:ottoawesome@tcp(127.0.0.1:3308)/mydbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("ims.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sql := sql.NewStorage(db)
	sql.Migrate()
}
