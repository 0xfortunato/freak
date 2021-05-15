package controllers

import "github.com/0xfortunato/freak/api/middlewares"

func (s *Server) initalizeRoutes() {

	// home route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// login route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
}
