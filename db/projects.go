package db

import (
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

func GetProjectById(projectId string) (models.GetProjectRes, error) {
	rows, err := resources.DB.Query(`
		SELECT 
			p.id, p.name, p.description, p.owner_id, p.created_at, p.updated_at,
			t.id, t.title, t.description, t.status, t.priority, t.project_id, t.assignee_id, t.due_date, t.created_at, t.updated_at
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		WHERE p.id = $1
	`, projectId)
	if err != nil {
		return models.GetProjectRes{}, err
	}
	defer rows.Close()

	var project models.GetProjectRes
	var tasks []models.GetTaskRes

	for rows.Next() {
		var t models.GetTaskRes
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &project.OwnerId, &project.CreatedAt, &project.UpdatedAt,
			&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.ProjectId, &t.AssigneeId, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return models.GetProjectRes{}, err
		}
		tasks = append(tasks, t)
	}

	project.Tasks = tasks
	return project, nil
}

func UpdateProjectById(projectId string, payload models.UpdateProjectReq) (models.GetProjectRes, error) {
	_, err := resources.DB.Exec("UPDATE projects SET name = $1, description = $2, updated_at = NOW() WHERE id = $3", payload.Name, payload.Description, projectId)
	if err != nil {
		return models.GetProjectRes{}, err
	}

	rows, err := resources.DB.Query(`
		SELECT 
			p.id, p.name, p.description, p.owner_id, p.created_at, p.updated_at,
			t.id, t.title, t.description, t.status, t.priority, t.project_id, t.assignee_id, t.due_date, t.created_at, t.updated_at
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		WHERE p.id = $1
	`, projectId)
	if err != nil {
		return models.GetProjectRes{}, err
	}
	defer rows.Close()

	var project models.GetProjectRes
	var tasks []models.GetTaskRes

	for rows.Next() {
		var t models.GetTaskRes
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &project.OwnerId, &project.CreatedAt, &project.UpdatedAt,
			&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.ProjectId, &t.AssigneeId, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return models.GetProjectRes{}, err
		}
		tasks = append(tasks, t)
	}

	project.Tasks = tasks
	return project, nil
}

func DeleteProjectById(projectId string) (bool, error) {
	query := `
		DELETE FROM projects WHERE id = $1
	`
	_, err := resources.DB.Exec(query, projectId)
	if err != nil {
		return false, err
	}

	return true, nil
}
