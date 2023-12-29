package models

import (
	"github.com/jinzhu/gorm"
	"PA2-Checkout/pkg/config"
)

var db *gorm.DB

type Checkout struct {
	gorm.Model
	// ID				uint   `gorm:"primary_key:auto_increment" json:"id"`
	User_id       	uint   `json:"user_id"`
	Order_number 	string `json:"order_number"`
	Name     		string `json:"name"`
	No_hp     		string `json:"no_hp"`
	Alamat   		string `json:"alamat"`
	Ucapan 			string `json:"ucapan"`
	Image    		string `json:"image"`
	Total     		uint   `json:"total"`
	Status     		string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Checkout{})
}

// fungsi unutuk mencreata user baru
func (b *Checkout) CreateCheckout() *Checkout {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//fungsi untuk mengambil semua data yang terdapat pada table
func GetAllCheckout() []Checkout {
	var Checkout []Checkout
	db.Find(&Checkout)
	return Checkout
}

// fungsi untuk mengambil data table sesuai dengan id yang di request
func GetCheckoutbyId(id int64) (*Checkout, *gorm.DB) {
	var getCheckout Checkout
	db := db.Where("id=?", id).Find(&getCheckout)
	return &getCheckout, db
}

// fungsi untuk menghapus data table sesuai dengan id yang di request
func DeleteCheckout(id int64) Checkout {
	var Checkout Checkout
	db.Where("id=?", id).Delete(Checkout)
	return Checkout
}
