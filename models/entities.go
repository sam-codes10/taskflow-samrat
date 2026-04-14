package models

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"-"`
	CreatedAt time.Time `db:"created_at"`
}

type Project struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	OwnerId     string    `db:"owner_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Task struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	Priority    string    `db:"priority"`
	ProjectId   string    `db:"project_id"`
	AssigneeId  string    `db:"assignee_id"`
	DueDate     time.Time `db:"due_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
