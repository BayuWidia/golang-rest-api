package routes

import (
	"net/http"
	"time"

	"github.com/bayuwidia/echo-rest/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo get!")
	})

	e.GET("/settime", func(c echo.Context) error {
		currentTime := time.Now()
		return c.String(http.StatusOK, currentTime.Format("2006-01-02 15:04:05"))
	})

	e.GET("/user", controllers.FetchAllUser)
	e.POST("/user", controllers.StoreUser)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)

	return e
	// e.Logger.Fatal(e.Start(":1234"))
}
