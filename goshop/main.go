package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/yourname/goshop/database"
	"github.com/yourname/goshop/handlers"
	"github.com/yourname/goshop/middleware"
)

func main() {
	 
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	app := fiber.New()

	database.ConnectDatabase()

	 
	app.Post("/users", handlers.CreateUser)
	app.Post("/login", handlers.Login)

 
	app.Use(middleware.JWTProtected())

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUserByID)
	app.Get("/products", handlers.GetProducts)

 
	app.Post("/products", middleware.AdminOnly(), handlers.CreateProduct)

 
	app.Post("/cart", handlers.AddToCart)
	app.Get("/cart/:userId", handlers.GetCart)
	app.Delete("/cart/:id", handlers.RemoveFromCart)

	app.Post("/checkout/:userId", handlers.Checkout)
	app.Get("/orders/:userId", handlers.GetOrders)
 
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
