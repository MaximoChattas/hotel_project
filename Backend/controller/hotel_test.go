package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"project/db"
	"project/dto"
	"project/model"
	"strings"
	"testing"
)

//REMEMBER TO CHANGE DB CONNECTION BEFORE EXECUTING TESTS
//db_name: miranda_test
//THESE TEST CLEAR THE "hotels" TABLE

type TestSuiteEnv struct {
	suite.Suite
	DB *gorm.DB
}

// Running before all tests are completed
func (suite *TestSuiteEnv) SetupSuite() {
	db.StartDbEngine()
	suite.DB = db.Db

	db.Db.Delete(&model.Hotel{})
	db.Db.DropTable(&model.Hotel{})
	db.Db.AutoMigrate(&model.Hotel{})
}

// Running after all tests are completed
func (suite *TestSuiteEnv) TearDownSuite() {
	db.Db.Delete(&model.Hotel{})
	db.Db.DropTable(&model.Hotel{})
	db.Db.AutoMigrate(&model.Hotel{})
	suite.DB.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}

func TestInsertHotel(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.POST("/hotel", InsertHotel)

	body := `{
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

	db.Db.DropTable(&model.Hotel{})
	db.Db.AutoMigrate(&model.Hotel{})

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
