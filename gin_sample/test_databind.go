package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Hobby    []string `json:"hobby"`
	Gender   string   `json:"gender"`
	City     string   `json:"city"`
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "test_form.html", nil)
}

func Register(c *gin.Context) {
	print(c.Request.Context())
	var user User
	c.ShouldBindJSON(&user)
	s, _ := json.Marshal(user)
	c.String(http.StatusOK, "user: %s", string(s))
}

func main() {

	e := gin.Default()
	e.LoadHTMLGlob("./templates/*")
	e.GET("/register", GoRegister)
	e.POST("/register", Register)
	e.Run(":8888")

}
