package controllers

import (
	"echo-api/models"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	phonenumber := c.FormValue("phonenumber")

	result, err := models.StoreEmployee(name, address, phonenumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	phonenumber := c.FormValue("phonenumber")

	// ParseInt
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateEmployee(conv_id, name, address, phonenumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEmployee(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
