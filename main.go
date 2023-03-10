package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "strings"
    "github.com/gofiber/swagger" 
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/limiter"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/BrandenM-PM/go-rest-api/initializers"
    "github.com/BrandenM-PM/go-rest-api/controllers"
    _ "github.com/BrandenM-PM/go-rest-api/docs"
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
    app.Use(recover.New())
    app.Use(limiter.New(limiter.Config{
        Next: func(c *fiber.Ctx) bool {
            return strings.Contains(c.Path(), "/swagger")
        },
        Expiration: 10 * time.Second,
        Max: 3,
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

    app.Get("/swagger/*", swagger.HandlerDefault) // default

    fmt.Println("Server is running on port 3000")
    port := os.Getenv("PORT")
    log.Fatal(app.Listen(port)) // listen and serve on 3000
}





// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Branden Morin
// @contact.email brandenmorin14@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
