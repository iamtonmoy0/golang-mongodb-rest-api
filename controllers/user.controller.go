package controllers

import (
	"fmt"
	"net/http"
	"os/user"

	"github.com/iamtonmoy0/mongodb-rest-api/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
return &UserController{s}
}

func (us UserController) getUser (w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//get user by id from db and return it as json object to the client
	id:= p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid:= bson.ObjectIdHex(id)
	u:= models.User{}

	if err:=uc.Session.DB("golang").C("users").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}
	uj ,err:=json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",uj)
}
createUser
deleteUser
