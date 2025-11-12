-- Rename User_id to Owner_id in project table for clarity
ALTER TABLE project RENAME COLUMN User_id TO Owner_id;
