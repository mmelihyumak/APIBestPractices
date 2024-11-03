package handlers

import (
	"net/http"
	"encoding/json"
	"goapibestpractices/middlewares"
	//"goapibestpractices/models"
	//"goapibestpractices/repository"
)

type AuthRequestModel struct{
	Username string
	Password string
}

type AuthResponseModel struct{
	IsSuccess bool
	Message string
	Token string
}

var user = AuthRequestModel{
	Username: "user",
	Password: "user123*",
}

func AuthHandler(w http.ResponseWriter, r *http.Request){

	var authRequest AuthRequestModel
	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if (user.Username != authRequest.Username || user.Password != authRequest.Password){
		http.Error(w, "Kullanıcı adı veya şifre hatalı", http.StatusNotFound)
		return 
	}

	token, err := middlewares.CreateToken(authRequest.Username)
	if err != nil{
		http.Error(w, "token oluşturulamadı", http.StatusInternalServerError)
		return
	}

	response := AuthResponseModel{
		IsSuccess: true,
		Message: "Token oluşturuldu",
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}