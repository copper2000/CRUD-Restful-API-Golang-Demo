package product_api

import (
	"../../config"
	"../../entities"
	"../../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	// get db
	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		products, err2 := productModel.FindAll()

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func FindAllWithPaging(response http.ResponseWriter, request *http.Request) {
	// get db
	db, err := config.GetDatabase()

	// get keyword value from url
	params := mux.Vars(request)
	pageIndex, _ := strconv.Atoi(params["index"])
	pageSize, _ := strconv.Atoi(params["size"])

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		products, err2 := productModel.FindAllWithPaging(pageIndex, pageSize)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func SearchByName(response http.ResponseWriter, request *http.Request) {

	// get keyword value from url
	params := mux.Vars(request)
	keyword := params["keyword"]
	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		products, err2 := productModel.SearchByName(keyword)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var product models.Product
	_ = json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		err2 := productModel.Create(&product)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product models.Product
	_ = json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		err2 := productModel.Update(&product)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	var product models.Product
	_ = json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		err2 := productModel.Delete(&product)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, http.StatusOK)
		}
	}
}

func SearchByPricesRange(response http.ResponseWriter, request *http.Request) {

	// get keyword value from url
	params := mux.Vars(request)
	min, _ := strconv.ParseFloat(params["min"], 64)
	max, _ := strconv.ParseFloat(params["max"], 64)

	db, err := config.GetDatabase()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := entities.ProductModel{Db: db}

		products, err2 := productModel.SearchByPricesRange(min, max)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
