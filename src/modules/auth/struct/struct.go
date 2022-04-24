package authstruct

type RegisterInput struct {
	Username string `form:"username" json:"username" binding:"required" `
	Email    string `form:"email" json:"email" binding:"required" validate:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}
