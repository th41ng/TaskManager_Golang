-- Create "tasks" table
CREATE TABLE `tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `done` bool NOT NULL DEFAULT 0,
  `priority` bigint NOT NULL,
  `project_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
