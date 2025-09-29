package main

import (
	"controllers/user"
	"log"
	"net/http"

	"github.com/Atul-Kumar-Rana/FocusPod/controllers"
	"github.com/Atul-Kumar-Rana/FocusPod/database"
	"github.com/Atul-Kumar-Rana/FocusPod/routes"
	"github.com/gorilla/mux"
)

func main(){
	database.ConnectDB()
	database.DB.AutoMigrate(&controllers.User{})
	r :=mux.NewRouter()
	routes.RegisterAuthRoutes(r)

	log.Println("server up and runing")
	log.Fatal(http.ListenAndServe(":8080",r))
}