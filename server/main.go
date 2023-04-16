package main

import (
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72"
	"net/http"
)


func init() {
	stripe.Key = "pk_test_51MuJinSHZTnL3EyWMnbuZ2O9HyvrhN2HIszaiDrVKhhs8NRsDL1JFIbPxs2Vf8xnp0Ruzk1xAj4ZNH41EkFI97YB00b775Qe8R"
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/checkout", CORSCheck(CheckoutCreator))
	r.HandleFunc("/event", CORSCheck(HandleEvent))

	http.ListenAndServe(":8080", r)
}
