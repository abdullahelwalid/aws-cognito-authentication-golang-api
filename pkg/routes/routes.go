package routes

import (
	"github.com/abdullahelwalid/tradelog-go/pkg/controllers"
	"github.com/abdullahelwalid/tradelog-go/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)


var MainRouter = func (app *fiber.App){
	app.Get("/auth", middleware.Authentication, controllers.AuthHandler)
	app.Get("/", controllers.Test)			
	app.Post("/signup", controllers.SignUp)
	app.Post("/confirmsignup", controllers.ConfirmSignUp)
	app.Post("/login", controllers.Login)
}
