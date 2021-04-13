package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JohnPCarvalho/desenvol-v/backend/models"
	"github.com/JohnPCarvalho/desenvol-v/backend/utils"
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
