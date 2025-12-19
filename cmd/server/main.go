package main

import (
	"net/http"

	"github.com/flp-fernandes/9-APIS/internal/entity"
	"github.com/flp-fernandes/9-APIS/internal/infra/database"
	"github.com/flp-fernandes/9-APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// _, err := configs.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", ProductHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
