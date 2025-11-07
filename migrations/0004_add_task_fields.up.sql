-- Add missing fields to task table
ALTER TABLE task ADD COLUMN description TEXT;
ALTER TABLE task ADD COLUMN due_date TIMESTAMP;
ALTER TABLE task ADD COLUMN completed BOOLEAN DEFAULT false NOT NULL;
