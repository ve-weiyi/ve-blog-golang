
-- ============================================
-- 最小运行数据 (Minimal Seed Data)
-- ============================================

-- ----------------------------
-- Records of t_user (管理员账号，密码: root)
-- ----------------------------
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`, `register_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (1, 'root', 'root', '$2a$10$2FQhHyejaB998v1GBVUQYu8MiLPdrgnDP1ozltfa1.LsWD6.P.A/.', '超级管理员', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'root@qq.com', '', '{\"gender\":0,\"intro\":\"hello!\",\"website\":\"https://blog.veweiyi.cn\"}', 0, 'email', '127.0.0.1', '本地', '2024-07-10 16:24:50', '2024-07-10 16:24:50');

-- ----------------------------
-- Records of t_role
-- ----------------------------
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_default`, `status`, `created_at`, `updated_at`)
VALUES (1, 0, 'root', '超级管理员', '拥有所有权限', 0, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_user_role
-- ----------------------------
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (1, '1', 1);

-- ----------------------------
-- Records of t_menu (基础后台菜单)
-- ----------------------------
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `visible`, `status`, `extra`, `created_at`, `updated_at`)
VALUES (1, 0, '/system', 'System', '/src/layout/index', '/system/user', '', '系统管理', 'el-icon-setting', 6, '', '', 0, 0, 1, 0, '{\"title\":\"系统管理\",\"icon\":\"el-icon-setting\",\"rank\":6}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `visible`, `status`, `extra`, `created_at`, `updated_at`)
VALUES (2, 1, '/system/user', 'User', '/src/views/admin/system/user/User', '', '', '用户列表', '', 1, '', '', 0, 0, 1, 0, '{\"title\":\"用户列表\",\"rank\":1}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `visible`, `status`, `extra`, `created_at`, `updated_at`)
VALUES (3, 1, '/system/role', 'Role', '/src/views/admin/system/role/Role', '', '', '角色管理', '', 2, '', '', 0, 0, 1, 0, '{\"title\":\"角色管理\",\"rank\":2}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `visible`, `status`, `extra`, `created_at`, `updated_at`)
VALUES (4, 1, '/system/menu', 'Menu', '/src/views/admin/system/menu/Menu', '', '', '菜单管理', '', 3, '', '', 0, 0, 1, 0, '{\"title\":\"菜单管理\",\"rank\":3}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `visible`, `status`, `extra`, `created_at`, `updated_at`)
VALUES (5, 1, '/system/api', 'Api', '/src/views/admin/system/api/Api', '', '', '接口管理', '', 4, '', '', 0, 0, 1, 0, '{\"title\":\"接口管理\",\"rank\":4}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_role_menu
-- ----------------------------
INSERT INTO `t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1, 1, 1);
INSERT INTO `t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (2, 1, 2);
INSERT INTO `t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (3, 1, 3);
INSERT INTO `t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (4, 1, 4);
INSERT INTO `t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (5, 1, 5);

-- ----------------------------
-- Records of t_config
-- ----------------------------
INSERT INTO `t_config` (`id`, `key`, `config`, `created_at`, `updated_at`)
VALUES (1, 'website_config', '{\"admin_url\":\"https://admin.veweiyi.cn\",\"website_avatar\":\"\",\"website_author\":\"\",\"website_create_time\":\"\",\"website_info\":{},\"website_intro\":\"\",\"website_name\":\"\",\"website_notice\":\"\",\"website_record_no\":\"\",\"website_feature\":{\"is_chat_room\":0,\"is_comment_review\":0,\"is_email_notice\":0,\"is_message_review\":0,\"is_music_player\":0,\"is_reward\":0},\"tourist_avatar\":\"\",\"user_avatar\":\"\",\"reward_qr_code\":{},\"social_login_list\":[],\"social_url_list\":[],\"websocket_url\":\"\"}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `t_config` (`id`, `key`, `config`, `created_at`, `updated_at`)
VALUES (2, 'about_me', '{\"content\":\"welcome to my blog!\"}', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_category
-- ----------------------------
INSERT INTO `t_category` (`id`, `category_name`, `created_at`, `updated_at`)
VALUES (1, '默认分类', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_tag
-- ----------------------------
INSERT INTO `t_tag` (`id`, `tag_name`, `created_at`, `updated_at`)
VALUES (1, '默认标签', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_article (示例文章)
-- ----------------------------
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `view_count`, `created_at`, `updated_at`)
VALUES (1, '1', 1, '', 'Hello World', '欢迎使用 ve-blog！这是一篇示例文章。', 1, '', 1, 0, 1, 0, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_article_tag
-- ----------------------------
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`)
VALUES (1, 1, 1);

-- ----------------------------
-- Records of t_friend (示例友链)
-- ----------------------------
INSERT INTO `t_friend` (`id`, `link_name`, `link_avatar`, `link_address`, `link_intro`, `created_at`, `updated_at`)
VALUES (1, 've-blog', '', 'https://github.com/ve-weiyi', '一个全栈博客系统', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_talk (示例说说)
-- ----------------------------
INSERT INTO `t_talk` (`id`, `user_id`, `content`, `images`, `is_top`, `status`, `like_count`, `created_at`, `updated_at`)
VALUES (1, '1', '博客运行成功！', '', 1, 1, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Records of t_notify_template (示例通知模板)
-- ----------------------------
INSERT INTO `t_notify_template` (`id`, `code`, `channel`, `scene`, `title`, `content`, `enabled`, `created_at`, `updated_at`)
VALUES (1, 'SMS_LOGIN', 'sms', 'login', '登录验证码', '您的验证码是：{code}，有效期 {expire} 分钟。', 1, '2026-01-01 00:00:00', '2026-01-01 00:00:00');

INSERT INTO `t_notify_template` (`id`, `code`, `channel`, `scene`, `title`, `content`, `enabled`, `created_at`, `updated_at`)
VALUES (2, 'EMAIL_REGISTER', 'email', 'register', '注册确认', '感谢您注册 {site_name}，请点击链接完成验证：{link}', 1, '2026-01-01 00:00:00', '2026-01-01 00:00:00');

INSERT INTO `t_notify_template` (`id`, `code`, `channel`, `scene`, `title`, `content`, `enabled`, `created_at`, `updated_at`)
VALUES (3, 'INBOX_NOTIFY', 'inbox', 'notify', '系统通知', '{title}\n{content}', 1, '2026-01-01 00:00:00', '2026-01-01 00:00:00');

-- ----------------------------
-- Records of t_notify_message (示例通知消息)
-- ----------------------------
INSERT INTO `t_notify_message` (`id`, `title`, `content`, `category`, `level`, `target_type`, `target_ids`, `status`, `published_at`, `published_by`, `created_at`, `updated_at`)
VALUES (1, '系统上线公告', 've-blog 博客系统正式上线运行！', 'system', 'info', 'all', '', 'published', '2026-05-10 16:30:00', 'root', '2026-05-10 16:24:50', '2026-05-10 16:30:00');

INSERT INTO `t_notify_message` (`id`, `title`, `content`, `category`, `level`, `target_type`, `target_ids`, `status`, `published_at`, `published_by`, `created_at`, `updated_at`)
VALUES (2, '功能更新预告', '即将推出评论通知和消息推送功能。', 'update', 'warning', 'all', '', 'draft', NULL, '', '2026-05-15 10:00:00', '2026-05-15 10:00:00');

INSERT INTO `t_notify_message` (`id`, `title`, `content`, `category`, `level`, `target_type`, `target_ids`, `status`, `published_at`, `published_by`, `created_at`, `updated_at`)
VALUES (3, '紧急维护通知', '6月1日凌晨2:00-4:00服务器维护升级。', 'maintenance', 'error', 'all', '', 'published', '2026-05-28 12:00:00', 'root', '2026-05-28 12:00:00', '2026-05-28 12:00:00');

INSERT INTO `t_notify_message` (`id`, `title`, `content`, `category`, `level`, `target_type`, `target_ids`, `status`, `published_at`, `published_by`, `created_at`, `updated_at`)
VALUES (4, '重要提醒', '请尽快完善个人资料信息。', 'remind', 'info', 'user_ids', 'root', 'published', '2026-06-01 09:00:00', 'root', '2026-06-01 09:00:00', '2026-06-01 09:00:00');

-- ----------------------------
-- Records of t_notify_record (示例投递记录)
-- ----------------------------
INSERT INTO `t_notify_record` (`id`, `message_id`, `channel`, `recipient`, `template_code`, `content`, `status`, `biz_id`, `error_msg`, `read_at`, `sent_at`, `created_at`)
VALUES (1, 1, 'inbox', 'root', '', 've-blog 博客系统正式上线运行！', 'read', '', '', '2026-05-10 18:00:00', NULL, '2026-05-10 16:30:00');

INSERT INTO `t_notify_record` (`id`, `message_id`, `channel`, `recipient`, `template_code`, `content`, `status`, `biz_id`, `error_msg`, `read_at`, `sent_at`, `created_at`)
VALUES (2, 3, 'inbox', 'root', '', '6月1日凌晨2:00-4:00服务器维护升级。', 'unread', '', '', NULL, NULL, '2026-05-28 12:00:00');

INSERT INTO `t_notify_record` (`id`, `message_id`, `channel`, `recipient`, `template_code`, `content`, `status`, `biz_id`, `error_msg`, `read_at`, `sent_at`, `created_at`)
VALUES (3, 3, 'sms', '13800138000', 'SMS_LOGIN', '6月1日凌晨2:00-4:00服务器维护升级。', 'sent', 'sms_biz_001', '', NULL, '2026-05-28 12:00:01', '2026-05-28 12:00:00');

INSERT INTO `t_notify_record` (`id`, `message_id`, `channel`, `recipient`, `template_code`, `content`, `status`, `biz_id`, `error_msg`, `read_at`, `sent_at`, `created_at`)
VALUES (4, 3, 'email', 'admin@example.com', 'EMAIL_REGISTER', '6月1日凌晨2:00-4:00服务器维护升级。', 'failed', 'email_biz_001', 'mailbox unavailable', NULL, NULL, '2026-05-28 12:00:00');

INSERT INTO `t_notify_record` (`id`, `message_id`, `channel`, `recipient`, `template_code`, `content`, `status`, `biz_id`, `error_msg`, `read_at`, `sent_at`, `created_at`)
VALUES (5, 4, 'inbox', 'root', '', '请尽快完善个人资料信息。', 'unread', '', '', NULL, NULL, '2026-06-01 09:00:00');


BEGIN;
INSERT INTO `t_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (1, 'website_config', '{\"admin_url\":\"https://admin.veweiyi.cn\",\"websocket_url\":\"wss://blog.veweiyi.cn/api/websocket\",\"tourist_avatar\":\"https://static.veweiyi.cn/blog/website/logo-20260108110231.jpg\",\"user_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175820.jpg\",\"website_feature\":{\"is_chat_room\":1,\"is_ai_assistant\":1,\"is_music_player\":1,\"is_comment_review\":0,\"is_email_notice\":1,\"is_message_review\":0,\"is_reward\":1},\"website_info\":{\"website_author\":\"与梦\",\"website_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-100-20251230185707.jpg\",\"website_create_time\":\"2022-01-17\",\"website_intro\":\"你能做的，岂止如此。\",\"website_name\":\"与梦\",\"website_notice\":\"\\u003cp\\u003e网站搭建问题请联系QQ \\u003ca href=\\\"tencent://message/?uin=791422171\\u0026Site=\\u0026Menu=yes\\\" target=\\\"_blank\\\" style=\\\"color: blue;\\\"\\u003e791422171\\u003c/a\\u003e。\\u003c/p\\u003e\\n\\u003cp\\u003e后台管理系统请访问 \\u003ca href=\\\"https://admin.veweiyi.cn\\\" target=\\\"_blank\\\" style=\\\"color: blue;\\\"\\u003ehttps://admin.veweiyi.cn\\u003c/a\\u003e\\u003c/p\\u003e\\n\\u003cp\\u003egithub项目地址 \\u003ca href=\\\"https://github.com/ve-weiyi/ve-blog-golang\\\" target=\\\"_blank\\\" style=\\\"color: blue;\\\"\\u003ehttps://github.com/ve-weiyi/ve-blog-golang\\u003c/a\\u003e\\u003c/p\\u003e\",\"website_record_no\":\"桂ICP备2023013735号-1\"},\"reward_qr_code\":{\"alipay_qr_code\":\"https://static.veweiyi.cn/blog/website/v-20260108110058.png\",\"weixin_qr_code\":\"https://static.veweiyi.cn/blog/website/809B3510BE3DB31CBDF93F689748177E-20260108105949.png\"},\"social_login_list\":[{\"name\":\"github\",\"platform\":\"github\",\"authorize_url\":\"\",\"enabled\":true},{\"name\":\"gitee\",\"platform\":\"gitee\",\"authorize_url\":\"\",\"enabled\":true},{\"name\":\"qq\",\"platform\":\"qq\",\"authorize_url\":\"\",\"enabled\":true},{\"name\":\"微信\",\"platform\":\"wechat\",\"authorize_url\":\"\",\"enabled\":true},{\"name\":\"哔哩哔哩\",\"platform\":\"bilibili\",\"authorize_url\":\"\",\"enabled\":false},{\"name\":\"微博\",\"platform\":\"weibo\",\"authorize_url\":\"\",\"enabled\":true},{\"name\":\"飞书\",\"platform\":\"feishu\",\"authorize_url\":\"\",\"enabled\":false}],\"social_url_list\":[{\"name\":\"github\",\"platform\":\"github\",\"link_url\":\"https://github.com/ve-weiyi\",\"enabled\":true},{\"name\":\"gitee\",\"platform\":\"gitee\",\"link_url\":\"https://gitee.com/wy791422171\",\"enabled\":true},{\"name\":\"qq\",\"platform\":\"qq\",\"link_url\":\"http://wpa.qq.com/msgrd?v=3\\u0026uin=791422171\\u0026site=qq\\u0026menu=yes\",\"enabled\":true},{\"name\":\"微信\",\"platform\":\"wechat\",\"link_url\":\"wy791422171\",\"enabled\":true}]}', '2021-08-09 19:37:30', '2026-01-20 11:20:16');
INSERT INTO `t_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (2, 'about_me', '{\"content\":\"welcome to my blog!\"}', '2024-11-15 17:57:20', '2024-11-15 17:57:20');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
