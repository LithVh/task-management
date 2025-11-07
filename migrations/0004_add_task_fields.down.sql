-- Rollback: Remove added fields from task table
ALTER TABLE task DROP COLUMN IF EXISTS completed;
ALTER TABLE task DROP COLUMN IF EXISTS due_date;
ALTER TABLE task DROP COLUMN IF EXISTS description;
