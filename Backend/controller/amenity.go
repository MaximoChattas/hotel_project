package controller

import (
	"net/http"
	"project/dto"
	"project/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertAmenity(c *gin.Context) {
	var amenityDto dto.AmenityDto
	err := c.BindJSON(&amenityDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amenityDto, er := service.Service.InsertAmenity(amenityDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, amenityDto)
}

func GetAmenities(c *gin.Context) {

	var amenitiesDto dto.AmenitiesDto

	amenitiesDto, err := service.Service.GetAmenities()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, amenitiesDto)
}
