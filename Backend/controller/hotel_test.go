package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"project/dto"
	"project/service"
	"strings"
	"testing"
	"time"
)

type TestHotel struct{}

func init() {
	service.HotelService = &TestHotel{}
}

func (t TestHotel) GetHotelById(id int) (dto.HotelDto, error) {

	if id > 10 {
		return dto.HotelDto{}, errors.New("hotel not found")
	}

	return dto.HotelDto{}, nil
}

func (t TestHotel) GetHotels() (dto.HotelsDto, error) {
	return dto.HotelsDto{}, nil
}

func (t TestHotel) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	return hotelDto, nil
}

func (t TestHotel) CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool {
	return true
}

func (t TestHotel) CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error) {
	reservationStart, _ := time.Parse("02-01-2006 15:04", startDate)
	reservationEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if reservationStart.After(reservationEnd) {
		return dto.HotelsDto{}, errors.New("a reservation cant end before it starts")
	}
	return dto.HotelsDto{}, nil
}

func TestInsertHotel(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.POST("/hotel", InsertHotel)

	body := `{
		"id": 1,
        "name": "Hotel Test",
        "room_amount": 10,
        "description": "Test hotel description",
        "street_name": "Test Street",
        "street_number": 123,
        "rate": 4.5
    }`

	req, err := http.NewRequest(http.MethodPost, "/hotel", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusCreated, w.Code)

	var response dto.HotelDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	a.NotEqual(0, response.Id)
}

func TestGetHotelById_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/:id", GetHotelById)

	req, err := http.NewRequest(http.MethodGet, "/hotel/400", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusNotFound, w.Code)

	expectedResponse := `{"error":"hotel not found"}`

	a.Equal(expectedResponse, w.Body.String())
}

func TestGetHotelById_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/:id", GetHotelById)

	req, err := http.NewRequest(http.MethodGet, "/hotel/1", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusOK, w.Code)

}

func TestCheckAllAvailability(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.GET("/availability", CheckAllAvailability)

	req, err := http.NewRequest(http.MethodGet, "/availability?start_date=16-06-2023+15:00&end_date=15-06-2023+11:00", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusBadRequest, w.Code)

	expectedResponse := `{"error":"a reservation cant end before it starts"}`

	a.Equal(w.Body.String(), expectedResponse)
}
