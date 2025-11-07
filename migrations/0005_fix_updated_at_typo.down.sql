-- Rollback: Restore typo
ALTER TABLE task RENAME COLUMN updated_at TO updated_at_at;
