package routes

import (
	"dasar_backend_go/src/controllers"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	http.HandleFunc("/products", controllers.Data_products)

	http.HandleFunc("/product/", controllers.Data_product)
}
