package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/0xfortunato/freak/api/controllers"
	"github.com/0xfortunato/freak/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("/../../.env"))
	if err != nil {
		log.Fatalf("error getting env %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("cannot connect to %s database\n", TestDbDriver)
			log.Fatal("this is the error:", err)
		} else {
			fmt.Printf("we are connected to the %s database\n", TestDbDriver)
		}
	}
}
