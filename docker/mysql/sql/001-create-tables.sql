DROP SCHEMA IF EXISTS `todo`;
CREATE SCHEMA `todo`;
USE `todo`;

---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF NOT EXISTS `users`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `created_at`       Datetime DEFAULT NULL,
  `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
