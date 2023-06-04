package controller

import (
	"net/http"
	"project/dto"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertUser(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.Service.InsertUser(userDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
	}

	c.JSON(http.StatusCreated, userDto)
}

func GetUserById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.Service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	var usersDto dto.UsersDto

	usersDto, err := service.Service.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserLogin(c *gin.Context) {
	var loginDto dto.UserLoginDto

	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, er := service.Service.UserLogin(loginDto)

	if er != nil {
		c.JSON(http.StatusUnauthorized, er.Error())
	}

	c.JSON(http.StatusAccepted, loginDto)
}
