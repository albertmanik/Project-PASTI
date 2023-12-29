package models

import (
	"github.com/jinzhu/gorm"
	"PA2-Toko/pkg/config"
)

var db *gorm.DB

type Toko struct {
	gorm.Model
	// ID				uint   `gorm:"primary_key:auto_increment" json:"id"`
	User_id       	uint   `json:"user_id"`
	Nama_toko     	string `json:"nama_toko"`
	Alamat   		string   `json:"alamat"`
	Status 			string   `json:"status"`
	Total_produk    uint   `json:"total_produk"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Toko{})
}

// fungsi unutuk mencreata user baru
func (b *Toko) CreateToko() *Toko {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//fungsi untuk mengambil semua data yang terdapat pada table
func GetAllToko() []Toko {
	var Toko []Toko
	db.Find(&Toko)
	return Toko
}

// fungsi untuk mengambil data table sesuai dengan id yang di request
func GetTokobyId(id int64) (*Toko, *gorm.DB) {
	var getToko Toko
	db := db.Where("id=?", id).Find(&getToko)
	return &getToko, db
}

// fungsi untuk menghapus data table sesuai dengan id yang di request
func DeleteToko(id int64) Toko {
	var Toko Toko
	db.Where("id=?", id).Delete(Toko)
	return Toko
}
