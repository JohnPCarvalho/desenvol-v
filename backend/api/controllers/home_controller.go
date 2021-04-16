package controllers

import (
	"net/http"

	"github.com/JohnPCarvalho/desenvol-v/backend/api/responses"
)

func (server *Server) Home (w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome API")
}
