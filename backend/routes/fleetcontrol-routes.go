package routes

import (
	"github.com/JohnPCarvalho/desenvol-v/backend/controllers"
	"github.com/gorilla/mux"
)

var RegisterFleetControlRoutes = func (router *mux.Router) {
	router.HandleFunc("/route", controllers.CreateTravel).Methods("POST")

}