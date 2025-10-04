package main
import (
	"log"
	// "net/http"
	"github.com/Atul-Kumar-Rana/FocusPod/models"
	"github.com/Atul-Kumar-Rana/FocusPod/database"
	"github.com/Atul-Kumar-Rana/FocusPod/routes"
	"github.com/labstack/echo/v4"
)


func main() {
	// connect DB
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	// start Echo
	e := echo.New()

	// register routes
	routes.RegisterAuthRoutes(e)

	log.Println("server up and running on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}