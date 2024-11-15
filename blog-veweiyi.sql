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

 Date: 16/11/2024 00:55:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_album
-- ----------------------------
DROP TABLE IF EXISTS `t_album`;
CREATE TABLE `t_album` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_name` varchar(32) NOT NULL DEFAULT '' COMMENT '相册名',
  `album_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '相册描述',
  `album_cover` varchar(255) NOT NULL DEFAULT '' COMMENT '相册封面',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='相册';

-- ----------------------------
-- Records of t_album
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_api
-- ----------------------------
DROP TABLE IF EXISTS `t_api`;
CREATE TABLE `t_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '分组id',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT 'api名称',
  `path` varchar(128) NOT NULL DEFAULT '' COMMENT 'api路径',
  `method` varchar(16) NOT NULL DEFAULT '' COMMENT 'api请求方法',
  `traceable` tinyint NOT NULL DEFAULT '0' COMMENT '是否追溯操作记录 0需要，1是',
  `is_disable` tinyint NOT NULL DEFAULT '0' COMMENT '是否禁用 0否 1是',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_path_method` (`path`,`method`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口';

-- ----------------------------
-- Records of t_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_article
-- ----------------------------
DROP TABLE IF EXISTS `t_article`;
CREATE TABLE `t_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '作者',
  `category_id` int NOT NULL DEFAULT '0' COMMENT '文章分类',
  `article_cover` varchar(1024) NOT NULL DEFAULT '' COMMENT '文章缩略图',
  `article_title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `article_content` longtext NOT NULL COMMENT '内容',
  `article_type` tinyint NOT NULL DEFAULT '0' COMMENT '文章类型 1原创 2转载 3翻译',
  `original_url` varchar(255) NOT NULL DEFAULT '' COMMENT '原文链接',
  `is_top` tinyint NOT NULL DEFAULT '2' COMMENT '是否置顶 0否 1是',
  `is_delete` tinyint NOT NULL DEFAULT '2' COMMENT '是否删除  0否 1是',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密 3评论可见',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发表时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章';

-- ----------------------------
-- Records of t_article
-- ----------------------------
BEGIN;
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (1, '1', 1, 'https://static.veweiyi.cn/blog/article/qinglong-20241115174624.jpg', '测试文章', '恭喜你成功运行了博客！', 1, '', 1, 2, 1, 1, '2024-11-15 17:46:29', '2024-11-15 23:49:28');
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`, `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (2, '1', 1, 'https://static.veweiyi.cn/blog/article/zhuqu-20241115182343.jpg', '草稿文章', '这是一篇草稿文章！', 1, '', 2, 2, 2, 1, '2024-11-15 18:22:24', '2024-11-15 23:49:30');
COMMIT;

-- ----------------------------
-- Table structure for t_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_article_tag`;
CREATE TABLE `t_article_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `article_id` int NOT NULL DEFAULT '0' COMMENT '文章id',
  `tag_id` int NOT NULL DEFAULT '0' COMMENT '标签id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_article_tag_1` (`article_id`) USING BTREE,
  KEY `fk_article_tag_2` (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章-标签关联';

-- ----------------------------
-- Records of t_article_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`) VALUES (6, 1, 1);
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`) VALUES (7, 2, 1);
COMMIT;

-- ----------------------------
-- Table structure for t_banner
-- ----------------------------
DROP TABLE IF EXISTS `t_banner`;
CREATE TABLE `t_banner` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '页面id',
  `banner_name` varchar(32) NOT NULL DEFAULT '' COMMENT '页面名',
  `banner_label` varchar(32) NOT NULL DEFAULT '' COMMENT '页面标签',
  `banner_cover` varchar(255) NOT NULL DEFAULT '' COMMENT '页面封面',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面';

-- ----------------------------
-- Records of t_banner
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_category
-- ----------------------------
DROP TABLE IF EXISTS `t_category`;
CREATE TABLE `t_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `category_name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章分类';

-- ----------------------------
-- Records of t_category
-- ----------------------------
BEGIN;
INSERT INTO `t_category` (`id`, `category_name`, `created_at`, `updated_at`) VALUES (1, '测试分类', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
COMMIT;

-- ----------------------------
-- Table structure for t_chat_message
-- ----------------------------
DROP TABLE IF EXISTS `t_chat_message`;
CREATE TABLE `t_chat_message` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `device_id` varchar(64) NOT NULL DEFAULT '' COMMENT '设备id',
  `topic_id` varchar(255) NOT NULL DEFAULT '' COMMENT '主题id,表示一个群会话',
  `reply_msg_id` varchar(255) NOT NULL DEFAULT '' COMMENT '回复消息id，at消息',
  `reply_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '回复用户列表,at用户列表',
  `chat_content` varchar(4096) NOT NULL DEFAULT '' COMMENT '聊天内容',
  `ip_address` varchar(64) NOT NULL DEFAULT '' COMMENT '用户ip 127.0.0.1',
  `ip_source` varchar(128) NOT NULL DEFAULT '' COMMENT '用户地址 广东省深圳市',
  `type` varchar(64) NOT NULL DEFAULT '' COMMENT '类型:chatgpt chatroom',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态:0正常 1编辑 2撤回 3删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='聊天消息';

-- ----------------------------
-- Records of t_chat_message
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '评论用户id',
  `topic_id` int NOT NULL DEFAULT '0' COMMENT '主题id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父评论id',
  `reply_msg_id` int NOT NULL DEFAULT '0' COMMENT '回复评论id',
  `reply_user_id` varchar(255) NOT NULL COMMENT '评论回复用户id',
  `comment_content` text NOT NULL COMMENT '评论内容',
  `ip_address` varchar(64) NOT NULL COMMENT 'ip地址 127.0.01',
  `ip_source` varchar(64) NOT NULL COMMENT 'ip来源 广东省',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '评论类型 1.文章 2.友链 3.说说',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0.正常 1.已编辑 2.已删除',
  `is_review` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核通过',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '评论点赞数量',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`,`reply_user_id`) USING BTREE,
  KEY `fk_comment_user` (`user_id`) USING BTREE,
  KEY `fk_comment_parent` (`parent_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='评论';

-- ----------------------------
-- Records of t_comment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_file_folder
-- ----------------------------
DROP TABLE IF EXISTS `t_file_folder`;
CREATE TABLE `t_file_folder` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `file_path` varchar(128) NOT NULL DEFAULT '' COMMENT '文件路径',
  `folder_name` varchar(128) NOT NULL DEFAULT '' COMMENT '文件夹名称',
  `folder_desc` varchar(128) NOT NULL DEFAULT '' COMMENT '文件夹描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_path` (`file_path`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件夹记录';

-- ----------------------------
-- Records of t_file_folder
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_file_upload
-- ----------------------------
DROP TABLE IF EXISTS `t_file_upload`;
CREATE TABLE `t_file_upload` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `file_path` varchar(128) NOT NULL DEFAULT '' COMMENT '文件路径',
  `file_name` varchar(128) NOT NULL DEFAULT '' COMMENT '文件名称',
  `file_type` varchar(128) NOT NULL DEFAULT '' COMMENT '文件类型',
  `file_size` int NOT NULL DEFAULT '0' COMMENT '文件大小',
  `file_md5` varchar(128) NOT NULL DEFAULT '' COMMENT '文件md5值',
  `file_url` varchar(256) NOT NULL DEFAULT '' COMMENT '上传路径',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_uid` (`user_id`) USING BTREE,
  KEY `idx_path` (`file_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='上传记录';

-- ----------------------------
-- Records of t_file_upload
-- ----------------------------
BEGIN;
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (1, '1', '/article', 'qinglong.jpg', '.jpg', 195293, '40dc6ddf12acd53fcc27722f4ca3b7f4', 'https://static.veweiyi.cn/blog/article/qinglong-20241115174624.jpg', '2024-11-15 17:46:26', '2024-11-15 17:46:26');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (2, '1', '/website', 'tiger.jpg', '.jpg', 57396, '142cdc069d224eec0092ae9ec63eebb8', 'https://static.veweiyi.cn/blog/website/tiger-20241115175746.jpg', '2024-11-15 17:57:47', '2024-11-15 17:57:47');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (3, '1', '/website', 'logo.jpg', '.jpg', 42497, 'f121d135f39f03e48da5fe5e8ced5b0a', 'https://static.veweiyi.cn/blog/website/logo-20241115175805.jpg', '2024-11-15 17:58:06', '2024-11-15 17:58:06');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (4, '1', '/website', 'tiger.jpg', '.jpg', 57396, '142cdc069d224eec0092ae9ec63eebb8', 'https://static.veweiyi.cn/blog/website/tiger-20241115175820.jpg', '2024-11-15 17:58:21', '2024-11-15 17:58:21');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`, `file_url`, `created_at`, `updated_at`) VALUES (5, '1', '/article', 'zhuqu.jpg', '.jpg', 195693, '970c74a565abccbc46b2cf5a333ffe3f', 'https://static.veweiyi.cn/blog/article/zhuqu-20241115182343.jpg', '2024-11-15 18:23:44', '2024-11-15 18:23:44');
COMMIT;

-- ----------------------------
-- Table structure for t_friend
-- ----------------------------
DROP TABLE IF EXISTS `t_friend`;
CREATE TABLE `t_friend` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `link_name` varchar(32) NOT NULL DEFAULT '' COMMENT '链接名',
  `link_avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '链接头像',
  `link_address` varchar(64) NOT NULL DEFAULT '' COMMENT '链接地址',
  `link_intro` varchar(100) NOT NULL DEFAULT '' COMMENT '链接介绍',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_friend_link_user` (`link_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='友链';

-- ----------------------------
-- Records of t_friend
-- ----------------------------
BEGIN;
INSERT INTO `t_friend` (`id`, `link_name`, `link_avatar`, `link_address`, `link_intro`, `created_at`, `updated_at`) VALUES (1, '与梦', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'https://blog.veweiyi.cn', '你能做的，岂止如此。', '2024-11-16 00:43:12', '2024-11-16 00:43:37');
COMMIT;

-- ----------------------------
-- Table structure for t_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_menu`;
CREATE TABLE `t_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父id',
  `path` varchar(64) NOT NULL DEFAULT '' COMMENT '路由路径',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '路由名称',
  `component` varchar(256) NOT NULL DEFAULT '' COMMENT '路由组件',
  `redirect` varchar(256) NOT NULL DEFAULT '' COMMENT '路由重定向',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '菜单类型',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单标题',
  `icon` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `rank` int NOT NULL DEFAULT '0' COMMENT '排序',
  `perm` varchar(64) NOT NULL DEFAULT '' COMMENT '权限标识',
  `params` varchar(256) NOT NULL DEFAULT '' COMMENT '路由参数',
  `keep_alive` tinyint NOT NULL DEFAULT '0' COMMENT '是否缓存',
  `always_show` tinyint NOT NULL DEFAULT '0' COMMENT '是否一直显示菜单',
  `is_hidden` tinyint NOT NULL DEFAULT '0' COMMENT '是否隐藏',
  `is_disable` tinyint NOT NULL DEFAULT '0' COMMENT '是否禁用',
  `extra` varchar(1024) NOT NULL DEFAULT '' COMMENT '菜单元数据',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_path` (`path`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Records of t_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `t_operation_log`;
CREATE TABLE `t_operation_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `nickname` varchar(64) DEFAULT '' COMMENT '用户昵称',
  `ip_address` varchar(255) DEFAULT '' COMMENT '操作ip',
  `ip_source` varchar(255) DEFAULT '' COMMENT '操作地址',
  `opt_module` varchar(32) DEFAULT '' COMMENT '操作模块',
  `opt_handler` varchar(32) DEFAULT '' COMMENT '操作方法',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='操作记录';

-- ----------------------------
-- Records of t_operation_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_photo
-- ----------------------------
DROP TABLE IF EXISTS `t_photo`;
CREATE TABLE `t_photo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_id` int NOT NULL DEFAULT '0' COMMENT '相册id',
  `photo_name` varchar(32) NOT NULL DEFAULT '' COMMENT '照片名',
  `photo_desc` varchar(64) NOT NULL DEFAULT '' COMMENT '照片描述',
  `photo_src` varchar(255) NOT NULL DEFAULT '' COMMENT '照片地址',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='照片';

-- ----------------------------
-- Records of t_photo
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_remark
-- ----------------------------
DROP TABLE IF EXISTS `t_remark`;
CREATE TABLE `t_remark` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
  `message_content` varchar(255) NOT NULL DEFAULT '' COMMENT '留言内容',
  `ip_address` varchar(64) NOT NULL DEFAULT '' COMMENT '用户ip 127.0.0.1',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT '用户地址 广东省深圳市',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态:0正常 1编辑 2撤回 3删除',
  `is_review` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核通过',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='留言';

-- ----------------------------
-- Records of t_remark
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父角色id',
  `role_name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名',
  `role_label` varchar(64) NOT NULL DEFAULT '' COMMENT '角色标签',
  `role_comment` varchar(64) NOT NULL DEFAULT '' COMMENT '角色备注',
  `is_disable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用  0否 1是',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认角色 0否 1是',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- ----------------------------
-- Records of t_role
-- ----------------------------
BEGIN;
INSERT INTO `t_role` (`id`, `parent_id`, `role_name`, `role_label`, `role_comment`, `is_disable`, `is_default`, `created_at`, `updated_at`) VALUES (1, 0, 'super-admin', '超级管理员', '', 0, 0, '2021-03-22 14:10:21', '2024-11-15 17:44:02');
COMMIT;

-- ----------------------------
-- Table structure for t_role_api
-- ----------------------------
DROP TABLE IF EXISTS `t_role_api`;
CREATE TABLE `t_role_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `api_id` int NOT NULL DEFAULT '0' COMMENT '接口id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-api关联';

-- ----------------------------
-- Records of t_role_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_role_menu`;
CREATE TABLE `t_role_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-菜单关联';

-- ----------------------------
-- Records of t_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_tag`;
CREATE TABLE `t_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `tag_name` varchar(32) NOT NULL DEFAULT '' COMMENT '标签名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_name` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='标签';

-- ----------------------------
-- Records of t_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_tag` (`id`, `tag_name`, `created_at`, `updated_at`) VALUES (1, '测试标签', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
COMMIT;

-- ----------------------------
-- Table structure for t_talk
-- ----------------------------
DROP TABLE IF EXISTS `t_talk`;
CREATE TABLE `t_talk` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '说说id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `content` varchar(2048) NOT NULL DEFAULT '' COMMENT '说说内容',
  `images` varchar(2048) NOT NULL DEFAULT '' COMMENT '图片',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1.公开 2.私密',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='说说';

-- ----------------------------
-- Records of t_talk
-- ----------------------------
BEGIN;
INSERT INTO `t_talk` (`id`, `user_id`, `content`, `images`, `is_top`, `status`, `like_count`, `created_at`, `updated_at`) VALUES (1, '1', '测试说说<img src=\"https://static.veweiyi.cn/emoji/qq/14@2x.gif\" width=\"24\" height=\"24\" alt=\"[微笑]\" style=\"margin: 0 1px;display: inline;vertical-align: text-bottom\">', 'null', 1, 1, 0, '2024-11-16 00:33:43', '2024-11-16 00:39:15');
COMMIT;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL COMMENT '用户id',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '用户密码',
  `nickname` varchar(64) NOT NULL COMMENT '用户昵称',
  `avatar` varchar(255) NOT NULL COMMENT '用户头像',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(64) DEFAULT '' COMMENT '手机号',
  `info` varchar(1024) NOT NULL COMMENT '用户信息',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: -1删除 0正常 1禁用',
  `login_type` varchar(64) NOT NULL DEFAULT '' COMMENT '注册方式',
  `ip_address` varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip',
  `ip_source` varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip 源',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_username` (`username`) USING BTREE,
  UNIQUE KEY `uk_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录信息';

-- ----------------------------
-- Records of t_user
-- ----------------------------
BEGIN;
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`, `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`) VALUES (1, '1', 'admin@qq.com', '$2a$10$ZINovpDg.FxFQRj6nhKDLOH55k19RDViybnVVn5EGuKQAcqChRs1e', '管理员', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'admin@qq.com', '', '{\"intro\":\"23\",\"website\":\"3\"}', 0, 'email', '127.0.0.1', '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2024-10-25 14:35:59');
COMMIT;

-- ----------------------------
-- Table structure for t_user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `t_user_login_history`;
CREATE TABLE `t_user_login_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录历史';

-- ----------------------------
-- Records of t_user_login_history
-- ----------------------------
BEGIN;
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (1, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:41:14', '1970-01-01 08:00:00', '2024-11-15 17:41:14', '2024-11-15 17:41:14');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (2, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:41:20', '1970-01-01 08:00:00', '2024-11-15 17:41:20', '2024-11-15 17:41:20');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (3, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:41:30', '1970-01-01 08:00:00', '2024-11-15 17:41:30', '2024-11-15 17:41:30');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (4, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:41:35', '1970-01-01 08:00:00', '2024-11-15 17:41:35', '2024-11-15 17:41:35');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (5, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:41:54', '1970-01-01 08:00:00', '2024-11-15 17:41:54', '2024-11-15 17:41:54');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (6, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 17:45:02', '1970-01-01 08:00:00', '2024-11-15 17:45:02', '2024-11-15 17:45:02');
INSERT INTO `t_user_login_history` (`id`, `user_id`, `login_type`, `agent`, `ip_address`, `ip_source`, `login_at`, `logout_at`, `created_at`, `updated_at`) VALUES (7, '1', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', '127.0.0.1', '本机地址', '2024-11-15 18:18:21', '1970-01-01 08:00:00', '2024-11-15 18:18:21', '2024-11-15 18:18:21');
COMMIT;

-- ----------------------------
-- Table structure for t_user_oauth
-- ----------------------------
DROP TABLE IF EXISTS `t_user_oauth`;
CREATE TABLE `t_user_oauth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `open_id` varchar(128) NOT NULL DEFAULT '' COMMENT '开发平台id，标识唯一用户',
  `platform` varchar(64) NOT NULL DEFAULT '' COMMENT '平台:手机号、邮箱、微信、飞书',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_oid_plat` (`open_id`,`platform`) USING BTREE,
  KEY `idx_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='第三方登录信息';

-- ----------------------------
-- Records of t_user_oauth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_user_role`;
CREATE TABLE `t_user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户-角色关联';

-- ----------------------------
-- Records of t_user_role
-- ----------------------------
BEGIN;
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`) VALUES (1, '1', 1);
COMMIT;

-- ----------------------------
-- Table structure for t_visit_history
-- ----------------------------
DROP TABLE IF EXISTS `t_visit_history`;
CREATE TABLE `t_visit_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `date` varchar(10) NOT NULL DEFAULT '' COMMENT '日期',
  `views_count` int NOT NULL DEFAULT '0' COMMENT '访问量',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面访问数量';

-- ----------------------------
-- Records of t_visit_history
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_website_config
-- ----------------------------
DROP TABLE IF EXISTS `t_website_config`;
CREATE TABLE `t_website_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `key` varchar(32) NOT NULL DEFAULT '' COMMENT '关键词',
  `config` varchar(2048) NOT NULL DEFAULT '' COMMENT '配置信息',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='网站配置表';

-- ----------------------------
-- Records of t_website_config
-- ----------------------------
BEGIN;
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (1, 'website_config', '{\"admin_url\":\"\",\"alipay_qr_code\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/17f234dc487c1bb5bbb732869be0eb53.jpg\",\"gitee\":\"https://gitee.com/wy791422171\",\"github\":\"https://github.com/ve-weiyi\",\"is_chat_room\":1,\"is_comment_review\":1,\"is_email_notice\":1,\"is_message_review\":0,\"is_music_player\":1,\"is_reward\":0,\"qq\":\"791422171\",\"social_login_list\":[\"qq\",\"github\",\"gitee\",\"feishu\",\"weibo\"],\"social_url_list\":[\"qq\",\"github\",\"gitee\"],\"tourist_avatar\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/5bfb96809bee5ba80a36811f0bf1d1ea.gif\",\"user_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175820.jpg\",\"website_author\":\"与梦\",\"website_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175746.jpg\",\"website_create_time\":\"2022-01-17\",\"website_intro\":\"你能做的，岂止如此。\",\"website_name\":\"与梦\",\"website_notice\":\"网站搭建问题请联系QQ 791422171。\",\"website_record_no\":\"桂ICP备2023013735号-1\",\"websocket_url\":\"wss://veweiyi.cn/api/websocket\",\"weixin_qr_code\":\"\"}', '2021-08-09 19:37:30', '2024-11-16 00:44:08');
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`) VALUES (2, 'about_me', '{\"content\":\"welcome to my blog!\"}', '2024-11-15 17:57:20', '2024-11-15 17:57:20');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
