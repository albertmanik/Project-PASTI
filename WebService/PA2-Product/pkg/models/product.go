package models

import (
	"github.com/jinzhu/gorm"
	"PA2-Product/pkg/config"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	// ID				uint   `gorm:"primary_key:auto_increment" json:"id"`
	User_id       	uint   `json:"user_id"`
	Name     		string `json:"name"`
	Category_id   	uint   `json:"category_id"`
	Toko_id 		uint   `json:"toko_id"`
	Harga    		uint   `json:"harga"`
	Kota     		string `json:"kota"`
	Deskripsi     	string `json:"deskripsi"`
	No_hp    		string `json:"no_hp"`
	Gambar     		string `json:"gambar"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

// fungsi unutuk mencreata user baru
func (b *Product) CreateProduct() *Product {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//fungsi untuk mengambil semua data yang terdapat pada table
func GetAllProduct() []Product {
	var Product []Product
	db.Find(&Product)
	return Product
}

// fungsi untuk mengambil data table sesuai dengan id yang di request
func GetProductbyId(id int64) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("id=?", id).Find(&getProduct)
	return &getProduct, db
}

// fungsi untuk menghapus data table sesuai dengan id yang di request
func DeleteProduct(id int64) Product {
	var Product Product
	db.Where("id=?", id).Delete(Product)
	return Product
}
