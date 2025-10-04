package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/Atul-Kumar-Rana/FocusPod/controllers"
)

func RegisterAuthRoutes(e *echo.Echo) {
	e.POST("/users/signup", controllers.SignUp)
	e.POST("/users/login", controllers.Login)
}
