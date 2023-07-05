package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"project/dto"
	"project/service"
	"strconv"
)

func InsertImages(c *gin.Context) {
	var imagesDto dto.ImagesDto

	id, _ := strconv.Atoi(c.Param("id"))

	form, _ := c.MultipartForm()
	files := form.File["images"]

	for i, file := range files {

		fileExt := path.Ext(file.Filename)

		//Filename as [hotel_id]-[image_number].[file_extension]
		fileName := fmt.Sprintf("%d-%d%s", id, i+1, fileExt)

		if err := c.SaveUploadedFile(file, "../Images/"+fileName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		imageDTO := dto.ImageDto{
			HotelId: id,
			Path:    "Images/" + fileName,
		}
		imagesDto = append(imagesDto, imageDTO)
	}

	imagesDto, err := service.Service.InsertImages(imagesDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}
