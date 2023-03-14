package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/renaldlie1/depd_go_echo/controllers"
	"github.com/renaldlie1/depd_go_echo/models"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		var res models.Response
		res.Code = http.StatusOK
		res.Status = true
		res.Message = "Hello, this is echo!"

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)

	return e
}
