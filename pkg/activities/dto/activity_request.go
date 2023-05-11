package dtoActivity

type CreateActivityBody struct {
	Title string `form:"title" binding:"required"`
	Email string `form:"email" binding:"required,email"`
}

type UpdateActivityBody struct {
	Title string `form:"title" binding:"required"`
}
