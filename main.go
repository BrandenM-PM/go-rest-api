package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/limiter"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/BrandenM-PM/go-rest-api/initializers"
    "github.com/BrandenM-PM/go-rest-api/controllers"
)

func init() {
    initializers.LoadEnvVars()
    initializers.ConnectToDB()
}

func main() {
    app := fiber.New(fiber.Config{
        ErrorHandler: initializers.CustomErrorHandler,
    })
    app.Static("/", "./public")

    app.Use(cors.New())
    app.Use(limiter.New(limiter.Config{
        Expiration: 1 * time.Second,
        Max: 5,
    }))
    app.Use(logger.New())

    // TODO: Add authentication middleware

    // TODO: add template engine
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{
            "message": "Hello, World!",
        })
    })

    app.Get("/articles", controllers.GetAllArticles)
    app.Get("/articles/:id", controllers.GetArticle)
    app.Post("/articles", controllers.CreateArticle)
    app.Patch("/articles/:id", controllers.UpdateArticle)
    app.Delete("/articles/:id", controllers.DeleteArticle)

    fmt.Println("Server is running on port 3000")
    port := os.Getenv("PORT")
    log.Fatal(app.Listen(port)) // listen and serve on 3000
}


