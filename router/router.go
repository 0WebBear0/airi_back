package router

import (
	"airi_back/view"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	//Users
	router.GET("/users", view.ShowAllUsers)
	router.POST("/user/add", view.CreateUser)

	router.Run()

}
