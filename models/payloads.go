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
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	ProjectId   string `json:"projectId"`
	AssigneeId  string `json:"assignee_id"`
	DueDate     string `json:"due_date"`
}
