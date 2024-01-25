package main

import (
	"dasar_backend_go/src/config"
	"dasar_backend_go/src/helper"
	"dasar_backend_go/src/routes"
	"fmt"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Router()
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
