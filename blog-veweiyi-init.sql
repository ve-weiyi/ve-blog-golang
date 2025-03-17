-- ----------------------------
-- Records of t_user
-- ----------------------------
BEGIN;
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_disable`, `is_default`,
                      `created_at`, `updated_at`)
VALUES (1, 0, 'super-admin', '超级管理员', '', 0, 0, '2021-03-22 14:10:21', '2024-11-15 17:44:02');

INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`, `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`) VALUES (1, '1', 'admin@qq.com', '$2a$10$ZINovpDg.FxFQRj6nhKDLOH55k19RDViybnVVn5EGuKQAcqChRs1e', '管理员', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'admin@qq.com', '', '{\"intro\":\"23\",\"website\":\"3\"}', 0, 'email', '127.0.0.1', '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2024-10-25 14:35:59');

INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`) VALUES (1, '1', 1);
COMMIT;


-- ----------------------------
-- Records of t_menu
-- ----------------------------
BEGIN;
INSERT INTO `blog-veweiyi`.`t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`, `updated_at`) VALUES (56, 0, '/system', '', '/src/layout/index', '/system/user', 0, '系统管理', 'el-icon-setting', 6, '', 'null', 0, 0, 0, 0, '{\"title\":\"系统管理\",\"icon\":\"el-icon-setting\",\"rank\":6,\"params\":\"null\"}', '2024-11-21 17:58:47', '2024-11-21 17:58:47');
INSERT INTO `blog-veweiyi`.`t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`, `updated_at`) VALUES (57, 56, '/system/user', 'User', '/src/views/admin/system/user/User', '', 1, '用户列表', '', 1, '', 'null', 0, 0, 0, 0, '{\"type\":1,\"title\":\"用户列表\",\"rank\":1,\"params\":\"null\"}', '2024-11-21 17:58:47', '2024-11-21 17:58:47');
INSERT INTO `blog-veweiyi`.`t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`, `updated_at`) VALUES (58, 56, '/system/role', 'Role', '/src/views/admin/system/role/Role', '', 1, '角色管理', '', 2, '', 'null', 0, 0, 0, 0, '{\"type\":1,\"title\":\"角色管理\",\"rank\":2,\"params\":\"null\"}', '2024-11-21 17:58:47', '2024-11-21 17:58:47');
INSERT INTO `blog-veweiyi`.`t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`, `updated_at`) VALUES (59, 56, '/system/menu', 'Menu', '/src/views/admin/system/menu/Menu', '', 1, '菜单管理', '', 3, '', 'null', 0, 0, 0, 0, '{\"type\":1,\"title\":\"菜单管理\",\"rank\":3,\"params\":\"null\"}', '2024-11-21 17:58:47', '2024-11-21 17:58:47');
INSERT INTO `blog-veweiyi`.`t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`, `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`, `updated_at`) VALUES (60, 56, '/system/api', 'Api', '/src/views/admin/system/api/Api', '', 1, '接口管理', '', 4, '', 'null', 0, 0, 0, 0, '{\"type\":1,\"title\":\"接口管理\",\"rank\":4,\"params\":\"null\"}', '2024-11-21 17:58:47', '2024-11-21 17:58:47');


INSERT INTO `blog-veweiyi`.`t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (56, 1, 56);
INSERT INTO `blog-veweiyi`.`t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (57, 1, 57);
INSERT INTO `blog-veweiyi`.`t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (58, 1, 58);
INSERT INTO `blog-veweiyi`.`t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (59, 1, 59);
INSERT INTO `blog-veweiyi`.`t_role_menu` (`id`, `role_id`, `menu_id`) VALUES (60, 1, 60);
COMMIT;



-- ----------------------------
-- Records of t_article
-- ----------------------------
BEGIN;
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (1, '1', 1, 'https://static.veweiyi.cn/blog/article/qinglong-20241115174624.jpg', '测试文章', '恭喜你成功运行了博客！', 1, '', 1, 2, 1, 1, '2024-11-15 17:46:29', '2024-11-15 23:49:28');
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (2, '1', 1, 'https://static.veweiyi.cn/blog/article/zhuqu-20241115182343.jpg', '草稿文章', '这是一篇草稿文章！', 1, '', 2, 2, 2, 1, '2024-11-15 18:22:24', '2024-11-15 23:49:30');
COMMIT;

-- ----------------------------
-- Records of t_article_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`) VALUES (6, 1, 1);
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`) VALUES (7, 2, 1);
COMMIT;

-- ----------------------------
-- Records of t_category
-- ----------------------------
BEGIN;
INSERT INTO `t_category` (`id`, `category_name`, `created_at`, `updated_at`) VALUES (1, '测试分类', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
COMMIT;

BEGIN;
INSERT INTO `t_tag` (`id`, `tag_name`, `created_at`, `updated_at`) VALUES (1, '测试标签', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
COMMIT;

-- ----------------------------
-- Records of t_friend
-- ----------------------------
BEGIN;
INSERT INTO `t_friend` (`id`, `link_name`, `link_avatar`, `link_address`, `link_intro`, `created_at`, `updated_at`) VALUES (1, '与梦', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'https://blog.veweiyi.cn', '你能做的，岂止如此。', '2024-11-16 00:43:12', '2024-11-16 00:43:37');
COMMIT;

----------------------
-- Records of t_talk
-- ----------------------------
BEGIN;
INSERT INTO `t_talk` (`id`, `user_id`, `content`, `images`, `is_top`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (1, '1', '测试说说<img src=\"https://static.veweiyi.cn/emoji/qq/14@2x.gif\" width=\"24\" height=\"24\" alt=\"[微笑]\" style=\"margin: 0 1px;display: inline;vertical-align: text-bottom\">', 'null', 1, 1, 0, '2024-11-16 00:33:43', '2024-11-16 00:39:15');
COMMIT;


BEGIN;
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (1, 'website_config', '{\"admin_url\":\"\",\"alipay_qr_code\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/17f234dc487c1bb5bbb732869be0eb53.jpg\",\"gitee\":\"https://gitee.com/wy791422171\",\"github\":\"https://github.com/ve-weiyi\",\"is_chat_room\":1,\"is_comment_review\":1,\"is_email_notice\":1,\"is_message_review\":0,\"is_music_player\":1,\"is_reward\":0,\"qq\":\"791422171\",\"social_login_list\":[\"qq\",\"github\",\"gitee\",\"feishu\",\"weibo\"],\"social_url_list\":[\"qq\",\"github\",\"gitee\"],\"tourist_avatar\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/5bfb96809bee5ba80a36811f0bf1d1ea.gif\",\"user_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175820.jpg\",\"website_author\":\"与梦\",\"website_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175746.jpg\",\"website_create_time\":\"2022-01-17\",\"website_intro\":\"你能做的，岂止如此。\",\"website_name\":\"与梦\",\"website_notice\":\"网站搭建问题请联系QQ 791422171。\",\"website_record_no\":\"桂ICP备2023013735号-1\",\"websocket_url\":\"wss://veweiyi.cn/api/websocket\",\"weixin_qr_code\":\"\"}', '2021-08-09 19:37:30', '2024-11-16 00:44:08');
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (2, 'about_me', '{\"content\":\"welcome to my blog!\"}', '2024-11-15 17:57:20', '2024-11-15 17:57:20');
COMMIT;
