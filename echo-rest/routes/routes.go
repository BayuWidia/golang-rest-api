package routes

import (
	"net/http"

	"github.com/bayuwidia/echo-rest/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func IsAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Success, you are on the secret admin main page!")
}

func Init() *echo.Echo {
	e := echo.New()

	g := e.Group("/admin")

	// g.Use(middleware.Logger())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.GET("/main", IsAdmin)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo get!")
	})

	e.GET("/user", controllers.GetAllUser)
	e.POST("/user", controllers.StoreUser)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)

	return e
	// e.Logger.Fatal(e.Start(":1234"))
}
