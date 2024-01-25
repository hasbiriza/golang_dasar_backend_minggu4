package controllers

import (
	"dasar_backend_go/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Data_products(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll())
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	} else if r.Method == "POST" {
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}
		// if product.Id <= 0 || product.Name == "" || product.Price <= 0 || product.Stock <= 0 {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	fmt.Fprintf(w, "Invalid product data- Isi Produk dengan Nilai yang benar")
		// 	return
		// }
		item := models.Product{
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		models.Post(&item)
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
}

func Data_product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/product/"):]
	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	http.Error(w, "ID Not Found", http.StatusNotFound)
	// 	return
	// }
	// var foundIndex = -1
	// for i, p := range models.Products {
	// 	if p.Id == id {
	// 		foundIndex = i
	// 		break
	// 	}
	// }
	// if foundIndex == -1 {
	// 	http.Error(w, "Id not Found", http.StatusNotFound)
	// 	return
	// }

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select(id))
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Product
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}
		// if updateProduct.Id <= 0 || updateProduct.Name == "" || updateProduct.Price <= 0 || updateProduct.Stock <= 0 {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	fmt.Fprintf(w, "Invalid product data")
		// 	return
		// }
		newProduct := models.Product{
			Name:  updateProduct.Name,
			Price: updateProduct.Price,
			Stock: updateProduct.Stock,
		}
		models.Updates(id, &newProduct)
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
		models.Deletes(id)
		// _ = append(models.Products[:foundIndex], models.Products[foundIndex+1:]...)
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
}
