package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Главная страница",
		})
	})

	router.POST("/", request)

	router.GET("/admin", admin)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "Авторизация",
		})
	})

	router.POST("/login", auth)

	router.Run()
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func auth(c *gin.Context) {
	var form Login

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "admin" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Неудачная авторизация"})
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}

func admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"admin": "admin",
	})
}

type Request struct {
	Number string `form:"number" json:"number" xml:"number"  binding:"required"`
	Text   string `form:"text" json:"text" xml:"text" binding:"required"`
}

func request(c *gin.Context) {
	var form Request

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Number != "" || form.Text != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Неверно заполненные поля"})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Заяка успешно отправлена"})
	}
}
