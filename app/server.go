package main

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
)

func server(e *echo.Echo) {
	log.Fatal(e.Start(viper.GetString("server.address")))
}
