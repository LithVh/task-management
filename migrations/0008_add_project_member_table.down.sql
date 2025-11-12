-- Drop project_member table and its indexes
DROP INDEX IF EXISTS idx_project_member_user_id;
DROP INDEX IF EXISTS idx_project_member_project_id;
DROP TABLE IF EXISTS project_member;
