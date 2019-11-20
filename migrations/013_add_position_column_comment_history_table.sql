-- +migrate Up

ALTER TABLE leave_comment_history ADD COLUMN position TEXT DEFAULT 'NA' NOT NULL;

-- +migrate Down
ALTER TABLE leave_comment_history DROP COLUMN IF EXISTS position;