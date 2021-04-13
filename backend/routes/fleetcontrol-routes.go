package routes

import (
	"github.com/gorilla/mux"
)

var RegisterFleetControlRoutes = func (router *mux.Router) {
	router.HandleFunc("/route", controllers.CreateTravel).Methods("POST")

}