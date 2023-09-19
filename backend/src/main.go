package main

import (
	"fmt"

	controlleruser "github.com/baron7151/praha-ddd-go/src/controller/user"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		panic("failed to load env file.")
	}
	infra.InitDB()
	e := echo.New()

	controlleruser.UserController(e)

	e.Logger.Fatal(e.Start(":8080"))
}

// func user(c echo.Context, db *gorm.DB) error {
// 	// Get team and member from the query string
// 	name := c.QueryParam("name")
// 	id := c.QueryParam("id")
// 	db.Create(&infra.User{UserName: name, UserId: id})
// 	return c.String(http.StatusOK, "name:"+name+", id:"+id)
// }
