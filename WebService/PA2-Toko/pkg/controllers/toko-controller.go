package controllers

import ( // mengimport library yang dibutuhkan
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"PA2-Toko/pkg/models"
	"PA2-Toko/pkg/utils"
)

var NewToko models.Toko

// fungsi yang digunakan memanggil seluruh data yang terdapat pada tabel
func GetToko(w http.ResponseWriter, r *http.Request) {
	newToko := models.GetAllToko()
	res, _ := json.Marshal(newToko)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan memanggil data yang sesuai dengan ID yang direquest
func GetTokoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TokoId := vars["TokoId"]
	ID, err := strconv.ParseInt(TokoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	TokoDetails, _ := models.GetTokobyId(ID)
	res, _ := json.Marshal(TokoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menmbah user baru pada database
func CreateToko(w http.ResponseWriter, r *http.Request) {
	CreateToko := &models.Toko{}
	utils.ParseBody(r, CreateToko)
	b := CreateToko.CreateToko()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menhapus data pada database
func DeleteToko(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TokoId := vars["TokoId"]
	ID, err := strconv.ParseInt(TokoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toko := models.DeleteToko(ID)
	res, _ := json.Marshal(toko)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi unutk mengedit data pada database
func UpdateToko(w http.ResponseWriter, r *http.Request) {
	var updateToko = &models.Toko{}
	utils.ParseBody(r, updateToko)
	vars := mux.Vars(r)
	TokoId := vars["TokoId"]
	ID, err := strconv.ParseInt(TokoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	TokoDetails, db := models.GetTokobyId(ID)
	if updateToko.Nama_toko != "" { // pengkondisian untuk mengubah nilai dari jurusan jika terdapat perubahan yang dilakukan
		TokoDetails.Nama_toko = updateToko.Nama_toko
	}
	if updateToko.User_id != 0 { // pengkondisian untuk mengubah nilai dari status aktif jika terdapat perubahan yang dilakukan
		TokoDetails.User_id = updateToko.User_id
	}
	if updateToko.Alamat != "" { // pengkondisian untuk mengubah nilai dari username jika terdapat perubahan yang dilakukan
		TokoDetails.Alamat = updateToko.Alamat
	}
	if updateToko.Status != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		TokoDetails.Status = updateToko.Status
	}
	if updateToko.Total_produk != 0 { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		TokoDetails.Total_produk = updateToko.Total_produk
	}
	db.Save(&TokoDetails) // mensave hasil perubahan
	res, _ := json.Marshal(TokoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
