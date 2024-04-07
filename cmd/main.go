package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SimonVillalonIT/filehub/internal/routers"
	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/labstack/echo/v4"
)

func main() {
    config := config.Config{}
	if err := config.Load(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	e := echo.New()

    routers.Setup(e, &config)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Server.Port)))
}
