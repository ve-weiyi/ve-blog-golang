-- ============================================
-- migrate.sql
-- 从 blog-struct.sql 当前结构迁移到代码模型
-- 执行前请确认已备份数据库！
-- ============================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ============================================
-- 1. 修复 t_user 排序规则
-- ============================================
ALTER TABLE `t_user` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- ============================================
-- 2. 修复 t_login_log 排序规则
-- ============================================
ALTER TABLE `t_login_log` CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- ============================================
-- 3. 重建 t_daily_stats（旧结构 pivot 到新结构）
-- ============================================
DROP TABLE IF EXISTS `t_daily_stats_new`;
CREATE TABLE `t_daily_stats_new`
(
    `id`             BIGINT      NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `date`           VARCHAR(10) NOT NULL DEFAULT '' COMMENT '统计日期 YYYY-MM-DD',
    `new_users`      BIGINT      NOT NULL DEFAULT 0 COMMENT '当日新增用户数',
    `total_users`    BIGINT      NOT NULL DEFAULT 0 COMMENT '累计用户数',
    `active_users`   BIGINT      NOT NULL DEFAULT 0 COMMENT '当日活跃用户数',
    `uv_count`       BIGINT      NOT NULL DEFAULT 0 COMMENT '当日独立访客数(UV)',
    `pv_count`       BIGINT      NOT NULL DEFAULT 0 COMMENT '当日页面浏览数(PV)',
    `total_uv_count` BIGINT      NOT NULL DEFAULT 0 COMMENT '累计访客数',
    `total_pv_count` BIGINT      NOT NULL DEFAULT 0 COMMENT '累计浏览量',
    `created_at`     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_date` (`date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='每日统计表';

-- pivot 旧数据：visit_type 1→uv_count, 2→pv_count，计算累计值
INSERT INTO `t_daily_stats_new` (`date`, `uv_count`, `pv_count`, `total_uv_count`, `total_pv_count`, `created_at`, `updated_at`)
SELECT
    `date`,
    `uv_count`,
    `pv_count`,
    SUM(`uv_count`) OVER (ORDER BY `date`)  AS `total_uv_count`,
    SUM(`pv_count`) OVER (ORDER BY `date`)  AS `total_pv_count`,
    `created_at`,
    `updated_at`
FROM (
    SELECT
        `date`,
        MAX(CASE WHEN `visit_type` = 1 THEN `view_count` ELSE 0 END) AS uv_count,
        MAX(CASE WHEN `visit_type` = 2 THEN `view_count` ELSE 0 END) AS pv_count,
        MIN(`created_at`) AS created_at,
        MAX(`updated_at`) AS updated_at
    FROM `t_daily_stats`
    GROUP BY `date`
) AS `pivoted`;

DROP TABLE IF EXISTS `t_daily_stats`;
RENAME TABLE `t_daily_stats_new` TO `t_daily_stats`;

-- ============================================
-- 4. 种子数据（如已存在则跳过）
-- ============================================

-- t_config
INSERT IGNORE INTO `t_config` (`id`, `key`, `config`, `created_at`, `updated_at`)
VALUES
(1, 'website_config', '{"admin_url":"https://admin.veweiyi.cn","website_avatar":"","website_author":"","website_create_time":"","website_info":{},"website_intro":"","website_name":"","website_notice":"","website_record_no":"","website_feature":{"is_chat_room":0,"is_comment_review":0,"is_email_notice":0,"is_message_review":0,"is_music_player":0,"is_reward":0},"tourist_avatar":"","user_avatar":"","reward_qr_code":{},"social_login_list":[],"social_url_list":[],"websocket_url":""}', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(2, 'about_me', '{"content":"welcome to my blog!"}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- t_notify_template
INSERT IGNORE INTO `t_notify_template` (`id`, `code`, `channel`, `scene`, `title`, `content`, `enabled`, `created_at`, `updated_at`)
VALUES
(1, 'SMS_LOGIN', 'sms', 'login', '登录验证码', '您的验证码是：{code}，有效期 {expire} 分钟。', 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(2, 'EMAIL_REGISTER', 'email', 'register', '注册确认', '感谢您注册 {site_name}，请点击链接完成验证：{link}', 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(3, 'INBOX_NOTIFY', 'inbox', 'notify', '系统通知', '{title}\n{content}', 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

SET FOREIGN_KEY_CHECKS = 1;
