package routes

import (
	authtransport "edx_go_2/src/modules/auth/transport"
	"github.com/gin-gonic/gin"
)

func DeclareRoute(r *gin.Engine) {

	authTranport := authtransport.NewTransport()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authTranport.Register)
	}
}
