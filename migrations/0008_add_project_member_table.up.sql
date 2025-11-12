-- Create project_member table for managing project collaborators
CREATE TABLE project_member (
    ID bigserial PRIMARY KEY,
    Project_id bigint REFERENCES project(ID) ON DELETE CASCADE NOT NULL,
    User_id UUID REFERENCES app_user(ID) ON DELETE CASCADE NOT NULL,
    UNIQUE(Project_id, User_id)
);

-- Create index for faster lookups
CREATE INDEX idx_project_member_project_id ON project_member(Project_id);
CREATE INDEX idx_project_member_user_id ON project_member(User_id);
