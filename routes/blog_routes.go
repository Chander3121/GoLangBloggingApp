package routes

import (
	"go-blog/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.Engine) {
	r.POST("/blogs", controllers.CreateBlog)
	r.GET("/blogs", controllers.GetBlogs)
	r.GET("/blogs/:id", controllers.GetBlog)
	r.PUT("/blogs/:id", controllers.UpdateBlog)
	r.PATCH("/blogs/:id", controllers.UpdateBlog)
	r.DELETE("/blogs/:id", controllers.DeleteBlog)
}
