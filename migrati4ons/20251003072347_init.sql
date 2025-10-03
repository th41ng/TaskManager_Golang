-- Create "users" table
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "projects" table
CREATE TABLE `projects` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `user_projects` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `projects_users_projects` (`user_projects`),
  CONSTRAINT `projects_users_projects` FOREIGN KEY (`user_projects`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "tasks" table
CREATE TABLE `tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `done` bool NOT NULL DEFAULT 0,
  `project_tasks` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `tasks_projects_tasks` (`project_tasks`),
  CONSTRAINT `tasks_projects_tasks` FOREIGN KEY (`project_tasks`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
