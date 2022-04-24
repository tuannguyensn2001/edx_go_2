package authtransport

import (
	"edx_go_2/src/app"
	authstruct "edx_go_2/src/modules/auth/struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *transport) Register(ctx *gin.Context) {
	var input authstruct.RegisterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	user, err := t.service.Register(ctx.Request.Context(), input)

	if err != nil {
		panic(app.NewErrorResponse("Register failed", http.StatusInternalServerError, nil, err))
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Register successfully", user))

}
