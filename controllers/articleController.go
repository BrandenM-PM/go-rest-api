package controllers

import (
	"github.com/BrandenM-PM/go-rest-api/initializers"
	"github.com/BrandenM-PM/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateArticle(c *fiber.Ctx) error{
    article := models.Article{
        Title: c.FormValue("title"),
        Content: c.FormValue("content"),
    }

    result := initializers.DB.Create(&article)
    if result.Error != nil { return result.Error }
    return c.JSON(&article)
}

func GetAllArticles(c *fiber.Ctx) error{
    var articles []models.Article
    result := initializers.DB.Find(&articles)
    if result.Error != nil { return result.Error }
    return c.JSON(&articles)
}

func GetArticle(c *fiber.Ctx) error{
    var article models.Article
    result := initializers.DB.First(&article, c.Params("id")) 
    if result.Error != nil { return result.Error }
    return c.JSON(&article)
}

func UpdateArticle(c *fiber.Ctx) error{
    var article models.Article
    result := initializers.DB.First(&article, c.Params("id"))
    if result.Error != nil { return result.Error }

    article.Title = c.FormValue("title")
    article.Content = c.FormValue("content")
    result = initializers.DB.Save(&article)
    if result.Error != nil { return result.Error }
    return c.JSON(&article)
}

func DeleteArticle(c *fiber.Ctx) error{
    result := initializers.DB.Delete(&models.Article{}, c.Params("id"))
    if result.Error != nil { return result.Error }
    return c.SendString("Article deleted successfully")
}

