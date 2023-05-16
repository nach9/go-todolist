package dtoActivity

type CreateActivityBody struct {
	Title string `form:"title" json:"title" binding:"required"`
	Email string `form:"email" json:"email" binding:"required,email"`
}

type UpdateActivityBody struct {
	Title string `form:"title" json:"title" binding:"required"`
}
