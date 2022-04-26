package chapterstruct

type CreateChapterInput struct {
	Name     string `form:"name" binding:"required"`
	CourseId int    `form:"courseId" binding:"required"`
}

type UpdateChapterInput struct {
	Name string `form:"name" binding:"required"`
}
