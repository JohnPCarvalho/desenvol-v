package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/JohnPCarvalho/desenvol-v/backend/api/auth"
	"github.com/JohnPCarvalho/desenvol-v/backend/api/models"
	"github.com/JohnPCarvalho/desenvol-v/backend/api/responses"
	"github.com/JohnPCarvalho/desenvol-v/backend/api/utils/formaterror"
)

func (server *Server) CreateTravel(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Travel := models.Travel{}
	err = json.Unmarshal(body, &Travel)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Travel.Prepare()
	err = Travel.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tid != Travel.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	TravelCreated, err := Travel.SaveTravel(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, TravelCreated.ID))
	responses.JSON(w, http.StatusCreated, TravelCreated)
}

func (server *Server) GetTravels(w http.ResponseWriter, r *http.Request) {

	Travell := models.Travel{}

	Travells, err := Travell.FindAllTravels(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, Travells)
}

func (server *Server) GetTravel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	Travel := models.Travel{}

	TravelReceived, err := Travel.FindTravelByID(server.DB, tid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, TravelReceived)
}

func (server *Server) UpdateTravel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the Travel id is valid
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the Travel exist
	Travel := models.Travel{}
	err = server.DB.Debug().Model(models.Travel{}).Where("id = ?", tid).Take(&Travel).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Travel not found"))
		return
	}

	// If a user attempt to update a Travel not belonging to him
	if uid != Travel.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data Traveled
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	TravelUpdate := models.Travel{}
	err = json.Unmarshal(body, &TravelUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != TravelUpdate.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	TravelUpdate.Prepare()
	err = TravelUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	TravelUpdate.ID = Travel.ID //this is important to tell the model the Travel id to update, the other update field are set above

	TravelUpdated, err := TravelUpdate.UpdateATravel(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, TravelUpdated)
}

func (server *Server) DeleteTravel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid Travel id given to us?
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the Travel exist
	Travel := models.Travel{}
	err = server.DB.Debug().Model(models.Travel{}).Where("id = ?", tid).Take(&Travel).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this Travell?
	if uid != Travel.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = Travel.DeleteATravel(server.DB, tid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", tid))
	responses.JSON(w, http.StatusNoContent, "")
}