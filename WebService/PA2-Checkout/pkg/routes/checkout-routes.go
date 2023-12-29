package routes

import (
"github.com/gorilla/mux"
"PA2-Checkout/pkg/controllers"
)

var RegisterCheckoutRoutes = func(router *mux.Router) {
router.HandleFunc("/checkout",controllers.CreateCheckout).Methods("POST")
router.HandleFunc("/checkout",controllers.GetCheckout).Methods("GET")
router.HandleFunc("/checkout/{CheckoutId}",controllers.GetCheckoutById).Methods("GET")
router.HandleFunc("/checkout/{CheckoutId}",controllers.UpdateCheckout).Methods("PUT")
router.HandleFunc("/checkout/{CheckoutId}",controllers.DeleteCheckout).Methods("DELETE")
}