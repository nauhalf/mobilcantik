package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/config"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/route"
)

func main() {

	//Load Environment
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err.Error())
	}

	e := echo.New()

	route.InitRoutes(e)

	// e.Logger.Fatal(e.Start(":1234"))

	db.Open()

	defer db.DBCon.Close()
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))
}
