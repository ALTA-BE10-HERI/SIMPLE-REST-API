package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var arrData []User

func GetAll(c echo.Context) error {
	res := map[string]interface{}{
		"message": "Get all data",
		"data":    arrData,
	}
	return c.JSON(http.StatusOK, res)
}

func InsertUser(c echo.Context) error {
	var tmp User
	err := c.Bind(&tmp) //input json ke variable
	if err != nil {
		log.Println("Cannot parser input to object", err.Error())
		return c.JSON(http.StatusInternalServerError, "Error dari server")
	}

	arrData = append(arrData, tmp)
	res := map[string]interface{}{
		"message": "Succes input data",
		"data":    tmp,
	}
	return c.JSON(http.StatusOK, res)
}

func GetSpesificUser(c echo.Context) error {
	param := c.Param("id")
	cnv, err := strconv.Atoi(param) //di convirt karena awalnya link di baca sebagai string
	if err != nil {
		log.Println("Cannot conver to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "canot conver id")
	}
	res := map[string]interface{}{
		"message": "Get data by id",
		"data":    arrData[cnv-1],
		// tanya tentang cnv-1 kok bisa menampilkan 1 data
	}
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.GET("/user", GetAll)
	e.POST("/user", InsertUser)
	e.GET("/user/:id", GetSpesificUser)

	fmt.Println("menjalankan program ...")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error()) //untuk memperhentikan program saat program ada yang eror
	}
}
