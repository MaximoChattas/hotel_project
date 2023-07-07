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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservationDto, er := service.ReservationService.InsertReservation(reservationDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, reservationDto)
}

func GetReservationById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var reservationDto dto.ReservationDto

	reservationDto, err := service.ReservationService.GetReservationById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservationDto)
}

func GetReservations(c *gin.Context) {

	var reservationsDto dto.ReservationsDto

	reservationsDto, err := service.ReservationService.GetReservations()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationsDto)
}

func GetReservationsByUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userReservations dto.UserReservationsDto

	userReservations, err := service.ReservationService.GetReservationsByUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userReservations)
}

func GetReservationsByUserRange(c *gin.Context) {

	var reservationsDto dto.ReservationsDto

	id, _ := strconv.Atoi(c.Param("id"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	reservationsDto, err := service.ReservationService.GetReservationsByUserRange(id, startDate, endDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationsDto)
}

func GetReservationsByHotel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var hotelReservations dto.HotelReservationsDto

	hotelReservations, err := service.ReservationService.GetReservationsByHotel(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelReservations)

}
