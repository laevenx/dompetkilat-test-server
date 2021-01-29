package controllers

import (
	"github.com/laevenx/dompetkilat-test-server/middlewares"
)

func (s *Server) InitializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")

	//Users routes
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")

	//Finance routes
	s.Router.HandleFunc("/finance", middlewares.SetMiddlewareJSON(s.CreateFinance)).Methods("POST")
	s.Router.HandleFunc("/finance", middlewares.SetMiddlewareJSON(s.GetFinances)).Methods("GET")

	//Sbn routes
	s.Router.HandleFunc("/sbn", middlewares.SetMiddlewareJSON(s.CreateSbn)).Methods("POST")
	s.Router.HandleFunc("/sbn", middlewares.SetMiddlewareJSON(s.GetAllSbn)).Methods("GET")

	//Productive Invoice routes
	s.Router.HandleFunc("/productiveinvoice", middlewares.SetMiddlewareJSON(s.CreateProductiveInvoice)).Methods("POST")
	s.Router.HandleFunc("/productiveinvoice", middlewares.SetMiddlewareJSON(s.GetAllProductiveInvoice)).Methods("GET")

	//Reksadana routes
	s.Router.HandleFunc("/reksadana", middlewares.SetMiddlewareJSON(s.CreateReksadana)).Methods("POST")
	s.Router.HandleFunc("/reksadana", middlewares.SetMiddlewareJSON(s.GetAllReksadana)).Methods("GET")

	//Conventional Invoice routes
	s.Router.HandleFunc("/conventionalinvoice", middlewares.SetMiddlewareJSON(s.CreateConventionalInvoice)).Methods("POST")
	s.Router.HandleFunc("/conventionalinvoice", middlewares.SetMiddlewareJSON(s.GetAllConventionalInvoice)).Methods("GET")

	//Conventional Osf routes
	s.Router.HandleFunc("/conventionalosf", middlewares.SetMiddlewareJSON(s.CreateConventionalOsf)).Methods("POST")
	s.Router.HandleFunc("/conventionalosf", middlewares.SetMiddlewareJSON(s.GetAllConventionalOsf)).Methods("GET")
}
