package routes

import (
	"github.com/Atul-Kumar-Rana/FocusPod/controllers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router){
	r.HandleFunc("/signup",controllers.Signup).Methods("POST")
	r.HandleFunc("/login",controllers.Login).Methods("POST")
	
}