-- +goose Up

-- insert user
INSERT INTO users (id, name, email, password)
VALUES (
    gen_random_uuid(),
    'Test User',
    'test@example.com',
    '$2a$12$t9IZPqggXw.0MLrQMG67EecHIfRoTCuVYLKCEy98/VaVg1TDcWUlO'
);

-- insert project
INSERT INTO projects (id, name, description, owner_id)
VALUES (
    gen_random_uuid(),
    'Sample Project',
    'This is a sample project',
    (SELECT id FROM users WHERE email = 'test@example.com')
);

-- insert 3 tasks with different statuses
INSERT INTO tasks (id, title, status, priority, project_id)
VALUES
(
    gen_random_uuid(),
    'Task 1',
    'todo',
    'high',
    (SELECT id FROM projects WHERE name = 'Sample Project')
),
(
    gen_random_uuid(),
    'Task 2',
    'in_progress',
    'medium',
    (SELECT id FROM projects WHERE name = 'Sample Project')
),
(
    gen_random_uuid(),
    'Task 3',
    'done',
    'low',
    (SELECT id FROM projects WHERE name = 'Sample Project')
);
