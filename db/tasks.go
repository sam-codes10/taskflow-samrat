package db

import (
	"taskflow-samrat/models"
	"taskflow-samrat/resources"
)

func GetAllTasksByProjectId(projectId, status, assignee_id string) ([]models.GetTaskRes, error) {
	var res []models.GetTaskRes
	baseQuery := `SELECT id,title,description,status,priority,project_id,assignee_id,due_date,created_at,updated_at FROM tasks WHERE project_id=$1 `
	var args []interface{}
	if status != "" && assignee_id != "" {
		baseQuery += `AND status=$2 AND assignee_id=$3 `
		args = []interface{}{projectId, status, assignee_id}
	} else if status != "" {
		baseQuery += `AND status=$2 `
		args = []interface{}{projectId, status}
	} else if assignee_id != "" {
		baseQuery += `AND assignee_id=$2 `
		args = []interface{}{projectId, assignee_id}
	} else {
		args = []interface{}{projectId}
	}
	rows, err := resources.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var record models.GetTaskRes
		err := rows.Scan(&record.ID, &record.Title, &record.Description, &record.Status, &record.Priority, &record.ProjectId, &record.AssigneeId, &record.DueDate, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, record)
	}
	return res, nil
}

func GetTaskById(taskId string) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	query := `SELECT id,title,description,status,priority,project_id,assignee_id,due_date,created_at,updated_at FROM tasks WHERE id=$1`
	err := resources.DB.QueryRow(query, taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func CreateTaskUsingProjectId(task models.CreateAndUpdateTaskReq, projectId string) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("INSERT INTO tasks (title,description,status,priority,project_id,assignee_id,due_date) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,title,description,status,priority,project_id,assignee_id,due_date,created_at,updated_at", task.Title, task.Description, task.Status, task.Priority, projectId, task.AssigneeId, task.DueDate).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func UpdateTaskById(taskId string, task models.CreateAndUpdateTaskReq) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("UPDATE tasks SET title=$1,description=$2,status=$3,priority=$4,assignee_id=$5,due_date=$6 WHERE id=$7 RETURNING id,title,description,status,priority,project_id,assignee_id,due_date,created_at,updated_at", task.Title, task.Description, task.Status, task.Priority, task.AssigneeId, task.DueDate, taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func DeleteTaskById(taskId string) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("DELETE FROM tasks WHERE id=$1 RETURNING id,title,description,status,priority,project_id,assignee_id,due_date,created_at,updated_at", taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func ValidateProjectIdAndOwnerId(projectId, ownerId string) (bool, error) {
	var res bool
	err := resources.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM projects WHERE id=$1 AND owner_id=$2)", projectId, ownerId).Scan(&res)
	if err != nil {
		return false, err
	}
	return res, nil
}

func ValidateTaskIdAndOwnerId(taskId, ownerId string) (bool, error) {
	var res bool
	err := resources.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks t JOIN projects p ON t.project_id = p.id WHERE t.id=$1 AND p.owner_id=$2)", taskId, ownerId).Scan(&res)
	if err != nil {
		return false, err
	}
	return res, nil
}
