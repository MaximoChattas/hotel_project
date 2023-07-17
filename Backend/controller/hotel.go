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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDto, er := service.HotelService.InsertHotel(hotelDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func GetHotelById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {

	var hotelsDto dto.HotelsDto

	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func CheckAllAvailability(c *gin.Context) {

	var hotelsDto dto.HotelsDto

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	hotelsDto, err := service.HotelService.CheckAllAvailability(startDate, endDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func DeleteHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.HotelService.DeleteHotel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}
