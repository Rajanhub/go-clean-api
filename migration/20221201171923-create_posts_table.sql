-- +migrate Up
CREATE TABLE IF NOT EXISTS `posts` (
  `id` BINARY(16) NOT NULL, 
  `title` LONGTEXT,
  `description` LONGTEXT,
  `user_id`  BINARY(16) DEFAULT NULL,
  `created_at` DATETIME DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  KEY `fk_posts_user` (`user_id`),
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `posts`;