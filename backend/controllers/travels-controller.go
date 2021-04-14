package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JohnPCarvalho/desenvol-v/backend/models"
	"github.com/JohnPCarvalho/desenvol-v/backend/utils"
	"github.com/gorilla/mux"
)

var NewTravel models.Travel

func CreateTravel(w http.ResponseWriter, r *http.Request) {
	CreateTravel := &models.Travel{}
	utils.ParseBody(r, CreateTravel)
	t := CreateTravel.CreateTravel()
	res, _ := json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTravel(w http.ResponseWriter, r *http.Request) {
	newTravels := models.GetAllTravels()
	res, _ := json.Marshal(newTravels)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTravelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	travelId := vars["travelId"]
	ID, err := strconv.ParseInt(travelId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	travelDetails, _ := models.GetTravelById(ID)
	res, _ := json.Marshal(travelDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTravel(w http.ResponseWriter, r *http.Request) {
	var updateTravel = &models.Travel{}
	utils.ParseBody(r, updateTravel)
	vars := mux.Vars(r)
	travelId := vars["travelId"]
	ID, err := strconv.ParseInt(travelId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	travelDetails, db := models.GetTravelById(ID)
	if updateTravel.TravelledKilometer != 0.0{
		travelDetails.TravelledKilometer = updateTravel.TravelledKilometer
	}
	if updateTravel.LiterSpent != 0.0 {
		travelDetails.LiterSpent = updateTravel.LiterSpent
	}
	if updateTravel.PricePerLiter != 0.0 {
		travelDetails.PricePerLiter = updateTravel.PricePerLiter
	} 
	if !updateTravel.CheckoutDate.IsZero() {
		travelDetails.CheckoutDate = updateTravel.CheckoutDate
	}
	db.Save(&travelDetails)
	res, _ := json.Marshal(travelDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	travelId := vars ["travelId"]
	ID, err := strconv.ParseInt(travelId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	travel := models.DeleteTravel(ID)
	res, _ := json.Marshal(travel)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}