package db

import (
	"taskflow-samrat/models"
	"taskflow-samrat/resources"
)

func GetAllTasksByProjectId(projectId, status, assignee_id string) ([]models.GetTaskRes, error) {
	var res []models.GetTaskRes
	baseQuery := `SELECT id,title,desciption,status,priority,project_id,assignee_id,due_date,created_at,updated_at FROM tasks WHERE project_id=$1`
	if status != "" && assignee_id != "" {
		baseQuery += `AND status =$2 AND assignee_id=$3`
	} else if status != "" {
		baseQuery += `AND status =$2 `
	} else {
		baseQuery += `AND assignee_id=$3`
	}
	rows, err := resources.DB.Query(baseQuery, projectId, status, assignee_id)
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
	query := `SELECT id,title,desciption,status,priority,project_id,assignee_id,due_date,created_at,updated_at FROM tasks WHERE id=$1`
	err := resources.DB.QueryRow(query, taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func CreateTaskUsingProjectId(task models.CreateAndUpdateTaskReq) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("INSERT INTO tasks (title,desciption,status,priority,project_id,assignee_id,due_date) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,title,desciption,status,priority,project_id,assignee_id,due_date,created_at,updated_at", task.Title, task.Description, task.Status, task.Priority, task.ProjectId, task.AssigneeId, task.DueDate).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func UpdateTaskById(taskId string, task models.CreateAndUpdateTaskReq) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("UPDATE tasks SET title=$1,desciption=$2,status=$3,priority=$4,project_id=$5,assignee_id=$6,due_date=$7 WHERE id=$8 RETURNING id,title,desciption,status,priority,project_id,assignee_id,due_date,created_at,updated_at", task.Title, task.Description, task.Status, task.Priority, task.ProjectId, task.AssigneeId, task.DueDate, taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return models.GetTaskRes{}, err
	}
	return res, nil
}

func DeleteTaskById(taskId string) (models.GetTaskRes, error) {
	var res models.GetTaskRes
	err := resources.DB.QueryRow("DELETE FROM tasks WHERE id=$1 RETURNING id,title,desciption,status,priority,project_id,assignee_id,due_date,created_at,updated_at", taskId).Scan(&res.ID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.ProjectId, &res.AssigneeId, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
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
	err := resources.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE id=$1 AND owner_id=$2)", taskId, ownerId).Scan(&res)
	if err != nil {
		return false, err
	}
	return res, nil
}
