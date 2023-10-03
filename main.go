package main

import (
	"os"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"github.com/iamtonmoy0/golang-mongodb-rest-api/controllers"
)



func main(){
	r:= httprouter.New()
	uc:= controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.getUser)
	r.POST("/user",uc.createUser)
	r.DELETE("/user/:id",uc.deleteUser)


}


// get session
func getSession() *mgo.Session{
	dbURL:= os.Getenv("DB")
s,err:=mgo.Dial(dbURL)
if err!=nil{
	panic("error connecting to db")
}
return s
}