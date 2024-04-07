package middlewares

import (
	"net/http"

	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/SimonVillalonIT/filehub/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
    Config *config.Config
}

func (a *AuthMiddleware) Init(next  echo.HandlerFunc) echo.HandlerFunc {
    return  func(c echo.Context) error {
        bearer := c.Request().Header.Get("Authorization")
        if bearer == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error":"missing token"})
        }
        tokenString := bearer[len("Bearer "):]
        if err := utils.ValidateToken(a.Config, tokenString); err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error":err.Error()})
        }
        return next(c)
    }
}
