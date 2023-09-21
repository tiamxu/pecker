package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tiamxu/pecker/controllers"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	userGroup := r.Group("user")
	{
		userGroup.GET("/login", controllers.LoginGet)
		userGroup.POST("/login", controllers.LoginPost)
		userGroup.POST("/register", controllers.Register)

	}
	gitlabGroup := r.Group("/gitlab")
	{
		gitlabGroup.GET("/list/projects", controllers.GetProjects)
	}

}

// func InitGitlabRouter(r *gin.RouterGroup) {

// }

// func InitRouter() *gin.Engine {
// 	router := gin.Default()
// 	router.GET("/", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "hello %s", name)
// 	})

// 	router.GET("/projects", controllers.GetProjects)
// 	return router
// }
