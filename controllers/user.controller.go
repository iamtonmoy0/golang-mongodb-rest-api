package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
return &UserController{s}
}

func (us UserController) getUser (w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//get user by id from db and return it as json object to the client

}
createUser
deleteUser
