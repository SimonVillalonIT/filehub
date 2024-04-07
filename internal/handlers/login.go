package handlers

import (
	"net/http"
	"time"

	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/SimonVillalonIT/filehub/pkg/utils"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	Config *config.Config
}

func (login *LoginHandler) Login(c echo.Context) error {

	if c.FormValue("username") == "" || c.FormValue("password") == "" {
		return c.String(http.StatusBadRequest, "Username or password is empty")
	}

	credentials := utils.Credentials{Username: c.FormValue("username"), Password: c.FormValue("password")}

	if !credentials.IsValidCredentials(*login.Config) {
		return c.String(http.StatusBadRequest,  "Wrong credentials")
	}

	jwt, err := utils.GenerateJWT(login.Config, credentials.Username, credentials.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "Bearer " + jwt
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
    c.Response().Header().Set("HX-Redirect", "/files/")
    return c.String(http.StatusSeeOther,"Logging successfully")
}
