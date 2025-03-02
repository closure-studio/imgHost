package middleware

import (
	"strings"

	"github.com/closure-studio/objectStorage/server/utils/consts"
	"github.com/closure-studio/objectStorage/server/utils/resp"
	"github.com/closure-studio/objectStorage/utils/env"
	"github.com/gofiber/fiber/v3"
)

const (
	ClientIdLowerCaseStr = "clientid"
)

// Auth 中间件
func Auth() fiber.Handler {
	return func(c fiber.Ctx) error {
		clientID, err := getClientID(c)
		if err != nil {
			return resp.Failed(c, fiber.StatusUnauthorized, err.Error())
		}

		// 验证 clientID 是否有效
		if isValidRedroidClientID(clientID) {
			return c.Next()
		}

		if isValidWebClientID(clientID) {
			// 走 JWT 认证
			if err := validateJWT(c, clientID); err != nil {
				return resp.Failed(c, fiber.StatusUnauthorized, err.Error())
			}
			return c.Next()
		}
		return resp.Failed(c, fiber.StatusUnauthorized, consts.ErrInValidClientId.Error())
	}
}

// **获取 ClientID**
func getClientID(c fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()

	for key, values := range headers {
		// 统一转换 key 为小写进行匹配
		if strings.ToLower(key) == ClientIdLowerCaseStr {
			if len(values) == 0 || strings.TrimSpace(values[0]) == "" {
				return "", consts.ErrClientIDEmpty
			}
			return values[0], nil
		}
	}

	return "", consts.ErrClientIDEmpty
}

func isValidRedroidClientID(clientID string) bool {
	return clientID == env.Instance.REDROID_CLIENT_ID
}

func isValidWebClientID(clientID string) bool {
	return clientID == env.Instance.WEBSITES_CLIENT_ID
}

// **JWT 认证**
func validateJWT(c fiber.Ctx, clientID string) error {

	return nil
}
