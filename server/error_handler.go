package server

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ValidError struct {
	Error
	Fields map[string]string `json:"fields"`
}

// Default error handler
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusConflict
	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	var ee validator.ValidationErrors
	if errors.As(err, &ee) {
		code = fiber.StatusBadRequest
		validErr := &ValidError{
			Error: Error{
				Message: "잘못된 요청값입니다.",
				Code:    code,
			},
			Fields: make(map[string]string),
		}

		for _, fieldErr := range ee {
			field := fieldErr.Field()
			field = strings.ToLower(field[:1]) + field[1:]
			validErr.Fields[field] = fieldErr.Tag()
		}

		return c.Status(code).JSON(validErr)
	}

	// Set Content-Type: application/json; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	// Return status code with error message
	return c.Status(code).JSON(&Error{
		Message: err.Error(),
		Code:    code,
	})
}
