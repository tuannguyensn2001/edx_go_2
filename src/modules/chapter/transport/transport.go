package chaptertransport

import (
	"edx_go_2/src/app"
	"edx_go_2/src/config"
	chapterrepository "edx_go_2/src/modules/chapter/repository"
	chapterservice "edx_go_2/src/modules/chapter/service"
	chapterstruct "edx_go_2/src/modules/chapter/struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type transport struct {
	service *chapterservice.Service
	conf    config.Config
}

func NewTransport() *transport {
	t := transport{conf: config.Conf}

	r := chapterrepository.NewRepository(t.conf.Db)
	s := chapterservice.NewService(r)

	t.service = s

	return &t
}

func (t *transport) Create(ctx *gin.Context) {
	var input chapterstruct.CreateChapterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	result, err := t.service.Create(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Create chapter successfully", result))

}

func (t *transport) Update(ctx *gin.Context) {
	chapterId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	var input chapterstruct.UpdateChapterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	chapter, err := t.service.Update(ctx, chapterId, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Update chapter successfully", chapter))

}

func (t *transport) Delete(ctx *gin.Context) {
	chapterId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	err = t.service.Delete(ctx.Request.Context(), chapterId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Delete chapter successfully", nil))

}
