package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"

	"github.com/laevenx/golang-crud-sql/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver string, DbUser string, DbPassword string, DbPort string, DbHost string, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Finance{}, &models.Sbn{}, &models.Reksadana{}, &models.ConventionalInvoice{}, &models.ConventionalOsf{}, &models.ProductiveInvoice{}) //database migration

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:4545"},
		AllowCredentials: true,
	})

	handler := c.Handler(server.Router)

	fmt.Println("Listening to port 4545")
	log.Fatal(http.ListenAndServe(addr, handler))
}
