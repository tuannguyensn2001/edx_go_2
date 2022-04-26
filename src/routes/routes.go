package routes

import (
	"edx_go_2/src/middlewares"
	authtransport "edx_go_2/src/modules/auth/transport"
	chaptertransport "edx_go_2/src/modules/chapter/transport"
	coursetransport "edx_go_2/src/modules/course/transport"
	"github.com/gin-gonic/gin"
)

func DeclareRoute(r *gin.Engine) {

	authTransport := authtransport.NewTransport()
	courseTransport := coursetransport.NewTransport()
	chapterTransport := chaptertransport.NewTransport()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authTransport.Register)
		v1.POST("/auth/login", authTransport.Login)
		v1.GET("/auth/me", middlewares.Auth, authTransport.Me)

		v1.POST("/courses/owner", middlewares.Auth, courseTransport.CreateByOwner)
		v1.PUT("/courses/owner/:id", middlewares.Auth, courseTransport.UpdateByOwner)
		v1.GET("/courses/:id", courseTransport.GetById)

		v1.POST("/chapters", middlewares.Auth, chapterTransport.Create)
		v1.PUT("/chapters/:id", middlewares.Auth, chapterTransport.Update)
		v1.DELETE("/chapters/:id", middlewares.Auth, chapterTransport.Delete)
	}
}
