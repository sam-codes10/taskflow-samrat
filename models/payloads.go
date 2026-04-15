package models

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateProjectReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateProjectReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateAndUpdateTaskReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required" oneOf:"todo,in_progress,done"`
	Priority    string `json:"priority" binding:"required" oneOf:"low,medium,high"`
	AssigneeId  string `json:"assigneeId" binding:"required"`
	DueDate     string `json:"dueDate" binding:"required"`
}
