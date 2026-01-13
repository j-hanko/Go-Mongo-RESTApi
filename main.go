package main

import (
	"net/http"

	"github.com/j-hanko/Go-Mongo-RESTApi/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	router := httprouter.New()
	userController := controllers.NewUserController(getSession())
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
	http.ListenAndServe(":9000", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://mongo:27017")
	if err != nil {
		panic(err)
	}
	return session
}
