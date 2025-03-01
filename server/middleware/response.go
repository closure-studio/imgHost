package middleware

import (
	"github.com/closure-studio/imgHost/server/utils/resp"
	"github.com/gofiber/fiber/v3"
)

func JSONResponseMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return resp.Failed(c, fiber.StatusInternalServerError, err.Error())
		}

		// 检查原始返回的 Content-Type
		contentType := c.Response().Header.ContentType()
		if string(contentType) == fiber.MIMEApplicationJSON {
			return nil // 已经是 JSON，不做任何处理
		}

		// 读取原始响应数据
		bodyBytes := c.Response().Body()
		bodyStr := string(bodyBytes)

		// 重新包装 JSON 响应
		return resp.Success(c, bodyStr, "Success")
	}
}
