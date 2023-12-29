package main

import ( // Mengimport library yang dibutuhkan
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"PA2-Toko/pkg/routes"
)

func main() {
	r := mux.NewRouter()                                // memanggil method yang digunkan untuk melakukan routing
	routes.RegisterTokoRoutes(r)                    // mendaftara
	http.Handle("/", r)                                 // mendaftarkan view ke paket http
	fmt.Print("Starting Server localhost:9020")         // mencetak string saat program berhasil terhubung ke server mysql
	log.Fatal(http.ListenAndServe("localhost:9020", r)) // digunakan untuk membuat sekaligus start server baru
}
