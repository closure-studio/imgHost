package resp

import "github.com/gofiber/fiber/v3"

// Response 统一格式
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func Success(c fiber.Ctx, data interface{}, message string) error {

	return c.Status(200).JSON(Response{
		Code:    1,
		Data:    data,
		Message: message,
	})
}

func Failed(c fiber.Ctx, code int, message string) error {
	return c.Status(500).JSON(Response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
}
