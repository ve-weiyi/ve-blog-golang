/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : blog-veweiyi

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 20/09/2024 15:55:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for album
-- ----------------------------
DROP TABLE IF EXISTS `album`;
CREATE TABLE `album` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_name` varchar(32) NOT NULL DEFAULT '' COMMENT '相册名',
  `album_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '相册描述',
  `album_cover` varchar(255) NOT NULL DEFAULT '' COMMENT '相册封面',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='相册';

-- ----------------------------
-- Records of album
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '分组id',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT 'api名称',
  `path` varchar(128) NOT NULL DEFAULT '' COMMENT 'api路径',
  `method` varchar(16) NOT NULL DEFAULT '' COMMENT 'api请求方法',
  `traceable` tinyint NOT NULL DEFAULT '0' COMMENT '是否追溯操作记录 0需要，1是',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1开，2关',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_path_method` (`path`,`method`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口';

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '作者',
  `category_id` int NOT NULL DEFAULT '0' COMMENT '文章分类',
  `article_cover` varchar(1024) NOT NULL DEFAULT '' COMMENT '文章缩略图',
  `article_title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `article_content` longtext NOT NULL COMMENT '内容',
  `article_type` tinyint NOT NULL DEFAULT '0' COMMENT '文章类型 1原创 2转载 3翻译',
  `original_url` varchar(255) NOT NULL DEFAULT '' COMMENT '原文链接',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶 0否 1是',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密 3评论可见',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发表时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章';

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (1, 1, 1, 'https://static.veweiyi.cn/blog/1/article/qinglong-20240920155237.jpg', '2024-09-20', '恭喜你成功运行博客！', 1, '', 0, 0, 1, 0, '2024-09-20 15:53:03', '2024-09-20 15:53:03');
COMMIT;

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `article_id` int NOT NULL DEFAULT '0' COMMENT '文章id',
  `tag_id` int NOT NULL DEFAULT '0' COMMENT '标签id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_article_tag_1` (`article_id`) USING BTREE,
  KEY `fk_article_tag_2` (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章-标签关联';

-- ----------------------------
-- Records of article_tag
-- ----------------------------
BEGIN;
INSERT INTO `article_tag` (`id`, `article_id`, `tag_id`) VALUES (1, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '页面id',
  `banner_name` varchar(32) NOT NULL DEFAULT '' COMMENT '页面名',
  `banner_label` varchar(32) NOT NULL DEFAULT '' COMMENT '页面标签',
  `banner_cover` varchar(255) NOT NULL DEFAULT '' COMMENT '页面封面',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面';

-- ----------------------------
-- Records of banner
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `category_name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章分类';

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` (`id`, `category_name`, `created_at`, `updated_at`) VALUES (1, '测试标签', '2024-09-20 15:53:03', '2024-09-20 15:53:03');
COMMIT;

-- ----------------------------
-- Table structure for chat_message
-- ----------------------------
DROP TABLE IF EXISTS `chat_message`;
CREATE TABLE `chat_message` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `chat_id` varchar(128) NOT NULL DEFAULT '' COMMENT '聊天id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `reply_msg_id` int NOT NULL DEFAULT '0' COMMENT '回复消息id',
  `content` varchar(4096) NOT NULL DEFAULT '' COMMENT '聊天内容',
  `ip_address` varchar(64) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip来源',
  `type` int NOT NULL DEFAULT '0' COMMENT '类型',
  `status` int NOT NULL DEFAULT '0' COMMENT '0正常 1撤回 2已编辑',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='聊天消息';

-- ----------------------------
-- Records of chat_message
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `content` varchar(1000) NOT NULL DEFAULT '' COMMENT '聊天内容',
  `ip_address` varchar(64) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip来源',
  `type` int NOT NULL DEFAULT '0' COMMENT '类型',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='聊天记录';

-- ----------------------------
-- Records of chat_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for chat_session
-- ----------------------------
DROP TABLE IF EXISTS `chat_session`;
CREATE TABLE `chat_session` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `chat_id` varchar(128) NOT NULL DEFAULT '' COMMENT '聊天id',
  `chat_title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `type` varchar(128) NOT NULL DEFAULT '' COMMENT '类型',
  `status` int NOT NULL DEFAULT '0' COMMENT '0正常 1删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='聊天会话';

-- ----------------------------
-- Records of chat_session
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `topic_id` int NOT NULL DEFAULT '0' COMMENT '主题id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父评论id',
  `session_id` int NOT NULL DEFAULT '0' COMMENT '会话id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '评论用户id',
  `reply_user_id` int NOT NULL DEFAULT '0' COMMENT '评论回复用户id',
  `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '评论点赞数量',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '评论类型 1.文章 2.友链 3.说说',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0.正常 1.已编辑 2.已删除',
  `is_review` tinyint NOT NULL DEFAULT '1' COMMENT '是否审核',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_comment_user` (`user_id`) USING BTREE,
  KEY `fk_comment_parent` (`parent_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='评论';

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `link_name` varchar(32) NOT NULL DEFAULT '' COMMENT '链接名',
  `link_avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '链接头像',
  `link_address` varchar(64) NOT NULL DEFAULT '' COMMENT '链接地址',
  `link_intro` varchar(100) NOT NULL DEFAULT '' COMMENT '链接介绍',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_friend_link_user` (`link_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='友链';

-- ----------------------------
-- Records of friend
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父id',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单标题',
  `path` varchar(64) NOT NULL DEFAULT '' COMMENT '路由路径',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '路由名称',
  `component` varchar(256) NOT NULL DEFAULT '' COMMENT '路由组件',
  `redirect` varchar(256) NOT NULL DEFAULT '' COMMENT '路由重定向',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '菜单类型',
  `rank` int NOT NULL DEFAULT '0' COMMENT '排序',
  `extra` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单元数据',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_path` (`path`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `nickname` varchar(64) DEFAULT '' COMMENT '用户昵称',
  `ip_address` varchar(255) DEFAULT '' COMMENT '操作ip',
  `ip_source` varchar(255) DEFAULT '' COMMENT '操作地址',
  `opt_module` varchar(32) DEFAULT '' COMMENT '操作模块',
  `opt_desc` varchar(255) DEFAULT '' COMMENT '操作描述',
  `request_url` varchar(255) DEFAULT '' COMMENT '请求地址',
  `request_method` varchar(32) DEFAULT '' COMMENT '请求方式',
  `request_header` varchar(1024) DEFAULT '' COMMENT '请求头参数',
  `request_data` varchar(4096) DEFAULT '' COMMENT '请求参数',
  `response_data` varchar(4096) DEFAULT '' COMMENT '返回数据',
  `response_status` int NOT NULL DEFAULT '0' COMMENT '响应状态码',
  `cost` varchar(32) NOT NULL DEFAULT '' COMMENT '耗时（ms）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='操作记录';

-- ----------------------------
-- Records of operation_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_id` int NOT NULL DEFAULT '0' COMMENT '相册id',
  `photo_name` varchar(32) NOT NULL DEFAULT '' COMMENT '照片名',
  `photo_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '照片描述',
  `photo_src` varchar(255) NOT NULL DEFAULT '' COMMENT '照片地址',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='照片';

-- ----------------------------
-- Records of photo
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for remark
-- ----------------------------
DROP TABLE IF EXISTS `remark`;
CREATE TABLE `remark` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `message_content` varchar(255) NOT NULL DEFAULT '' COMMENT '留言内容',
  `ip_address` varchar(64) NOT NULL DEFAULT '' COMMENT '用户ip',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT '用户地址',
  `is_review` tinyint NOT NULL DEFAULT '1' COMMENT '是否审核',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='留言';

-- ----------------------------
-- Records of remark
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父角色id',
  `role_domain` varchar(64) NOT NULL DEFAULT '0' COMMENT '角色域',
  `role_name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名',
  `role_comment` varchar(64) NOT NULL DEFAULT '' COMMENT '角色备注',
  `is_disable` tinyint NOT NULL DEFAULT '0' COMMENT '是否禁用  0否 1是',
  `is_default` tinyint NOT NULL DEFAULT '0' COMMENT '是否默认角色 0否 1是',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_api
-- ----------------------------
DROP TABLE IF EXISTS `role_api`;
CREATE TABLE `role_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `api_id` int NOT NULL DEFAULT '0' COMMENT '接口id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-api关联';

-- ----------------------------
-- Records of role_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-菜单关联';

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tag_name` varchar(32) NOT NULL DEFAULT '' COMMENT '标签名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_name` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='标签';

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
INSERT INTO `tag` (`id`, `tag_name`, `created_at`, `updated_at`) VALUES (1, '测试分类', '2024-09-20 15:53:03', '2024-09-20 15:53:03');
COMMIT;

-- ----------------------------
-- Table structure for talk
-- ----------------------------
DROP TABLE IF EXISTS `talk`;
CREATE TABLE `talk` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '说说id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `content` varchar(2048) NOT NULL DEFAULT '' COMMENT '说说内容',
  `images` varchar(2048) NOT NULL DEFAULT '' COMMENT '图片',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1.公开 2.私密',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='说说';

-- ----------------------------
-- Records of talk
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for upload_record
-- ----------------------------
DROP TABLE IF EXISTS `upload_record`;
CREATE TABLE `upload_record` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `label` varchar(128) NOT NULL DEFAULT '' COMMENT '标签',
  `file_name` varchar(64) NOT NULL DEFAULT '' COMMENT '文件名称',
  `file_size` int NOT NULL DEFAULT '0' COMMENT '文件大小',
  `file_md5` varchar(128) NOT NULL DEFAULT '' COMMENT '文件md5值',
  `file_url` varchar(256) NOT NULL DEFAULT '' COMMENT '上传路径',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='上传记录';

-- ----------------------------
-- Records of upload_record
-- ----------------------------
BEGIN;
INSERT INTO `upload_record` (`id`, `user_id`, `label`, `file_name`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (1, 1, 'article', 'qinglong.jpg', 195293, '40dc6ddf12acd53fcc27722f4ca3b7f4', 'https://static.veweiyi.cn/blog/1/article/qinglong-20240920155237.jpg', '2024-09-20 15:52:39', '2024-09-20 15:52:39');
COMMIT;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `nickname` varchar(64) NOT NULL COMMENT '用户昵称',
  `avatar` varchar(255) NOT NULL COMMENT '用户头像',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(64) DEFAULT '' COMMENT '手机号',
  `info` varchar(1024) NOT NULL COMMENT '用户信息',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: -1删除 0正常 1禁用',
  `login_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '注册方式',
  `ip_address` varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip 源',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录信息';

-- ----------------------------
-- Records of user_account
-- ----------------------------
BEGIN;
INSERT INTO `user_account` (`id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`, `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`) VALUES (1, 'admin@qq.com', '$2a$10$ZINovpDg.FxFQRj6nhKDLOH55k19RDViybnVVn5EGuKQAcqChRs1e', '管理员', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'admin@qq.com', '', '{\"intro\":\"23\",\"website\":\"3\"}', 0, 'email', '127.0.0.1', '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2024-09-03 11:00:21');
COMMIT;

-- ----------------------------
-- Table structure for user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `user_login_history`;
CREATE TABLE `user_login_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `login_type` varchar(64) NOT NULL DEFAULT '' COMMENT '登录类型',
  `agent` varchar(255) NOT NULL DEFAULT '' COMMENT '代理',
  `ip_address` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip host',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip 源',
  `login_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `logout_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登出时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uk_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录历史';

-- ----------------------------
-- Records of user_login_history
-- ----------------------------
BEGIN;
INSERT INTO `user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (1, 1, 'oauth', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-09-20 15:50:50', '1970-01-01 08:00:00', '2024-09-20 15:50:50', '2024-09-20 15:50:50');
COMMIT;

-- ----------------------------
-- Table structure for user_oauth
-- ----------------------------
DROP TABLE IF EXISTS `user_oauth`;
CREATE TABLE `user_oauth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `open_id` varchar(128) NOT NULL DEFAULT '' COMMENT '开发平台id，标识唯一用户',
  `platform` varchar(64) NOT NULL DEFAULT '' COMMENT '平台:手机号、邮箱、微信、飞书',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_oid_plat` (`open_id`,`platform`) USING BTREE,
  KEY `idx_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='第三方登录信息';

-- ----------------------------
-- Records of user_oauth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户-角色关联';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for visit_history
-- ----------------------------
DROP TABLE IF EXISTS `visit_history`;
CREATE TABLE `visit_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `date` varchar(10) NOT NULL DEFAULT '' COMMENT '日期',
  `views_count` int NOT NULL DEFAULT '0' COMMENT '访问量',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面访问数量';

-- ----------------------------
-- Records of visit_history
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for website_config
-- ----------------------------
DROP TABLE IF EXISTS `website_config`;
CREATE TABLE `website_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `key` varchar(32) NOT NULL DEFAULT '' COMMENT '关键词',
  `config` varchar(2048) NOT NULL DEFAULT '' COMMENT '配置信息',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='网站配置表';

-- ----------------------------
-- Records of website_config
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
