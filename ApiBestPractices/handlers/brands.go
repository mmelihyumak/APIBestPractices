package handlers

import (
	"net/http"
	"encoding/json"
	"goapibestpractices/models"
	"goapibestpractices/repository"
	"goapibestpractices/cache"
)

var redisClient = cache.GetRedisClient()
var redisCache = cache.NewRedisCache(redisClient)

type BrandRequestModel struct {
	Name    string `json:"name"`
	Segment int    `json:"segment"`
}

type BrandResponseModel struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"Message"`
}


// @Summary      Marka ekle
// @Description  Bu endpoint yeni bir marka ekler
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  BrandResponseModel
// @Router       /api/brand/{id} [post]
func BrandHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var brandRequest BrandRequestModel
	err := json.NewDecoder(r.Body).Decode(&brandRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if models.DB == nil {
		http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
		return
	}



	repository.AddBrand(models.Brand(brandRequest))

	responseBody := BrandResponseModel{
		IsSuccess: true,
		Message:   "POST request successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseBody)
}