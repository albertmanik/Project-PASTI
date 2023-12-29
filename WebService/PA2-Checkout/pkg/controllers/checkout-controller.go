package controllers

import ( // mengimport library yang dibutuhkan
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"PA2-Checkout/pkg/models"
	"PA2-Checkout/pkg/utils"
)

var NewCheckout models.Checkout

// fungsi yang digunakan memanggil seluruh data yang terdapat pada tabel
func GetCheckout(w http.ResponseWriter, r *http.Request) {
	newCheckout := models.GetAllCheckout()
	res, _ := json.Marshal(newCheckout)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan memanggil data yang sesuai dengan ID yang direquest
func GetCheckoutById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CheckoutId := vars["CheckoutId"]
	ID, err := strconv.ParseInt(CheckoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CheckoutDetails, _ := models.GetCheckoutbyId(ID)
	res, _ := json.Marshal(CheckoutDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menmbah user baru pada database
func CreateCheckout(w http.ResponseWriter, r *http.Request) {
	CreateCheckout := &models.Checkout{}
	utils.ParseBody(r, CreateCheckout)
	b := CreateCheckout.CreateCheckout()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menhapus data pada database
func DeleteCheckout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CheckoutId := vars["CheckoutId"]
	ID, err := strconv.ParseInt(CheckoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	checkout := models.DeleteCheckout(ID)
	res, _ := json.Marshal(checkout)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi unutk mengedit data pada database
func UpdateCheckout(w http.ResponseWriter, r *http.Request) {
	var updateCheckout = &models.Checkout{}
	utils.ParseBody(r, updateCheckout)
	vars := mux.Vars(r)
	CheckoutId := vars["CheckoutId"]
	ID, err := strconv.ParseInt(CheckoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CheckoutDetails, db := models.GetCheckoutbyId(ID)
	if updateCheckout.User_id != 0 { // pengkondisian untuk mengubah nilai dari jurusan jika terdapat perubahan yang dilakukan
		CheckoutDetails.User_id = updateCheckout.User_id
	}
	if updateCheckout.Order_number != "" { // pengkondisian untuk mengubah nilai dari angkatan jika terdapat perubahan yang dilakukan
		CheckoutDetails.Order_number = updateCheckout.Order_number
	}
	if updateCheckout.Name != "" { // pengkondisian untuk mengubah nilai dari status aktif jika terdapat perubahan yang dilakukan
		CheckoutDetails.Name = updateCheckout.Name
	}
	if updateCheckout.No_hp != "" { // pengkondisian untuk mengubah nilai dari username jika terdapat perubahan yang dilakukan
		CheckoutDetails.No_hp = updateCheckout.No_hp
	}
	if updateCheckout.Alamat != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		CheckoutDetails.Alamat = updateCheckout.Alamat
	}
	if updateCheckout.Ucapan != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		CheckoutDetails.Ucapan = updateCheckout.Ucapan
	}
	if updateCheckout.Image != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		CheckoutDetails.Image = updateCheckout.Image
	}
	if updateCheckout.Total != 0 { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		CheckoutDetails.Total = updateCheckout.Total
	}
	if updateCheckout.Status != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		CheckoutDetails.Status = updateCheckout.Status
	}
	db.Save(&CheckoutDetails) // mensave hasil perubahan
	res, _ := json.Marshal(CheckoutDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
