# TaskManagement API

> RESTful API for task management built with Go, Gin, and PostgreSQL

A clean, well-structured task management API demonstrating professional Go backend development practices. Perfect for portfolio and job interviews.

## Features

- **JWT Authentication** - Secure user registration and login
- **User Management** - Profile viewing and updates
- **Project Organization** - Create and manage projects
- **Task Tracking** - Full CRUD operations with status management
- **Priority System** - Low, medium, high task priorities
- **Due Dates** - Track task deadlines
- **Owner-based Authorization** - Users can only access their own data
- **Clean Architecture** - Controller-Service-Repository pattern
- **Feature-based Structure** - Organized by domain, not by layer

## Tech Stack

- **Go 1.21+** - Modern, efficient backend language
- **Gin** - Fast HTTP web framework
- **GORM** - Type-safe ORM
- **PostgreSQL 15** - Reliable relational database
- **JWT** - Stateless authentication
- **Docker** - Containerized deployment
- **bcrypt** - Secure password hashing

## Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 15
- Docker & Docker Compose (optional)

### Installation

1. **Clone the repository**
```bash
git clone <your-repo-url>
cd task-management
```

2. **Copy environment variables**
```bash
cp .env.example .env
```

3. **Start PostgreSQL** (using Docker)
```bash
docker-compose up -d postgres
```

4. **Run migrations**
```bash
make migrate-up
```

5. **Start the server**
```bash
make run
```

The API will be available at `http://localhost:8080`

### Using Docker

Start everything with Docker Compose:
```bash
make docker-up
```

## API Endpoints

### Authentication
```
POST   auth/register   - Register new user
POST   auth/login      - Login user
POST   /api/v1/auth/logout     - Logout user
```

### Users
```
GET    users/me        - Get current user profile
PUT    users/me        - Update user profile
```

### Projects
```
GET    projects        - List all user's projects
POST   projects        - Create new project
GET    projects/:id    - Get project by ID
PUT    projects/:id    - Update project
DELETE projects/:id    - Delete project
```

### Tasks
```
GET    projects/:projectId/tasks  - List tasks in project
POST   projects/:projectId/tasks  - Create task in project
GET    tasks/:id                  - Get task by ID
PUT    tasks/:id                  - Update task
DELETE tasks/:id                  - Delete task
PATCH  tasks/:id/complete         - Toggle task completion
```

### Filters
Tasks can be filtered by query parameters:
- `?status=todo|in_progress|completed`
- `?priority=low|medium|high`

## Project Structure

```
task-management/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── auth/                   # Authentication feature
│   │   ├── controller.go
│   │   ├── service.go
│   │   ├── routes.go
│   │   └── dto.go
│   ├── user/                   # User feature
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── controller.go
│   │   ├── routes.go
│   │   └── dto.go
│   ├── project/                # Project feature
│   │   └── ...
│   ├── task/                   # Task feature
│   │   └── ...
│   ├── config/                 # Configuration
│   │   └── config.go
│   ├── database/               # Database connection
│   │   └── database.go
│   ├── middleware/             # HTTP middlewares
│   │   ├── auth.go
│   │   └── cors.go
│   └── utils/                  # Utility functions
│       ├── jwt.go
├── migrations/                 # Database migrations
├── .env                        # Environment variables template
├── docker-compose.yml         # Docker services
├── Dockerfile                 # API container
└── README.md                  # This file
```

## Database Schema

### app_user
- `id` - UUID (Primary Key)
- `name` - VARCHAR(255)
- `email` - VARCHAR(255) UNIQUE
- `password_hash` - VARCHAR(255)
- `created_at` - TIMESTAMP

### project
- `id` - BIGSERIAL (Primary Key)
- `user_id` - UUID (Foreign Key)
- `name` - VARCHAR(255)
- `description` - VARCHAR(255)
- `created_at` - TIMESTAMP
- `updated_at` - TIMESTAMP

### task
- `id` - BIGSERIAL (Primary Key)
- `project_id` - BIGINT (Foreign Key)
- `title` - VARCHAR(255)
- `description` - TEXT
- `status` - VARCHAR(20) (todo, in_progress, completed)
- `priority` - VARCHAR(20) (low, medium, high)
- `due_date` - TIMESTAMP
- `completed` - BOOLEAN
- `created_at` - TIMESTAMP
- `updated_at` - TIMESTAMP


## Example Usage

### 1. Register a new user
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "abc@gmail.com",
    "password": "123"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Create a project
```bash
curl -X POST http://localhost:8080/projects \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "My Project",
    "description": "Project description"
  }'
```

### 4. Create a task
```bash
curl -X POST http://localhost:8080/projects/1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Implement feature X",
    "description": "Description here",
    "status": "todo",
    "priority": "high"
  }'
```




