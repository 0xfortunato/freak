package controllers

import (
	"fmt"
	"log"

	"github.com/0xfortunato/freak/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type Server struct {
	DB     *gorm.DB
	Router *echo.Echo
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("cannot connect to %s database", Dbdriver)
			log.Fatal("error: ", err)
		} else {
			fmt.Printf("we are connected to the %s database", Dbdriver)
		}
	}

	// database migration
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = echo.New()

	server.initializeRoutes()
}
