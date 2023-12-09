package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	r.GET("")
	r.POST()
	r.DELETE()

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("connection string")
	if err != nil {
		panic(err)
	}
	return s
}
