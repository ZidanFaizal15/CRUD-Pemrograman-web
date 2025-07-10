package controllers

import (
	"barang-bekas/config"
	"barang-bekas/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	config.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if user.ID == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Email atau password salah!"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/products")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Email dan password wajib diisi"})
		return
	}

	// Cek apakah user sudah ada
	var existing models.User
	config.DB.Where("email = ?", email).First(&existing)
	if existing.ID != 0 {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Email sudah terdaftar"})
		return
	}

	user := models.User{
		Email:    email,
		Password: password, // nanti bisa pakai bcrypt
	}
	config.DB.Create(&user)

	c.Redirect(http.StatusFound, "/login")
}
