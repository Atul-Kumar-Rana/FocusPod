package main
import (
	"log"
	"net/http"
	"github.com/Atul-Kumar-Rana/FocusPod/models"
	"github.com/Atul-Kumar-Rana/FocusPod/database"
	"github.com/Atul-Kumar-Rana/FocusPod/routes"
	"github.com/gorilla/mux"
)

func main(){
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})
	r :=mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	log.Println("server up and runing")
	log.Fatal(http.ListenAndServe(":8080",r))
}