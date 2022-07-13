package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ALTA-BE10-DIFA/API2/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) GetAllData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp []model.User
		err := uc.DB.Find(&tmp).Error

		if err != nil {
			log.Println("Cannot retrieve object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Server Error")
		}
		res := map[string]interface{}{
			"message": "Get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp model.User

		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}
		err = uc.DB.Where("ID = ?", cnv).First(&tmp).Error
		if err != nil {
			log.Println("There is a problem with data", err.Error())
			return c.JSON(http.StatusBadRequest, "no data")
		}
		res := map[string]interface{}{
			"message": "Get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp model.User
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Server Error")
		}

		err = uc.DB.Create(&tmp).Error
		if err != nil {
			log.Println("Cannot create object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Server Error")
		}
		res := map[string]interface{}{
			"message": "Berhasil input data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		qry := map[string]interface{}{}
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "Cannot convert id")
		}
		var tmp model.User
		err = c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		if tmp.Nama != "" {
			qry["nama"] = tmp.Nama
		}
		if tmp.Email != "" {
			qry["email"] = tmp.Email
		}
		if tmp.Password != "" {
			qry["password"] = tmp.Password
		}

		var ret model.User
		err = uc.DB.Model(&ret).Where("ID = ?", cnv).Updates(qry).Error
		if err != nil {
			log.Println("Cannot update data", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot Update")
		}
		res := map[string]interface{}{
			"message": "Succes update data",
			"data":    ret,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "Cannot convert id")
		}
		err = uc.DB.Where("ID = ?", cnv).Delete(&model.User{}).Error
		if err != nil {
			log.Println("Cannot delete data", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}
		res := map[string]interface{}{
			"message": "Success delete data",
		}
		return c.JSON(http.StatusOK, res)
	}
}
