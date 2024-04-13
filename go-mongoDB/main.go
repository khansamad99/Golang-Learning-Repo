package main

import (
	"fmt"
	"go-mongoDB/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/test", uc.Hello)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	err := http.ListenAndServe("localhost:9000", r)

	if err == nil {
		fmt.Println("Server is Running")
	}

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://127.0.0.1:27017/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return s
}
