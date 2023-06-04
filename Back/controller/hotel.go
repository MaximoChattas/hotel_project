package controller

import (
	"net/http"
	"project/dto"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.Service.InsertHotel(hotelDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
	}

	c.JSON(http.StatusCreated, userDto)
}

func GetHotelById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.Service.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {

	var hotelsDto dto.HotelsDto

	hotelsDto, err := service.Service.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}
