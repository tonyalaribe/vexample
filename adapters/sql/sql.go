package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/tonyalaribe/vexample"
)

type Provider struct {
	db *gorm.DB
}

func New() (*Provider, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&vexample.Product{})

	// Create
	db.Create(&vexample.Product{Title: "product title"})

	// Read
	var product vexample.Product
	db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

	return &Provider{db}, nil
}

func (p *Provider) Cleanup() {
	p.db.Close()
}
