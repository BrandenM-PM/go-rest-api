package initializers

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "errors"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError

    // Check if the error is a DB error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        code = fiber.StatusNotFound
    }
    
    // Retrieve the custom status code if it's a *fiber.Error
    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    }

    c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
    return c.Status(code).JSON(&fiber.Map{"error": err.Error()})
}
