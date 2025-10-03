-- Modify "users" table
ALTER TABLE `users` ADD UNIQUE INDEX `username` (`username`);
-- Modify "projects" table
ALTER TABLE `projects` RENAME COLUMN `task_project` TO `user_projects`, DROP INDEX `task_project`, ADD INDEX `projects_users_projects` (`user_projects`), DROP FOREIGN KEY `projects_tasks_project`, ADD CONSTRAINT `projects_users_projects` FOREIGN KEY (`user_projects`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "tasks" table
ALTER TABLE `tasks` RENAME COLUMN `user_tasks` TO `project_tasks`, DROP INDEX `tasks_users_tasks`, ADD INDEX `tasks_projects_tasks` (`project_tasks`), DROP FOREIGN KEY `tasks_users_tasks`, ADD CONSTRAINT `tasks_projects_tasks` FOREIGN KEY (`project_tasks`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
