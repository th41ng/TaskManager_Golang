-- Modify "users" table
ALTER TABLE `users` RENAME COLUMN `user_name` TO `username`, MODIFY COLUMN `password` varchar(255) NOT NULL;
