
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE "jobs" ADD COLUMN control_group_csv_path TEXT;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "jobs" DROP COLUMN control_group_csv_path;

