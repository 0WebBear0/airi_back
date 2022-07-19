package router

import (
	"airi_back/socket"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body1 struct {
	// json tag to de-serialize json body
	Name    string `json:"name"`
	Message string `json:"message"`
}

func Router1() {

	router := gin.Default()

	// Маршрут 1
	// Соответствие / user / zhangsan не соответствует / user / zhangsan / user
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "you name is "+name)
	})

	// Маршрут 2
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/getAndpost", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")

		body := Body1{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.String(http.StatusOK, "Получить метод: id:"+id+"; страница:"+page+"; Метод публикации: имя:"+body.Name+"; сообщение:"+body.Message)
	})

	//webSockets

	router.GET("/ws", func(c *gin.Context) {
		socket.Wshandler(c.Writer, c.Request)
	})

	router.Run()

}
