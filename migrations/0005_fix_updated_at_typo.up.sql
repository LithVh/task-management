-- Fix typo in task table column name
ALTER TABLE task RENAME COLUMN updated_at_at TO updated_at;
