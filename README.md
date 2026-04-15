# TaskFlow Backend

## 1. Overview
**TaskFlow** is a minimalistic but fully-featured task management system designed for robust backend execution. It allows users to register, securely authenticate via JWT, create distinct Work Projects, and natively append priority-tracked Tasks inside those projects with role-isolated scoping. 

This Backend is constructed securely in **Golang (Gin)** operating identically across environments via Docker Orchestration and powered efficiently by raw PostgreSQL (lib/pq) bypassing ORMs to maximize performance and flexibility.

## 2. Architecture Decisions
**MVC Pattern Architecture:**
We strictly separated concerns implementing an enhanced Controller-Service-Repository Model:
- **Routers & Controllers:** `routers/` handles path mapping whilst `controllers/` serves purely as data validators extracting JSON binding logic, yielding directly to Services.
- **Services:** Heavy business logic and HTTP standard mapping sit cleanly isolated from frameworks inside `services/`.
- **Repository (DB Layer):** Querying happens inside `db/` using explicitly parameterized `database/sql` strings. No ORMs were used to avoid N+1 querying issues.
- **Unified Payloads (`models/` & `apiRes/`):** All structured mapping executes on strongly-typed structures ensuring exact alignment to Frontend contracts natively executing mapping.

## 3. Running Locally
To run this application seamlessly, you only need Docker installed.

```bash
git clone https://github.com/sam-codes10/taskflow-samrat.git
cd taskflow-samrat
cp .env.example .env
docker compose up --build -d
```
The API spins up silently in the background at `http://localhost:8080`.

## 4. Running Migrations
Zero manual intervention is required. Database synchronisation is natively handled by the Multi-Stage orchestrator. 

Upon `docker compose up`, an automated `entrypoint.sh` bash routine intercepts execution flow routing exclusively to `goose` to compile all 6 Upward automated Schema SQL files (including data seeding keys). Only after PostgreSQL returns healthy does internal API process execution boot up.

## 5. Test Credentials
The migration scripts automatically seed test entities you can test endpoints against natively using the following credentials:
```text
Email:    test@example.com
Password: password123
```

## 6. API Reference & cURL Commands
There is a full interactive Swagger Documentation page mounted locally directly mapped over the application logic:
👉 **[Swagger UI: http://localhost:8080/swagger/task-flow-sam/index.html#/](http://localhost:8080/swagger/task-flow-sam/index.html#/)**

#### Execute the following Raw cURLs testing workflows:

**Authentication**
```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name": "Samrat", "email": "samrat@test.com", "password": "password123"}'

# Login (Extract Token returned here!)
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "samrat@test.com", "password": "password123"}'
```

**Projects API** (Replace `<TOKEN>` with your standard Bearer)
```bash
# Create Project
curl -X POST http://localhost:8080/projects \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"name": "New Alpha Platform", "description": "Main workflow"}'

# Get All User Projects
curl -X GET http://localhost:8080/projects -H "Authorization: Bearer <TOKEN>"

# Get Specific Project Tasks
curl -X GET http://localhost:8080/projects/<PROJECT_ID> -H "Authorization: Bearer <TOKEN>"

# Update Project Definition
curl -X PATCH http://localhost:8080/projects/<PROJECT_ID> \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Alpha Platform", "description": "v2 Updated"}'

# Delete Entire Project Context
curl -X DELETE http://localhost:8080/projects/<PROJECT_ID> -H "Authorization: Bearer <TOKEN>"
```

**Tasks API**
```bash
# Create nested Task under Project
curl -X POST http://localhost:8080/projects/<PROJECT_ID>/tasks \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"title": "Implement WebSockets", "priority": "high", "status": "todo"}'

# Fetch Tasks using Filters (Status / Assignee)
curl -X GET "http://localhost:8080/projects/<PROJECT_ID>/tasks?status=todo&assignee_id=<USER_ID>" -H "Authorization: Bearer <TOKEN>"

# Update Task Scopes
curl -X PATCH http://localhost:8080/tasks/<TASK_ID> \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"status": "in_progress", "priority": "low"}'

# Delete specific Task
curl -X DELETE http://localhost:8080/tasks/<TASK_ID> -H "Authorization: Bearer <TOKEN>"
```

## 7. What I'd Do With More Time
While prioritizing strict scope isolation and stable scaling, given excessive resource capacities, the system would immediately scale adding Redis.
*   **Redis Caching Execution:** We'd cache repetitive identical database read queries (ie: Global Workspace `GET /projects`) into an ultra-fast Redis pipeline slashing generic traffic IO onto PostgreSQL allowing the standard Database pool to natively target complex analytical tasks and task manipulations.
