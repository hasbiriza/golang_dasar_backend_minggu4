package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

var Products = []Product{
	{
		Id: 1, Name: "Baju", Price: 70000, Stock: 12,
	}, {
		Id: 2, Name: "kemeja", Price: 100000, Stock: 20,
	}, {
		Id: 3, Name: "jeans", Price: 90000, Stock: 10,
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := json.Marshal(Products)
			if err != nil {
				http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		} else if r.Method == "POST" {
			var product Product
			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Fprintf(w, "Gagal Decode")
				return
			}
			if product.Id <= 0 || product.Name == "" || product.Price <= 0 || product.Stock <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid product data- Isi Produk dengan Nilai yang benar")
				return
			}
			Products = append(Products, product)
			w.WriteHeader(http.StatusCreated)
			msg := map[string]string{
				"Message": "Product Created",
			}
			res, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
				return
			}
			w.Write(res)
		} else {
			http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		idParam := r.URL.Path[len("/product/"):]
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "ID Not Found", http.StatusNotFound)
			return
		}
		var foundIndex = -1
		for i, p := range Products {
			if p.Id == id {
				foundIndex = i
				break
			}
		}
		if foundIndex == -1 {
			http.Error(w, "Id not Found", http.StatusNotFound)
			return
		}

		if r.Method == "GET" {
			res, err := json.Marshal(Products[foundIndex])
			if err != nil {
				http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		} else if r.Method == "PUT" {
			var updateProduct Product
			err := json.NewDecoder(r.Body).Decode(&updateProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Fprintf(w, "Gagal Decode boss")
				return
			}
			if updateProduct.Id <= 0 || updateProduct.Name == "" || updateProduct.Price <= 0 || updateProduct.Stock <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid product data")
				return
			}
			Products[foundIndex] = updateProduct
			msg := map[string]string{
				"Message": "Product Updated",
			}
			res, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
				return
			}
			w.Write(res)
		} else if r.Method == "DELETE" {
			_ = append(Products[:foundIndex], Products[foundIndex+1:]...)
			msg := map[string]string{
				"Message": "Product Deleted",
			}
			res, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
				return
			}
			w.Write(res)
		} else {
			http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
