package controllers

import ( // mengimport library yang dibutuhkan
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"PA2-Product/pkg/models"
	"PA2-Product/pkg/utils"
)

var NewProduct models.Product

// fungsi yang digunakan memanggil seluruh data yang terdapat pada tabel
func GetProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := models.GetAllProduct()
	res, _ := json.Marshal(newProduct)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan memanggil data yang sesuai dengan ID yang direquest
func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProductId := vars["ProductId"]
	ID, err := strconv.ParseInt(ProductId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ProductDetails, _ := models.GetProductbyId(ID)
	res, _ := json.Marshal(ProductDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menmbah user baru pada database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	b := CreateProduct.CreateProduct()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi yang digunakan ketika menhapus data pada database
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProductId := vars["ProductId"]
	ID, err := strconv.ParseInt(ProductId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// fungsi unutk mengedit data pada database
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &models.Product{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	ProductId := vars["ProductId"]
	ID, err := strconv.ParseInt(ProductId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ProductDetails, db := models.GetProductbyId(ID)
	if updateProduct.Name != "" { // pengkondisian untuk mengubah nilai dari jurusan jika terdapat perubahan yang dilakukan
		ProductDetails.Name = updateProduct.Name
	}
	if updateProduct.Toko_id != 0 { // pengkondisian untuk mengubah nilai dari angkatan jika terdapat perubahan yang dilakukan
		ProductDetails.Toko_id = updateProduct.Toko_id
	}
	if updateProduct.User_id != 0 { // pengkondisian untuk mengubah nilai dari status aktif jika terdapat perubahan yang dilakukan
		ProductDetails.User_id = updateProduct.User_id
	}
	if updateProduct.Harga != 0 { // pengkondisian untuk mengubah nilai dari username jika terdapat perubahan yang dilakukan
		ProductDetails.Harga = updateProduct.Harga
	}
	if updateProduct.Kota != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		ProductDetails.Kota = updateProduct.Kota
	}
	if updateProduct.Deskripsi != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		ProductDetails.Deskripsi = updateProduct.Deskripsi
	}
	if updateProduct.Gambar != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		ProductDetails.Gambar = updateProduct.Gambar
	}
	if updateProduct.No_hp != "" { // pengkondisian untuk mengubah nilai dari email akademik jika terdapat perubahan yang dilakukan
		ProductDetails.No_hp = updateProduct.No_hp
	}
	db.Save(&ProductDetails) // mensave hasil perubahan
	res, _ := json.Marshal(ProductDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
