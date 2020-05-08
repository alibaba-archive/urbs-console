-- Keywords and Reserved Words https://dev.mysql.com/doc/refman/5.7/en/keywords.html name status channel value values group user
CREATE DATABASE IF NOT EXISTS `urbs_console` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `urbs_console`.`operation_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `operator` varchar(63) NOT NULL,
  `object` varchar(189) NOT NULL,
  `action` varchar(63) NOT NULL,
  `content` TEXT NOT NULL,
  `description` varchar(8190) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_object` (`object`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `urbs_console`.`urbs_ac_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `uid` varchar(63) NOT NULL,
  `name` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uid` (`uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `urbs_console`.`urbs_ac_acl` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `subject` varchar(63) NOT NULL,
  `object` varchar(256) NOT NULL,
  `permission` varchar(63) NOT NULL,
  PRIMARY KEY (`id`),
  KEY KEY `idx_subject` (`subject`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;