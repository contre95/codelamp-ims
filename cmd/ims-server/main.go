package main

import (
	"codelamp-ims/pkg/domain/contacts"
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

	contactRepo := sql.NewSQLRepository(db)
	contact := contacts.NewService(contactRepo)

	// Migrate the schema
	//db.AutoMigrate(&Product{})

	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42

	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//// DeleteContact - delete product
	////db.DeleteContact(&product, 1)
}
