package routes

import (
"github.com/gorilla/mux"
"PA2-Toko/pkg/controllers"
)

var RegisterTokoRoutes = func(router *mux.Router) {
router.HandleFunc("/toko",controllers.CreateToko).Methods("POST")
router.HandleFunc("/toko",controllers.GetToko).Methods("GET")
router.HandleFunc("/toko/{TokoId}",controllers.GetTokoById).Methods("GET")
router.HandleFunc("/toko/{TokoId}",controllers.UpdateToko).Methods("PUT")
router.HandleFunc("/toko/{TokoId}",controllers.DeleteToko).Methods("DELETE")
}