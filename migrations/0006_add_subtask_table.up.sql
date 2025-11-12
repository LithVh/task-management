CREATE TABLE subtask (
    id BIGSERIAL PRIMARY KEY,
    task_id BIGINT REFERENCES task(id) NOT NULL,
    assigned_to UUID REFERENCES app_user(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'todo',
    priority VARCHAR(20),
    due_date TIMESTAMP,
    completed BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_subtask_task_id ON subtask(task_id);

CREATE INDEX idx_subtask_assigned_to ON subtask(assigned_to);
