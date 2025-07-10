package routes

import (
	"barang-bekas/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Template dan static
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})

	// Auth routes
	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.GET("/register", controllers.ShowRegister)
	r.POST("/register", controllers.Register)

	// Produk routes
	product := r.Group("/products")
	{
		product.GET("", controllers.Index)
		product.GET("/new", controllers.New)
		product.POST("", controllers.Create)
		product.GET("/edit/:id", controllers.Edit)
		product.POST("/update/:id", controllers.Update)
		product.GET("/delete/:id", controllers.Delete)
		product.GET("/export/pdf", controllers.ExportPDF)
		product.GET("/export/excel", controllers.ExportExcel)
	}

	return r
}
