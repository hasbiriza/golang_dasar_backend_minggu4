package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	Name  string
	Price int
	Stock int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World, Ini lagi coba Air Golang , mantap")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			products := []product{ 
				{Name: "Baju", Price: 70000, Stock: 12},
				{Name: "Celana", Price: 50000, Stock: 20},
				{Name: "Topi", Price: 30000, Stock: 15},
			}
			res, err := json.Marshal(products)
			if err != nil {
				http.Error(w, "Gagal Konversi JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		} else {
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server Running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
