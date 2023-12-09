package main

import (
	"log"
	"net/http"

	"github.com/iamtonmoy0/golang-mongodb-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	// listen on port 7000
	log.Fatal(http.ListenAndServe(":7000", r))
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("connection string")
	if err != nil {
		panic(err)
	}

	return s
}
