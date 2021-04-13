package models

import (
	"time"

	"github.com/JohnPCarvalho/desenvol-v/backend/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Travel struct {
	gorm.Model
	//Id			string `json:"id"`
	TravelledKilometer		float32 `json:"travelledKilometers"`
	LiterSpent						float32 `json:"literSpent"`
	PricePerLiter					float32 `json:"pricePerLiter"`
	CheckoutDate					time.Time `json:"checkoutDate"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Travel{})
}

func (t *Travel) CreateTravel() *Travel {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func getAllTravels() []Travel {
	var Travels []Travel
	db.Find(&Travels)
	return Travels
}

func getTravelById(Id int64) (*Travel, *gorm.DB) {
	var getTravel Travel
	db := db.Where("ID = ?", Id).Find(&getTravel)
	return &getTravel, db
}

func DeleteTravel(Id int64) Travel {
	var travel Travel
	db.Where("Id= ?", Id).Delete(&travel)
	return travel
}