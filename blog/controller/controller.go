package controller

import (
	"blog/dao"
	"blog/model"
	"html/template"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func ListUser(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println(username, password)
	user := model.User{
		Username: username,
		Password: password,
	}

	log.Printf("%v\n", user)
	dao.Mgr.Register(&user)

	c.Redirect(301, "/")
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "user undefined")
		log.Println("user undefined")
	} else {
		if u.Password != password {
			log.Println("password error")
			c.HTML(200, "login.html", "密码错误")
		} else {
			log.Println("login sucess")
			c.Redirect(301, "/")
		}
	}
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "post_index.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)
	c.Redirect(301, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func GoToPost(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	post := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(post.Content))

	c.HTML(200, "post_detail.html", gin.H{
		"Title":   post.Title,
		"Content": template.HTML(content),
	})
}
