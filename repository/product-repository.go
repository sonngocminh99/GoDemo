package repository

import (
	"log"

	"github.com/sonngocminh99/GoDemo/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(b entity.Product) entity.Product
	UpdateProduct(b entity.Product) entity.Product
	DeleteProduct(b entity.Product)
	AllProducts() []entity.Product
	FindProductById(id int) entity.Product
}

type ProductConnection struct {
	ProductConnection *gorm.DB
}

func NewProductConnection(dbConn *gorm.DB) *ProductConnection {
	return &ProductConnection{
		ProductConnection: dbConn,
	}

}

func (db *ProductConnection) InsertProduct(b entity.Product) entity.Product {
	db.ProductConnection.Save(&b)
	return b
}

func (db *ProductConnection) UpdateProduct(b entity.Product) entity.Product {
	db.ProductConnection.Save(&b)
	return b
}

func (db *ProductConnection) DeleteProduct(b entity.Product) {
	db.ProductConnection.Delete(&b)
}

func (db *ProductConnection) AllProducts() []entity.Product {
	var product []entity.Product
	//db.ProductConnection.Find(entity.Tabler.TableName(product))
	return product
}

func (db *ProductConnection) FindProductById(id int) entity.Product {
	var product entity.Product
	db.ProductConnection.Where("id = ?", id).First(entity.Tabler.TableName(product))
	log.Println(db.ProductConnection.Table("ecommerce_product"))
	return product
}
