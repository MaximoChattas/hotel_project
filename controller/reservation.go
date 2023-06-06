package controller

import (
	"net/http"
	"project/dto"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertReservation(c *gin.Context) {
	var reservationDto dto.ReservationDto
	err := c.BindJSON(&reservationDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservationDto, er := service.Service.InsertReservation(reservationDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusCreated, reservationDto)
}

func GetReservationById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var reservationDto dto.ReservationDto

	reservationDto, err := service.Service.GetReservationById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, reservationDto)
}

func GetReservations(c *gin.Context) {

	var reservationsDto dto.ReservationsDto

	reservationsDto, err := service.Service.GetReservations()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, reservationsDto)
}

func GetReservationsByUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userReservations dto.UserReservationsDto

	userReservations, err := service.Service.GetReservationsByUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, userReservations)
}

func GetReservationsByHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hotelReservations dto.HotelReservationsDto

	hotelReservations, err := service.Service.GetReservationsByHotel(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelReservations)

}
