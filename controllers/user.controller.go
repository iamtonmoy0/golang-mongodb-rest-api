package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamtonmoy0/golang-mongodb-rest-api/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// get user by id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}
	if err := uc.session.DB("golang-mongo").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

// create user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	uc.session.DB("golang-mongo").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(w, "%s\n", uj)
}

// delete user

func (uc.UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
