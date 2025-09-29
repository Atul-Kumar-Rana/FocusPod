package controllers

import (
	"encoding/json"
	"FocusPod./models"
	// "hash"
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/Atul-Kumar-Rana/FocusPod/database"
	// "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
func Signup(w http.ResponseWriter, r *http.Request){
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	hash,err :=bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)

	if err != nil{
		panic(err)
	}
	u.Password=string(hash)
	database.DB.Create(&u)
	w.Write(([]byte("User created")))

}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	var dbUser models.User
	database.DB.Where("email = ?", input.Email).First(&dbUser)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	// generate JWT
	claims := jwt.MapClaims{
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}