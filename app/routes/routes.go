package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_web_apps/todos"
)

func SetUpRoutes(app *fiber.App) {
	//Your routes here
	//Example
	api := app.Group("/api")
	// API routes
	api.Get("/todos", todos.GetTodos)
	api.Post("/todos", todos.AddTodo)
	api.Put("/todos/:id", todos.UpdateTodo)
	api.Delete("/todos/:id", todos.DeleteTodo)

	// Serve the frontend (public directory)
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app.Static("/", "./public")
	// JWT Middleware
	// To use middleware uncomment this
	//middleware.JWTConfig(app)
	// Put your endpoint below here to Make auth with JWT
}
