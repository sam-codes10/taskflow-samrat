package db

import (
	"errors"
	"taskflow-samrat/models"
	"taskflow-samrat/resources"
)

func CreateProject(payload models.CreateProjectReq, ownerId string) (models.CreateProjectRes, error) {
	var project models.CreateProjectRes
	err := resources.DB.QueryRow("INSERT INTO projects (name, description, owner_id) VALUES ($1, $2, $3) RETURNING *", payload.Name, payload.Description, ownerId).Scan(&project.ID, &project.Name, &project.Description, &project.OwnerId, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return models.CreateProjectRes{}, err
	}
	return project, nil
}

func GetAllProjects(ownerId string) ([]models.GetAllProjectRes, error) {
	rows, err := resources.DB.Query(`
		SELECT 
			p.id, p.name, p.description, p.owner_id, p.created_at, p.updated_at,
			t.id, t.title, t.description, t.status, t.priority, t.project_id, t.assignee_id, t.due_date, t.created_at, t.updated_at
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		WHERE p.owner_id = $1
	`, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projectMap := make(map[string]*models.GetAllProjectRes)

	for rows.Next() {
		var p models.GetAllProjectRes
		var t models.GetTaskRes

		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.OwnerId, &p.CreatedAt, &p.UpdatedAt,
			&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.ProjectId, &t.AssigneeId, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if existing, ok := projectMap[p.ID]; ok {
			existing.Tasks = append(existing.Tasks, t)
		} else {
			p.Tasks = []models.GetTaskRes{t}
			projectMap[p.ID] = &p
		}
	}

	var projects []models.GetAllProjectRes
	for _, v := range projectMap {
		projects = append(projects, *v)
	}

	return projects, nil
}

func GetProjectById(projectId string, ownerId string) (models.GetProjectRes, error) {
	rows, err := resources.DB.Query(`
		SELECT 
			p.id, p.name, p.description, p.owner_id, p.created_at, p.updated_at,
			t.id, t.title, t.description, t.status, t.priority, t.project_id, t.assignee_id, t.due_date, t.created_at, t.updated_at
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		WHERE p.id = $1 AND p.owner_id = $2
	`, projectId, ownerId)
	if err != nil {
		return models.GetProjectRes{}, err
	}
	defer rows.Close()

	var project models.GetProjectRes
	var tasks []models.GetTaskRes

	for rows.Next() {
		var task models.GetTaskRes
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &project.OwnerId, &project.CreatedAt, &project.UpdatedAt,
			&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.ProjectId, &task.AssigneeId, &task.DueDate, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return models.GetProjectRes{}, err
		}
		tasks = append(tasks, task)
	}

	project.Tasks = tasks
	return project, nil
}

func UpdateProjectById(projectId string, payload models.UpdateProjectReq, ownerId string) (models.GetProjectRes, error) {
	_, err := resources.DB.Exec("UPDATE projects SET name = $1, description = $2, updated_at = NOW() WHERE id = $3 AND owner_id = $4", payload.Name, payload.Description, projectId, ownerId)
	if err != nil {
		return models.GetProjectRes{}, err
	}

	return GetProjectById(projectId, ownerId)
}

func DeleteProjectById(projectId, ownerId string) (bool, error) {
	query := `
		DELETE FROM projects WHERE id = $1 AND owner_id = $2
	`
	result, err := resources.DB.Exec(query, projectId, ownerId)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, errors.New("project not found")
	}

	return true, nil
}
