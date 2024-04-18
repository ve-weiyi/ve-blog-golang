DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `account`
(
    `id`            int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`      varchar(64)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`      varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
    `status`        tinyint      NOT NULL DEFAULT '0' COMMENT '状态: -1删除 0正常 1禁用',
    `register_type` varchar(64)  NOT NULL DEFAULT '' COMMENT '注册方式',
    `ip_address`    varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip',
    `ip_source`     varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip 源',
    `created_at`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_username` (`username`,`register_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录信息';
