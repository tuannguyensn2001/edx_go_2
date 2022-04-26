package coursestruct

import enumcourse "edx_go_2/src/enums/course"

type CreateCourseInput struct {
	Status      enumcourse.Status `form:"status" binding:"required"`
	Name        string            `json:"name" form:"name" binding:"required"`
	Description string            `json:"description" form:"description" binding:"required"`
	ImageUrl    string            `json:"imageUrl" form:"imageUrl" binding:"required"`
	Price       float64           `form:"price" binding:"required"`
	UserId      int
}

type UpdateCourseInput struct {
	Status      enumcourse.Status `form:"status" binding:"required"`
	Name        string            `json:"name" form:"name" binding:"required"`
	Description string            `json:"description" form:"description" binding:"required"`
	ImageUrl    string            `json:"imageUrl" form:"imageUrl" binding:"required"`
	Price       float64           `form:"price" binding:"required"`
}
