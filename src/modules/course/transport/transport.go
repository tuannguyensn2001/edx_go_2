package coursetransport

import (
	"edx_go_2/src/app"
	"edx_go_2/src/config"
	courserepository "edx_go_2/src/modules/course/repository"
	courseservice "edx_go_2/src/modules/course/service"
	coursestruct "edx_go_2/src/modules/course/struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type transport struct {
	service *courseservice.Service
	conf    config.Config
}

func NewTransport() *transport {
	t := transport{
		conf: config.Conf,
	}

	r := courserepository.NewRepository(t.conf.Db)
	s := courseservice.NewService(r)

	t.service = s

	return &t

}

func (t *transport) CreateByOwner(ctx *gin.Context) {
	var input coursestruct.CreateCourseInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	userId, _ := ctx.Get("user_id")

	input.UserId = userId.(int)

	course, err := t.service.Create(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Create course successfully", course))

}

func (t *transport) UpdateByOwner(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))

	userId, _ := ctx.Get("user_id")

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	var input coursestruct.UpdateCourseInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	course, err := t.service.Update(ctx.Request.Context(), courseId, input, userId.(int))

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Update course successfully", course))

}

func (t *transport) GetById(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(app.ErrInvalidRequest(err))
	}

	course, err := t.service.FindById(ctx.Request.Context(), courseId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, app.NewResponse("Get course successfully", course))
}
