package dtoTodo

type CreateTodoBody struct {
	Title           string `form:"title" json:"title" binding:"required"`
	ActivityGroupID int64  `form:"activity_group_id" json:"activity_group_id" binding:"required,number"`
	IsActive        bool   `form:"is_active" json:"is_active" binding:"required,boolean"`
}

type UpdateTodoBody struct {
	Title    *string `form:"title" json:"title"`
	IsActive *bool   `form:"is_active" json:"is_active"`
	Priority *string `form:"priority" json:"priority"`
}

type GetAllParam struct {
	ActivityGroupID *int64 `form:"activity_group_id" json:"activity_group_id"`
}
