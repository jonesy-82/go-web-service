package main

import (
	"inventoryservice/product"
	"inventoryservice/receipt"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	product.SetupRoutes(basePath)
	receipt.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
