package main

import (
	"./apis/product_api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/product_api/find-all", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product_api/find-all/{index}/{size}", product_api.FindAllWithPaging).Methods("GET")
	router.HandleFunc("/api/product_api/search/{keyword}", product_api.SearchByName).Methods("GET")
	router.HandleFunc("/api/product_api/search/{min}/{max}", product_api.SearchByPricesRange).Methods("GET")
	router.HandleFunc("/api/product_api/create", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product_api/update", product_api.Update).Methods("POST")
	router.HandleFunc("/api/product_api/delete", product_api.Delete).Methods("POST")

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println(err)
	}
}
