package app

import (
	log "github.com/sirupsen/logrus"
	"project/controller"
)

func mapUrls() {

	// Add all methods and its mappings
	router.POST("/user", controller.InsertUser)
	router.GET("/user/:id", controller.GetUserById)
	router.GET("/user", controller.GetUsers)

	router.POST("/hotel", controller.InsertHotel)
	router.GET("/hotel/:id", controller.GetHotelById)
	router.GET("/hotel", controller.GetHotels)
	router.POST("/hotel/:id/images", controller.InsertImages)

	router.POST("/reserve", controller.InsertReservation)
	router.GET("/reservation/:id", controller.GetReservationById)
	router.GET("/reservation", controller.GetReservations)
	router.GET("/user/reservations/:id", controller.GetReservationsByUser)
	router.GET("/user/reservations/:id/range", controller.GetReservationsByUserRange)
	router.GET("/hotel/reservations/:id", controller.GetReservationsByHotel)
	router.DELETE("/reservation/:id", controller.DeleteReservation)

	router.POST("/amenity", controller.InsertAmenity)
	router.GET("/amenity", controller.GetAmenities)

	router.GET("/image/:id", controller.GetImageById)

	router.POST("/login", controller.UserLogin)

	router.GET("/availability", controller.CheckAllAvailability)

	log.Info("Finishing mappings configurations")
}
