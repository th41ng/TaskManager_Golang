-- Modify "tasks" table
ALTER TABLE `tasks` ADD COLUMN `priority` bigint NOT NULL AFTER `done`;
