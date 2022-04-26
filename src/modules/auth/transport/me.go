package authtransport

import (
	"edx_go_2/src/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *transport) Me(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	user, err := t.service.Me(ctx.Request.Context(), userId.(int))

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Get information user successfully", user))

}
