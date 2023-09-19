package controlleruser

import (
	"fmt"
	"log"
	"net/http"

	appuser "github.com/baron7151/praha-ddd-go/src/app/user"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"github.com/baron7151/praha-ddd-go/src/infra/queryservice"
	"github.com/labstack/echo/v4"
)

func UserController(e *echo.Echo) {
	db, err := infra.ConnectDB()
	if err != nil {
		panic(err)
	}
	userDataQS := queryservice.NewUserDataQS(db)
	getUserDataUsecase := appuser.NewGetUserDataUsecase(userDataQS)

	e.GET("/user", func(c echo.Context) error {
		username := c.QueryParam("user_name")
		fmt.Println(username)
		result, err := getUserDataUsecase.GetUserData(username)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, result)
	})
	//e.POST("/user", createUser)
	//e.PATCH("/user", updateUser)
}
