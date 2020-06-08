package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/config"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/route"
	"golang.org/x/crypto/bcrypt"
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

	password := []byte("admin")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err)
	defer db.DBCon.Close()
	e.Logger.Fatal(e.Start(":1234"))
}
