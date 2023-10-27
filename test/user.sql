CREATE TABLE `user`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT,
    `country_code` tinyint unsigned NOT NULL DEFAULT '86' COMMENT '国家代码',
    `phone`        varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `password`     varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
    `salt`         varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
    `openid`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'openid',
    `status`       tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否禁用 0-禁用  1-正常',
    `created_at`   datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `phone_unique` (`phone`),
    KEY            `openid_index` (`openid`),
    KEY            `phone_status_index` (`phone`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';