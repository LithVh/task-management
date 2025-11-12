-- Revert Owner_id back to User_id
ALTER TABLE project RENAME COLUMN Owner_id TO User_id;
