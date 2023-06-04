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

	router.POST("/reserve", controller.InsertReservation)
	router.GET("/reservation/:id", controller.GetReservationById)
	router.GET("/reservation", controller.GetReservations)
	router.GET("/user/reservations/:id", controller.GetReservationsByUser)
	router.GET("hotel/reservations/:id", controller.GetReservationsByHotel)

	router.POST("/login", controller.UserLogin)

	log.Info("Finishing mappings configurations")
}
