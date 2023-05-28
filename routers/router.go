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

}
func InitGitlabRouter(r *gin.RouterGroup) {
	gitlabGroup := r.Group("/gitlab")
	{
		gitlabGroup.GET("/list/projects", controllers.GetProjects)
	}
}
func User(r *gin.RouterGroup) {

	userGroup := r.Group("user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
}

// func InitRouter() *gin.Engine {
// 	router := gin.Default()
// 	router.GET("/", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "hello %s", name)
// 	})

// 	router.GET("/projects", controllers.GetProjects)
// 	return router
// }
