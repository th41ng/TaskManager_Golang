-- Create "users" table
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `password` bigint NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "tasks" table
CREATE TABLE `tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `done` bool NOT NULL DEFAULT 0,
  `user_tasks` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `tasks_users_tasks` (`user_tasks`),
  CONSTRAINT `tasks_users_tasks` FOREIGN KEY (`user_tasks`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
