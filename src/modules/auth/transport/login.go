package authtransport

import (
	"edx_go_2/src/app"
	authstruct "edx_go_2/src/modules/auth/struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *transport) Login(ctx *gin.Context) {
	var input authstruct.LoginInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	resp, err := t.service.Login(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Login successfully", resp))

}
