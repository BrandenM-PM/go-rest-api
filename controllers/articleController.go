package controllers

import (
    "github.com/BrandenM-PM/go-rest-api/initializers"
    "github.com/BrandenM-PM/go-rest-api/models"
    "github.com/gofiber/fiber/v2"
)
// CreateArticle godoc
// @Summary Creates an Article
// @Produce json
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Success 200 {object} models.Article
// @Router /articles [post]
func CreateArticle(c *fiber.Ctx) error{
    article := models.Article{
        Title: c.FormValue("title"),
        Content: c.FormValue("content"),
    }

    result := initializers.DB.Create(&article)
    if result.Error != nil { return result.Error }
    return c.JSON(&article)
}

// GetAllArticle godoc
// @Summary Retrieves all Articles 
// @Produce json
// @Success 200 {object} []models.Article
// @Router /articles [get]
func GetAllArticles(c *fiber.Ctx) error{
    var articles []models.Article
    result := initializers.DB.Find(&articles)
    if result.Error != nil { return result.Error }
    return c.JSON(&articles)
}

// GetArticle godoc
// @Summary Retrieves an Article based on given ID
// @Produce json
// @Param id path integer true "Article ID"
// @Success 200 {object} models.Article
// @Router /articles/{id} [get]
func GetArticle(c *fiber.Ctx) error{
    var article models.Article
    result := initializers.DB.First(&article, c.Params("id")) 
    if result.Error != nil { return result.Error }
    return c.JSON(&article)
}
// UpdateArticle godoc
// @Summary Updates an Article based on given ID
// @Produce json
// @Param id path integer true "Article ID"
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Success 200 {object} models.Article
// @Router /articles/{id} [patch]
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

// DeleteArticle godoc
// @Summary Deletes an Article based on given ID
// @Produce json
// @Param id path integer true "Article ID"
// @Success 200 {object} models.Article
// @Router /articles/{id} [delete]
func DeleteArticle(c *fiber.Ctx) error{
    result := initializers.DB.Delete(&models.Article{}, c.Params("id"))
    if result.Error != nil { return result.Error }
    return c.JSON(&fiber.Map{
        "message": "Article deleted successfully",
    })
}

