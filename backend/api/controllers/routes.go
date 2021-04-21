package controllers

import "github.com/JohnPCarvalho/desenvol-v/backend/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(middlewares.SetCORS(s.Login))).Methods("POST", "OPTIONS")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/travels", middlewares.SetMiddlewareJSON(s.CreateTravel)).Methods("POST")
	s.Router.HandleFunc("/travels", middlewares.SetMiddlewareJSON(s.GetTravel)).Methods("GET")
	s.Router.HandleFunc("/travels/{id}", middlewares.SetMiddlewareJSON(s.GetTravel)).Methods("GET")
	s.Router.HandleFunc("/travels/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTravel))).Methods("PUT")
	s.Router.HandleFunc("/travels/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteTravel)).Methods("DELETE")
}