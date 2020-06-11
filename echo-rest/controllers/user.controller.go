package controllers

import (
	"net/http"
	"strconv"

	"github.com/bayuwidia/echo-rest/models"
	"github.com/bayuwidia/echo-rest/utils"
	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")

	hash, _ := utils.HashPassword(password)

	// match := utils.CheckPasswordHash(password, hash)
	// fmt.Println("Match:   ", match)

	result, err := models.StoreUser(name, email, hash, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")

	hash, _ := utils.HashPassword(password)

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(convId, name, email, hash, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
