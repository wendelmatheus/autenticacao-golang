package handlers

import (
	"encoding/json"
	"go-jwt/models"
	"go-jwt/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Password hashing error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Nome:  input.Nome,
		Email: input.Email,
		Senha: hash,
	}

	err = models.CreateUser(&user)
	if err != nil {
		log.Println("Error creating user:", err.Error())
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	log.Println("Email received:", creds.Email)
	log.Println("Password received:", creds.Password)

	user, err := models.GetUserByEmail(creds.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	ok := utils.CheckPasswordHash(creds.Password, user.Senha)

	if !ok {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	hash, _ := utils.HashPassword(u.Senha)
	u.Senha = hash

	err := models.UpdateUser(id, u.Nome, u.Email, u.Senha)
	if err != nil {
		http.Error(w, "Error when updating", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := models.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error when deleting", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
