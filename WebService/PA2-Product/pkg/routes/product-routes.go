package routes

import (
"github.com/gorilla/mux"
"PA2-Product/pkg/controllers"
)

var RegisterProductsRoutes = func(router *mux.Router) {
router.HandleFunc("/product",controllers.CreateProduct).Methods("POST")
router.HandleFunc("/product",controllers.GetProduct).Methods("GET")
router.HandleFunc("/product/{ProductId}",controllers.GetProductById).Methods("GET")
router.HandleFunc("/product/{ProductId}",controllers.UpdateProduct).Methods("PUT")
router.HandleFunc("/product/{ProductId}",controllers.DeleteProduct).Methods("DELETE")
}