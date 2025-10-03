-- Modify "projects" table
ALTER TABLE `projects` ADD COLUMN `name` varchar(255) NOT NULL AFTER `id`, ADD COLUMN `task_project` bigint NULL, ADD UNIQUE INDEX `task_project` (`task_project`), ADD CONSTRAINT `projects_tasks_project` FOREIGN KEY (`task_project`) REFERENCES `tasks` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
