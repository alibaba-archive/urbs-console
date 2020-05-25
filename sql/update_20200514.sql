CREATE TABLE IF NOT EXISTS `urbs_console`.`urbs_lock` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `expire_at` datetime(3) NOT NULL,
    `name` varchar(127) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;