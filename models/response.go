package models

import "time"

type UserAuthRes struct {
	ID    string `json:"userId"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type CreateProjectRes struct {
	ID          string    `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     string    `json:"ownerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetAllProjectRes struct {
	ID          string       `json:"projectId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerId     string       `json:"ownerId"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Tasks       []GetTaskRes `json:"tasks"`
}

type GetProjectRes struct {
	ID          string       `json:"projectId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerId     string       `json:"ownerId"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Tasks       []GetTaskRes `json:"tasks"`
}

type GetTaskRes struct {
	ID          string    `json:"taskId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	ProjectId   string    `json:"projectId"`
	AssigneeId  string    `json:"assigneeId"`
	DueDate     time.Time `json:"dueDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
