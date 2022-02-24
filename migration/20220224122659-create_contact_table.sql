-- +migrate Up
CREATE TABLE `contact` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0: Male, 1: Female, 2: Secret',
    `is_disable` tinyint(4) NOT NULL DEFAULT '0',
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +migrate Down
DROP TABLE `contact`;
