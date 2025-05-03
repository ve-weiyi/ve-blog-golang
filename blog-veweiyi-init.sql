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

 Date: 06/05/2025 11:33:41
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_album
-- ----------------------------
DROP TABLE IF EXISTS `t_album`;
CREATE TABLE `t_album`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `album_name`  varchar(32)  NOT NULL DEFAULT '' COMMENT '相册名',
    `album_desc`  varchar(64)  NOT NULL DEFAULT '' COMMENT '相册描述',
    `album_cover` varchar(255) NOT NULL DEFAULT '' COMMENT '相册封面',
    `is_delete`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
    `status`      tinyint      NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密',
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='相册';

-- ----------------------------
-- Records of t_album
-- ----------------------------
BEGIN;
INSERT INTO `t_album` (`id`, `album_name`, `album_desc`, `album_cover`, `is_delete`, `status`, `created_at`,
                       `updated_at`)
VALUES (1, '壁纸', '壁纸', 'https://static.veweiyi.cn/blog/album/wusheng-20250502030442.jpg', 0, 1,
        '2025-05-02 03:04:45', '2025-05-02 03:04:45');
COMMIT;

-- ----------------------------
-- Table structure for t_api
-- ----------------------------
DROP TABLE IF EXISTS `t_api`;
CREATE TABLE `t_api`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id`  int          NOT NULL DEFAULT '0' COMMENT '分组id',
    `name`       varchar(128) NOT NULL DEFAULT '' COMMENT 'api名称',
    `path`       varchar(128) NOT NULL DEFAULT '' COMMENT 'api路径',
    `method`     varchar(16)  NOT NULL DEFAULT '' COMMENT 'api请求方法',
    `traceable`  tinyint      NOT NULL DEFAULT '0' COMMENT '是否追溯操作记录 0需要，1是',
    `is_disable` tinyint      NOT NULL DEFAULT '0' COMMENT '是否禁用 0否 1是',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_path_method` (`path`,`method`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口';

-- ----------------------------
-- Records of t_api
-- ----------------------------
BEGIN;
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (1, 0, 'account', 'account', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (2, 1, '修改用户角色', '/admin_api/v1/account/update_account_roles', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (3, 1, '查询在线用户列表', '/admin_api/v1/account/find_account_online_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (4, 1, '获取用户分布地区', '/admin_api/v1/account/find_account_area_analysis', 'POST', 0, 0,
        '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (5, 1, '修改用户状态', '/admin_api/v1/account/update_account_status', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (6, 1, '修改用户密码', '/admin_api/v1/account/update_account_password', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (7, 1, '查询用户列表', '/admin_api/v1/account/find_account_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (8, 1, '查询用户登录历史', '/admin_api/v1/account/find_account_login_history_list', 'POST', 0, 0,
        '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (9, 0, 'admin-api', 'admin-api', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (10, 9, 'ping', '/admin_api/v1/ping', 'GET', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (11, 0, 'album', 'album', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (12, 11, '查询相册', '/admin_api/v1/album/get_album', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (13, 11, '删除相册', '/admin_api/v1/album/delete_album', 'DELETE', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (14, 11, '分页获取相册列表', '/admin_api/v1/album/find_album_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (15, 11, '更新相册', '/admin_api/v1/album/update_album', 'PUT', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (16, 11, '创建相册', '/admin_api/v1/album/add_album', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (17, 0, 'api', 'api', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (18, 17, '创建api路由', '/admin_api/v1/api/add_api', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (19, 17, '批量删除api路由', '/admin_api/v1/api/batch_delete_api', 'DELETE', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (20, 17, '删除api路由', '/admin_api/v1/api/delete_api', 'DELETE', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (21, 17, '更新api路由', '/admin_api/v1/api/update_api', 'PUT', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (22, 17, '同步api列表', '/admin_api/v1/api/sync_api_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (23, 17, '清空接口列表', '/admin_api/v1/api/clean_api_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (24, 17, '分页获取api路由列表', '/admin_api/v1/api/find_api_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (25, 0, 'article', 'article', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (26, 25, '查询文章', '/admin_api/v1/article/get_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (27, 25, '删除文章', '/admin_api/v1/article/delete_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (28, 25, '导出文章列表', '/admin_api/v1/article/export_article_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (29, 25, '回收文章', '/admin_api/v1/article/recycle_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (30, 25, '添加文章', '/admin_api/v1/article/add_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (31, 25, '置顶文章', '/admin_api/v1/article/top_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (32, 25, '保存文章', '/admin_api/v1/article/update_article', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (33, 25, '查询文章列表', '/admin_api/v1/article/find_article_list', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (34, 0, 'auth', 'auth', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (35, 34, '重置密码', '/admin_api/v1/user/reset_password', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (36, 34, '注销', '/admin_api/v1/logoff', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (37, 34, '绑定邮箱', '/admin_api/v1/bind_user_email', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (38, 34, '第三方登录', '/admin_api/v1/oauth_login', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (39, 34, '发送绑定邮箱验证码', '/admin_api/v1/send_bind_email', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (40, 34, '发送重置密码邮件', '/admin_api/v1/user/send_reset_email', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (41, 34, '登录', '/admin_api/v1/login', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (42, 34, '注册', '/admin_api/v1/register', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (43, 34, '发送注册账号邮件', '/admin_api/v1/send_register_email', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (44, 34, '第三方登录授权地址', '/admin_api/v1/oauth_authorize_url', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (45, 34, '登出', '/admin_api/v1/logout', 'POST', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (46, 0, 'category', 'category', '', 0, 0, '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (47, 46, '分页获取文章分类列表', '/admin_api/v1/category/find_category_list', 'POST', 0, 0,
        '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (48, 46, '更新文章分类', '/admin_api/v1/category/update_category', 'PUT', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (49, 46, '批量删除文章分类', '/admin_api/v1/category/batch_delete_category', 'DELETE', 0, 0,
        '2025-04-29 23:17:01', '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (50, 46, '创建文章分类', '/admin_api/v1/category/add_category', 'POST', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (51, 46, '删除文章分类', '/admin_api/v1/category/delete_category', 'DELETE', 0, 0, '2025-04-29 23:17:01',
        '2025-04-29 23:17:01');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (52, 0, 'comment', 'comment', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (53, 52, '查询评论列表(后台)', '/admin_api/v1/comment/find_comment_back_list', 'POST', 0, 0,
        '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (54, 52, '批量删除评论', '/admin_api/v1/comment/batch_delete_comment', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (55, 52, '删除评论', '/admin_api/v1/comment/delete_comment', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (56, 52, '更新评论审核状态', '/admin_api/v1/comment/update_comment_review', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (57, 0, 'file', 'file', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (58, 57, '创建文件目录', '/admin_api/v1/file/add_file_folder', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (59, 57, '上传文件', '/admin_api/v1/file/upload_file', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (60, 57, '上传文件列表', '/admin_api/v1/file/multi_upload_file', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (61, 57, '获取文件列表', '/admin_api/v1/file/list_upload_file', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (62, 57, '删除文件列表', '/admin_api/v1/file/deletes_file', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (63, 57, '分页获取文件列表', '/admin_api/v1/file/find_file_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (64, 0, 'friend', 'friend', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (65, 64, '分页获取友链列表', '/admin_api/v1/friend/find_friend_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (66, 64, '更新友链', '/admin_api/v1/friend/update_friend', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (67, 64, '批量删除友链', '/admin_api/v1/friend/batch_delete_friend', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (68, 64, '创建友链', '/admin_api/v1/friend/add_friend', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (69, 64, '删除友链', '/admin_api/v1/friend/delete_friend', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (70, 0, 'menu', 'menu', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (71, 70, '批量删除菜单', '/admin_api/v1/menu/batch_delete_menu', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (72, 70, '同步菜单列表', '/admin_api/v1/menu/sync_menu_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (73, 70, '删除菜单', '/admin_api/v1/menu/delete_menu', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (74, 70, '创建菜单', '/admin_api/v1/menu/add_menu', 'POST', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (75, 70, '清空菜单列表', '/admin_api/v1/menu/clean_menu_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (76, 70, '分页获取菜单列表', '/admin_api/v1/menu/find_menu_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (77, 70, '更新菜单', '/admin_api/v1/menu/update_menu', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (78, 0, 'operation_log', 'operation_log', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (79, 78, '分页获取操作记录列表', '/admin_api/v1/operation_log/find_operation_log_list', 'POST', 0, 0,
        '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (80, 78, '删除操作记录', '/admin_api/v1/operation_log/deletes_operation_log', 'DELETE', 0, 0,
        '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (81, 0, 'page', 'page', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (82, 81, '创建页面', '/admin_api/v1/page/add_page', 'POST', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (83, 81, '分页获取页面列表', '/admin_api/v1/page/find_page_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (84, 81, '删除页面', '/admin_api/v1/page/delete_page', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (85, 81, '更新页面', '/admin_api/v1/page/update_page', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (86, 0, 'photo', 'photo', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (87, 86, '分页获取照片列表', '/admin_api/v1/photo/find_photo_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (88, 86, '创建照片', '/admin_api/v1/photo/add_photo', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (89, 86, '更新照片', '/admin_api/v1/photo/update_photo', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (90, 86, '删除照片', '/admin_api/v1/photo/delete_photo', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (91, 86, '批量删除照片', '/admin_api/v1/album/batch_delete_photo', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (92, 0, 'remark', 'remark', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (93, 92, '删除留言', '/admin_api/v1/remark/delete_remark', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (94, 92, '更新留言', '/admin_api/v1/remark/update_remark_review', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (95, 92, '分页获取留言列表', '/admin_api/v1/remark/find_remark_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (96, 92, '批量删除留言', '/admin_api/v1/remark/batch_delete_remark', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (97, 0, 'role', 'role', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (98, 97, '更新角色菜单权限', '/admin_api/v1/role/update_role_menus', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (99, 97, '删除角色', '/admin_api/v1/role/delete_role', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (100, 97, '更新角色', '/admin_api/v1/role/update_role', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (101, 97, '获取角色资源列表', '/admin_api/v1/role/find_role_resources', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (102, 97, '更新角色接口权限', '/admin_api/v1/role/update_role_apis', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (103, 97, '创建角色', '/admin_api/v1/role/add_role', 'POST', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (104, 97, '分页获取角色列表', '/admin_api/v1/role/find_role_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (105, 97, '批量删除角色', '/admin_api/v1/role/batch_delete_role', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (106, 0, 'tag', 'tag', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (107, 106, '创建标签', '/admin_api/v1/tag/add_tag', 'POST', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (108, 106, '批量删除标签', '/admin_api/v1/tag/batch_delete_tag', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (109, 106, '分页获取标签列表', '/admin_api/v1/tag/find_tag_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (110, 106, '更新标签', '/admin_api/v1/tag/update_tag', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (111, 106, '删除标签', '/admin_api/v1/tag/delete_tag', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (112, 0, 'talk', 'talk', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (113, 112, '分页获取说说列表', '/admin_api/v1/talk/find_talk_list', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (114, 112, '更新说说', '/admin_api/v1/talk/update_talk', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (115, 112, '创建说说', '/admin_api/v1/talk/add_talk', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (116, 112, '删除说说', '/admin_api/v1/talk/delete_talk', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (117, 112, '查询说说', '/admin_api/v1/talk/get_talk', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (118, 0, 'user', 'user', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (119, 118, '修改用户信息', '/admin_api/v1/user/update_user_info', 'POST', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (120, 118, '获取用户角色', '/admin_api/v1/user/get_user_roles', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (121, 118, '获取用户信息', '/admin_api/v1/user/get_user_info', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (122, 118, '查询用户登录历史', '/admin_api/v1/user/get_user_login_history_list', 'POST', 0, 0,
        '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (123, 118, '获取用户接口权限', '/admin_api/v1/user/get_user_apis', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (124, 118, '获取用户菜单权限', '/admin_api/v1/user/get_user_menus', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (125, 0, 'visit_log', 'visit_log', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (126, 125, '分页获取操作记录列表', '/admin_api/v1/visit_log/find_visit_log_list', 'POST', 0, 0,
        '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (127, 125, '删除操作记录', '/admin_api/v1/visit_log/deletes_visit_log', 'DELETE', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (128, 0, 'website', 'website', '', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (129, 128, '更新网站配置', '/admin_api/v1/admin/update_website_config', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (130, 128, '获取网站配置', '/admin_api/v1/admin/get_website_config', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (131, 128, '获取服务器信息', '/admin_api/v1/admin/system_state', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (132, 128, '获取关于我的信息', '/admin_api/v1/admin/about_me', 'GET', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (133, 128, '更新关于我的信息', '/admin_api/v1/admin/about_me', 'PUT', 0, 0, '2025-04-29 23:17:02',
        '2025-04-29 23:17:02');
INSERT INTO `t_api` (`id`, `parent_id`, `name`, `path`, `method`, `traceable`, `is_disable`, `created_at`, `updated_at`)
VALUES (134, 128, '获取后台首页信息', '/admin_api/v1/admin', 'GET', 0, 0, '2025-04-29 23:17:02', '2025-04-29 23:17:02');
COMMIT;

-- ----------------------------
-- Table structure for t_article
-- ----------------------------
DROP TABLE IF EXISTS `t_article`;
CREATE TABLE `t_article`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`         varchar(64)   NOT NULL DEFAULT '' COMMENT '作者',
    `category_id`     int           NOT NULL DEFAULT '0' COMMENT '文章分类',
    `article_cover`   varchar(1024) NOT NULL DEFAULT '' COMMENT '文章缩略图',
    `article_title`   varchar(64)   NOT NULL DEFAULT '' COMMENT '标题',
    `article_content` longtext      NOT NULL COMMENT '内容',
    `article_type`    tinyint       NOT NULL DEFAULT '0' COMMENT '文章类型 1原创 2转载 3翻译',
    `original_url`    varchar(255)  NOT NULL DEFAULT '' COMMENT '原文链接',
    `is_top`          tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶 0否 1是',
    `is_delete`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
    `status`          tinyint       NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密 3评论可见',
    `like_count`      int           NOT NULL DEFAULT '0' COMMENT '点赞数',
    `view_count`      int           NOT NULL DEFAULT '0' COMMENT '查看数',
    `created_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发表时间',
    `updated_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章';

-- ----------------------------
-- Records of t_article
-- ----------------------------
BEGIN;
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`,
                         `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `view_count`,
                         `created_at`, `updated_at`)
VALUES (1, 'admin', 1, 'https://static.veweiyi.cn/blog/article/qinglong-20241115174624.jpg', '测试文章',
        '恭喜你成功运行了博客！', 1, '', 1, 2, 1, 1, 1, '2024-11-15 17:46:29', '2025-05-02 03:57:07');
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`,
                         `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `view_count`,
                         `created_at`, `updated_at`)
VALUES (2, 'admin', 1, 'https://static.veweiyi.cn/blog/article/zhuqu-20241115182343.jpg', '草稿文章',
        '这是一篇草稿文章！', 1, '', 2, 2, 2, 1, 0, '2024-11-15 18:22:24', '2025-05-02 03:57:09');
INSERT INTO `t_article` (`id`, `user_id`, `category_id`, `article_cover`, `article_title`, `article_content`,
                         `article_type`, `original_url`, `is_top`, `is_delete`, `status`, `like_count`, `view_count`,
                         `created_at`, `updated_at`)
VALUES (3, 'root', 2, 'https://static.veweiyi.cn/blog/cover/zhuque.jpg', '从零开始部署博客系统',
        '# 博客系统部署指南\n\n## 文档说明\n\n本文档是博客系统部署的总体指南，详细介绍了从零开始部署一个完整博客系统的全过程。适合有一定基础的开发者参考。\n\n## 部署流程概览\n\n1. 前期准备\n    - [服务器选型与购买](资源准备（1）服务器准备)\n    - [域名注册与备案](资源准备（2）域名准备)\n    - [SSL证书配置](资源准备（3）SSL证书配置)\n    - [存储服务配置](资源准备（4）存储服务配置)\n\n2. 环境搭建\n    - [Docker环境配置](环境搭建（1）Docker.md)\n    - [MySQL数据库配置](环境搭建（2）MySQL.md)\n    - [Redis缓存配置](环境搭建（3）Redis.md)\n    - [RabbitMQ消息队列配置](环境搭建（4）RabbitMQ.md)\n\n3. 本地开发\n    - [开发环境配置](开发环境（1）基础配置.md)\n    - [配置文件说明](开发环境（2）配置文件.md)\n    - [开发调试指南](开发环境（3）调试指南.md)\n\n4. 生产部署\n    - [服务部署流程](生产部署（1）服务部署.md)\n    - [Nginx配置指南](生产部署（2）Nginx配置.md)\n    - [HTTPS配置说明](生产部署（3）HTTPS配置.md)\n\nTODO\n\n5. 运维管理\n    - [监控系统配置](运维管理（1）监控系统.md)\n    - [日志管理方案](运维管理（2）日志管理.md)\n    - [备份策略说明](运维管理（3）备份策略.md)\n    - [安全加固指南](运维管理（4）安全加固.md)\n\n6. 故障排查\n    - [常见问题处理](故障排查（1）常见问题.md)\n    - [性能优化指南](故障排查（2）性能优化.md)\n\n7. 更新维护\n    - [版本更新流程](更新维护（1）版本更新.md)\n    - [定期维护计划](更新维护（2）维护计划.md)\n',
        1, '', 1, 2, 1, 0, 2, '2025-05-05 00:20:03', '2025-05-05 00:35:30');
COMMIT;

-- ----------------------------
-- Table structure for t_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_article_tag`;
CREATE TABLE `t_article_tag`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `article_id` int NOT NULL DEFAULT '0' COMMENT '文章id',
    `tag_id`     int NOT NULL DEFAULT '0' COMMENT '标签id',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `fk_article_tag_1` (`article_id`) USING BTREE,
    KEY          `fk_article_tag_2` (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章-标签关联';

-- ----------------------------
-- Records of t_article_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`)
VALUES (1, 1, 1);
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`)
VALUES (2, 2, 1);
INSERT INTO `t_article_tag` (`id`, `article_id`, `tag_id`)
VALUES (6, 3, 2);
COMMIT;

-- ----------------------------
-- Table structure for t_category
-- ----------------------------
DROP TABLE IF EXISTS `t_category`;
CREATE TABLE `t_category`
(
    `id`            int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `category_name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名',
    `created_at`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文章分类';

-- ----------------------------
-- Records of t_category
-- ----------------------------
BEGIN;
INSERT INTO `t_category` (`id`, `category_name`, `created_at`, `updated_at`)
VALUES (1, '测试分类', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
INSERT INTO `t_category` (`id`, `category_name`, `created_at`, `updated_at`)
VALUES (2, '网站搭建', '2025-05-05 00:20:03', '2025-05-05 00:20:03');
COMMIT;

-- ----------------------------
-- Table structure for t_chat
-- ----------------------------
DROP TABLE IF EXISTS `t_chat`;
CREATE TABLE `t_chat`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     varchar(64)   NOT NULL DEFAULT '' COMMENT '用户id',
    `terminal_id` varchar(64)   NOT NULL DEFAULT '' COMMENT '设备id',
    `ip_address`  varchar(64)   NOT NULL DEFAULT '' COMMENT '用户ip 127.0.0.1',
    `ip_source`   varchar(128)  NOT NULL DEFAULT '' COMMENT '用户地址 广东省深圳市',
    `type`        varchar(64)   NOT NULL DEFAULT '' COMMENT '类型:chatgpt chatroom',
    `content`     varchar(4096) NOT NULL DEFAULT '' COMMENT '聊天内容',
    `status`      int           NOT NULL DEFAULT '0' COMMENT '状态:0正常 1编辑 2撤回 3删除',
    `created_at`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='聊天消息';

-- ----------------------------
-- Records of t_chat
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`         varchar(64)  NOT NULL DEFAULT '' COMMENT '评论用户id',
    `topic_id`        int          NOT NULL DEFAULT '0' COMMENT '主题id',
    `parent_id`       int          NOT NULL DEFAULT '0' COMMENT '父评论id',
    `reply_msg_id`    int          NOT NULL DEFAULT '0' COMMENT '回复评论id',
    `reply_user_id`   varchar(255) NOT NULL COMMENT '评论回复用户id',
    `comment_content` text         NOT NULL COMMENT '评论内容',
    `ip_address`      varchar(64)  NOT NULL COMMENT 'ip地址 127.0.01',
    `ip_source`       varchar(64)  NOT NULL COMMENT 'ip来源 广东省',
    `type`            tinyint      NOT NULL DEFAULT '0' COMMENT '评论类型 1.文章 2.友链 3.说说',
    `status`          tinyint      NOT NULL DEFAULT '0' COMMENT '状态 0.正常 1.已编辑 2.已删除',
    `is_review`       tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核通过',
    `like_count`      int          NOT NULL DEFAULT '0' COMMENT '评论点赞数量',
    `created_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`, `reply_user_id`) USING BTREE,
    KEY               `fk_comment_user` (`user_id`) USING BTREE,
    KEY               `fk_comment_parent` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='评论';

-- ----------------------------
-- Records of t_comment
-- ----------------------------
BEGIN;
INSERT INTO `t_comment` (`id`, `user_id`, `topic_id`, `parent_id`, `reply_msg_id`, `reply_user_id`, `comment_content`,
                         `ip_address`, `ip_source`, `type`, `status`, `is_review`, `like_count`, `created_at`,
                         `updated_at`)
VALUES (1, 'veweiyi', 0, 0, 0, '',
        '名称：与梦\n\n简介：你能做的，岂止如此。\n\n头像：https://static.veweiyi.cn/blog/website/tiger-20241115175746.jpg',
        '', '', 2, 0, 1, 0, '2025-05-05 02:11:43', '2025-05-05 02:11:43');
COMMIT;

-- ----------------------------
-- Table structure for t_file_folder
-- ----------------------------
DROP TABLE IF EXISTS `t_file_folder`;
CREATE TABLE `t_file_folder`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`     varchar(64)  NOT NULL DEFAULT '' COMMENT '用户id',
    `file_path`   varchar(128) NOT NULL DEFAULT '' COMMENT '文件路径',
    `folder_name` varchar(128) NOT NULL DEFAULT '' COMMENT '文件夹名称',
    `folder_desc` varchar(128) NOT NULL DEFAULT '' COMMENT '文件夹描述',
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
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
CREATE TABLE `t_file_upload`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    varchar(64)  NOT NULL DEFAULT '' COMMENT '用户id',
    `file_path`  varchar(128) NOT NULL DEFAULT '' COMMENT '文件路径',
    `file_name`  varchar(128) NOT NULL DEFAULT '' COMMENT '文件名称',
    `file_type`  varchar(128) NOT NULL DEFAULT '' COMMENT '文件类型',
    `file_size`  int          NOT NULL DEFAULT '0' COMMENT '文件大小',
    `file_md5`   varchar(128) NOT NULL DEFAULT '' COMMENT '文件md5值',
    `file_url`   varchar(256) NOT NULL DEFAULT '' COMMENT '上传路径',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_uid` (`user_id`) USING BTREE,
    KEY          `idx_path` (`file_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='上传记录';

-- ----------------------------
-- Records of t_file_upload
-- ----------------------------
BEGIN;
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (1, 'admin', '/', 'carousel', '', 0, '', '', '2025-05-02 01:30:43', '2025-05-02 03:57:23');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (2, 'admin', '/carousel', 'qinglong.jpg', '.jpg', 644494, '40dc6ddf12acd53fcc27722f4ca3b7f4',
        'https://static.veweiyi.cn/blog/carousel/qinglong-20250502013551.jpg', '2025-05-02 01:35:52',
        '2025-05-02 03:57:24');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (3, 'admin', '/carousel', 'baihu.jpg', '.jpg', 503323, '14f43565288c3fa2f8c2388a9289925e',
        'https://static.veweiyi.cn/blog/carousel/baihu-20250502013551.jpg', '2025-05-02 01:35:52',
        '2025-05-02 03:57:26');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (4, 'admin', '/carousel', 'zhuque.jpg', '.jpg', 672552, 'cf85e1676f1d4762268db22da610f12c',
        'https://static.veweiyi.cn/blog/carousel/zhuque-20250502013551.jpg', '2025-05-02 01:35:52',
        '2025-05-02 03:57:28');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (5, 'admin', '/carousel', 'xuanwu.jpg', '.jpg', 905978, 'c2cdf2b78315e9ea3ec37b968f4448e7',
        'https://static.veweiyi.cn/blog/carousel/xuanwu-20250502013551.jpg', '2025-05-02 01:35:52',
        '2025-05-02 03:57:31');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (6, 'admin', '/carousel', 'qilin.jpg', '.jpg', 257392, 'b0ab156ff05892669456be53f26bb3ae',
        'https://static.veweiyi.cn/blog/carousel/qilin-20250502013552.jpg', '2025-05-02 01:35:53',
        '2025-05-02 03:57:32');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (7, 'admin', '/carousel', 'wusheng.jpg', '.jpg', 395488, '7c576bcd6d5d56bcda8d9a39d1992f8a',
        'https://static.veweiyi.cn/blog/carousel/wusheng-20250502013552.jpg', '2025-05-02 01:35:53',
        '2025-05-02 03:57:33');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (9, 'admin', '/album', 'wusheng.jpg', '.jpg', 395488, '7c576bcd6d5d56bcda8d9a39d1992f8a',
        'https://static.veweiyi.cn/blog/album/wusheng-20250502030442.jpg', '2025-05-02 03:04:43',
        '2025-05-02 03:57:35');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (10, 'admin', '/photo', 'qinglong.jpg', '.jpg', 644494, '40dc6ddf12acd53fcc27722f4ca3b7f4',
        'https://static.veweiyi.cn/blog/photo/qinglong-20250502030854.jpg', '2025-05-02 03:08:55',
        '2025-05-02 03:57:36');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (11, 'admin', '/photo', 'qilin.jpg', '.jpg', 257392, 'b0ab156ff05892669456be53f26bb3ae',
        'https://static.veweiyi.cn/blog/photo/qilin-20250502030854.jpg', '2025-05-02 03:08:55', '2025-05-02 03:57:38');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (12, 'admin', '/photo', 'xuanwu.jpg', '.jpg', 905978, 'c2cdf2b78315e9ea3ec37b968f4448e7',
        'https://static.veweiyi.cn/blog/photo/xuanwu-20250502030854.jpg', '2025-05-02 03:08:55', '2025-05-02 03:57:39');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (13, 'admin', '/photo', 'zhuque.jpg', '.jpg', 672552, 'cf85e1676f1d4762268db22da610f12c',
        'https://static.veweiyi.cn/blog/photo/zhuque-20250502030854.jpg', '2025-05-02 03:08:55', '2025-05-02 03:57:40');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (14, 'admin', '/photo', 'baihu.jpg', '.jpg', 503323, '14f43565288c3fa2f8c2388a9289925e',
        'https://static.veweiyi.cn/blog/photo/baihu-20250502030854.jpg', '2025-05-02 03:08:55', '2025-05-02 03:57:42');
INSERT INTO `t_file_upload` (`id`, `user_id`, `file_path`, `file_name`, `file_type`, `file_size`, `file_md5`,
                             `file_url`, `created_at`, `updated_at`)
VALUES (15, 'admin', '/photo', 'wusheng.jpg', '.jpg', 395488, '7c576bcd6d5d56bcda8d9a39d1992f8a',
        'https://static.veweiyi.cn/blog/photo/wusheng-20250502030854.jpg', '2025-05-02 03:08:55',
        '2025-05-02 03:57:43');
COMMIT;

-- ----------------------------
-- Table structure for t_friend
-- ----------------------------
DROP TABLE IF EXISTS `t_friend`;
CREATE TABLE `t_friend`
(
    `id`           int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `link_name`    varchar(32)  NOT NULL DEFAULT '' COMMENT '链接名',
    `link_avatar`  varchar(255) NOT NULL DEFAULT '' COMMENT '链接头像',
    `link_address` varchar(64)  NOT NULL DEFAULT '' COMMENT '链接地址',
    `link_intro`   varchar(100) NOT NULL DEFAULT '' COMMENT '链接介绍',
    `created_at`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_name` (`link_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='友链';

-- ----------------------------
-- Records of t_friend
-- ----------------------------
BEGIN;
INSERT INTO `t_friend` (`id`, `link_name`, `link_avatar`, `link_address`, `link_intro`, `created_at`, `updated_at`)
VALUES (1, '与梦', 'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'https://blog.veweiyi.cn',
        '你能做的，岂止如此。', '2024-11-16 00:43:12', '2024-11-16 00:43:37');
COMMIT;

-- ----------------------------
-- Table structure for t_login_log
-- ----------------------------
DROP TABLE IF EXISTS `t_login_log`;
CREATE TABLE `t_login_log`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    varchar(64)  NOT NULL DEFAULT '' COMMENT '用户id',
    `login_type` varchar(64)  NOT NULL DEFAULT '' COMMENT '登录类型',
    `app_name`   varchar(64)  NOT NULL DEFAULT '' COMMENT 'app名称',
    `os`         varchar(64)  NOT NULL DEFAULT '' COMMENT '操作系统',
    `browser`    varchar(64)  NOT NULL DEFAULT '' COMMENT '浏览器',
    `ip_address` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip host',
    `ip_source`  varchar(255) NOT NULL DEFAULT '' COMMENT 'ip 源',
    `login_at`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
    `logout_at`  datetime              DEFAULT NULL COMMENT '登出时间',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录历史';

-- ----------------------------
-- Records of t_login_log
-- ----------------------------
BEGIN;
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (1, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[', '', '2025-05-02 21:09:38',
        '2025-05-02 21:17:41', '2025-05-02 21:09:38', '2025-05-02 21:17:41');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (2, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[', '', '2025-05-02 21:17:43',
        '2025-05-02 21:45:53', '2025-05-02 21:17:43', '2025-05-02 21:45:54');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (3, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', 'http', '', '2025-05-02 21:45:57',
        '2025-05-02 22:03:50', '2025-05-02 21:45:57', '2025-05-02 22:03:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (4, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', 'http', '', '2025-05-02 22:03:53',
        '2025-05-02 22:05:29', '2025-05-02 22:03:53', '2025-05-02 22:05:30');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (5, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[', '', '2025-05-02 22:05:31',
        '2025-05-02 22:06:52', '2025-05-02 22:05:31', '2025-05-02 22:06:53');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (6, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[', '', '2025-05-02 22:06:54',
        '2025-05-02 22:12:23', '2025-05-02 22:06:54', '2025-05-02 22:12:23');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (7, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:51246', '本机地址',
        '2025-05-02 22:12:25', '2025-05-02 23:15:48', '2025-05-02 22:12:25', '2025-05-02 23:15:49');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (8, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:56029', '本机地址',
        '2025-05-02 23:15:51', NULL, '2025-05-02 23:15:51', '2025-05-02 23:15:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (9, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:62754', '本机地址',
        '2025-05-03 05:13:26', NULL, '2025-05-03 05:13:26', '2025-05-03 05:13:26');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (10, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:62815', '本机地址',
        '2025-05-03 05:13:36', '2025-05-03 05:14:36', '2025-05-03 05:13:36', '2025-05-03 05:14:36');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (11, 'admin', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:63250', '本机地址',
        '2025-05-03 05:14:41', '2025-05-03 05:14:46', '2025-05-03 05:14:41', '2025-05-03 05:14:46');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (12, 'veweiyi', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:56897', '本机地址',
        '2025-05-03 05:44:04', '2025-05-03 05:44:58', '2025-05-03 05:44:04', '2025-05-03 05:44:58');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (13, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:55812', '本机地址',
        '2025-05-03 16:18:22', NULL, '2025-05-03 16:18:22', '2025-05-03 16:18:22');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (14, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:55885', '本机地址',
        '2025-05-03 16:18:33', '2025-05-03 16:22:24', '2025-05-03 16:18:33', '2025-05-03 16:22:25');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (15, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:57361', '本机地址',
        '2025-05-03 16:22:27', '2025-05-03 16:23:41', '2025-05-03 16:22:27', '2025-05-03 16:23:42');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (16, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:57926', '本机地址',
        '2025-05-03 16:23:44', '2025-05-03 16:23:46', '2025-05-03 16:23:44', '2025-05-03 16:23:47');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (17, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:58117', '本机地址',
        '2025-05-03 16:24:09', '2025-05-03 16:24:18', '2025-05-03 16:24:09', '2025-05-03 16:24:19');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (18, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:59306', '本机地址',
        '2025-05-03 16:27:10', NULL, '2025-05-03 16:27:10', '2025-05-03 16:27:10');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (19, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:59330', '本机地址',
        '2025-05-03 16:27:12', NULL, '2025-05-03 16:27:12', '2025-05-03 16:27:12');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (20, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60143', '本机地址',
        '2025-05-03 16:29:49', NULL, '2025-05-03 16:29:49', '2025-05-03 16:29:49');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (21, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60446', '本机地址',
        '2025-05-03 16:30:35', '2025-05-03 16:30:39', '2025-05-03 16:30:35', '2025-05-03 16:30:39');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (22, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60527', '本机地址',
        '2025-05-03 16:30:40', '2025-05-03 16:30:44', '2025-05-03 16:30:40', '2025-05-03 16:30:44');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (23, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60595', '本机地址',
        '2025-05-03 16:30:47', '2025-05-03 16:30:50', '2025-05-03 16:30:47', '2025-05-03 16:30:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (24, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60695', '本机地址',
        '2025-05-03 16:30:56', '2025-05-03 16:31:15', '2025-05-03 16:30:56', '2025-05-03 16:31:15');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (25, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60879', '本机地址',
        '2025-05-03 16:31:16', '2025-05-03 16:31:50', '2025-05-03 16:31:16', '2025-05-03 16:31:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (26, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:61116', '本机地址',
        '2025-05-03 16:31:53', '2025-05-03 16:36:23', '2025-05-03 16:31:53', '2025-05-03 16:36:24');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (27, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:63378', '本机地址',
        '2025-05-03 16:38:22', '2025-05-03 16:38:27', '2025-05-03 16:38:22', '2025-05-03 16:38:28');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (28, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:63456', '本机地址',
        '2025-05-03 16:38:30', '2025-05-03 17:14:02', '2025-05-03 16:38:30', '2025-05-03 17:14:02');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (29, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:59559', '本机地址',
        '2025-05-03 17:14:04', NULL, '2025-05-03 17:14:04', '2025-05-03 17:14:04');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (30, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:61143', '本机地址',
        '2025-05-04 22:34:54', '2025-05-05 04:12:51', '2025-05-04 22:34:54', '2025-05-05 04:12:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (31, 'veweiyi', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:64925', '本机地址',
        '2025-05-05 01:03:20', '2025-05-05 05:38:21', '2025-05-05 01:03:20', '2025-05-05 05:38:21');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (32, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:63845', '本机地址',
        '2025-05-05 04:17:29', '2025-05-05 04:17:32', '2025-05-05 04:17:29', '2025-05-05 04:17:33');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (33, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:56273', '本机地址',
        '2025-05-05 04:42:57', '2025-05-05 04:43:02', '2025-05-05 04:42:57', '2025-05-05 04:43:02');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (34, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:56951', '本机地址',
        '2025-05-05 04:44:36', '2025-05-05 04:44:44', '2025-05-05 04:44:36', '2025-05-05 04:44:44');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (35, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60313', '本机地址',
        '2025-05-05 04:52:37', '2025-05-05 04:52:54', '2025-05-05 04:52:37', '2025-05-05 04:52:55');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (36, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:61843', '本机地址',
        '2025-05-05 04:56:14', '2025-05-05 04:56:41', '2025-05-05 04:56:14', '2025-05-05 04:56:41');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (37, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:51866', '本机地址',
        '2025-05-05 05:19:49', '2025-05-05 05:21:46', '2025-05-05 05:19:49', '2025-05-05 05:21:47');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (38, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:52646', '本机地址',
        '2025-05-05 05:22:02', '2025-05-05 05:22:27', '2025-05-05 05:22:02', '2025-05-05 05:22:28');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (39, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '127.0.0.1:55220', '本机地址', '2025-05-05 06:17:24', '2025-05-05 06:17:42', '2025-05-05 06:17:24',
        '2025-05-05 06:17:43');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (40, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '127.0.0.1:55844', '本机地址', '2025-05-05 06:18:49', '2025-05-05 06:19:30', '2025-05-05 06:18:49',
        '2025-05-05 06:19:30');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (41, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:61626', '本机地址', '2025-05-05 06:35:17', NULL, '2025-05-05 06:35:17', '2025-05-05 06:35:17');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (42, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:61765', '本机地址', '2025-05-05 06:35:31', NULL, '2025-05-05 06:35:31', '2025-05-05 06:35:31');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (43, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:62924', '本机地址', '2025-05-05 06:38:37', NULL, '2025-05-05 06:38:37', '2025-05-05 06:38:37');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (44, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:63761', '本机地址', '2025-05-05 06:41:14', '2025-05-05 06:43:10', '2025-05-05 06:41:14',
        '2025-05-05 06:43:11');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (45, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:64451', '本机地址', '2025-05-05 06:43:15', '2025-05-05 06:43:42', '2025-05-05 06:43:15',
        '2025-05-05 06:43:43');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (46, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:59956', '本机地址', '2025-05-05 19:33:12', '2025-05-05 19:33:21', '2025-05-05 19:33:12',
        '2025-05-05 19:33:22');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (47, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:60667', '本机地址', '2025-05-05 19:35:02', '2025-05-05 19:35:06', '2025-05-05 19:35:02',
        '2025-05-05 19:35:06');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (48, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:62952', '本机地址', '2025-05-05 19:39:43', '2025-05-05 19:39:48', '2025-05-05 19:39:43',
        '2025-05-05 19:39:49');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (49, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:63858', '本机地址', '2025-05-05 20:18:42', '2025-05-05 20:20:20', '2025-05-05 20:18:42',
        '2025-05-05 20:20:21');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (50, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:64599', '本机地址', '2025-05-05 20:20:31', NULL, '2025-05-05 20:20:31', '2025-05-05 20:20:31');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (51, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:64774', '本机地址', '2025-05-05 20:20:55', NULL, '2025-05-05 20:20:55', '2025-05-05 20:20:55');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (52, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 'gitee', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:65132', '本机地址', '2025-05-05 20:21:51', NULL, '2025-05-05 20:21:51', '2025-05-05 20:21:51');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (53, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:49595', '本机地址',
        '2025-05-05 20:23:55', NULL, '2025-05-05 20:23:55', '2025-05-05 20:23:55');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (54, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:56129', '本机地址', '2025-05-05 21:29:49', '2025-05-05 21:30:08', '2025-05-05 21:29:49',
        '2025-05-05 21:30:09');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (55, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:56508', '本机地址',
        '2025-05-05 21:30:49', NULL, '2025-05-05 21:30:49', '2025-05-05 21:30:49');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (56, '582d1d06-e496-4d87-888f-9910c09a9149', 'github', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '[::1]:56580', '本机地址', '2025-05-05 21:30:56', '2025-05-05 21:35:21', '2025-05-05 21:30:56',
        '2025-05-05 21:35:22');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (57, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:58099', '本机地址',
        '2025-05-05 21:35:23', '2025-05-05 22:29:30', '2025-05-05 21:35:23', '2025-05-05 22:29:31');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (58, 'root', 'email', 'admin-web', 'Intel Mac OS X 10_15_7', 'Chrome', '[::1]:60598', '本机地址',
        '2025-05-05 22:29:39', '2025-05-05 22:56:06', '2025-05-05 22:29:39', '2025-05-05 22:56:07');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (59, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:62760', '本机地址',
        '2025-05-05 22:35:52', '2025-05-05 22:36:34', '2025-05-05 22:35:52', '2025-05-05 22:36:34');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (60, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:53677', '本机地址',
        '2025-05-05 22:59:46', '2025-05-05 23:13:41', '2025-05-05 22:59:46', '2025-05-05 23:13:42');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (61, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:65428', '本机地址',
        '2025-05-05 23:32:55', '2025-05-05 23:33:24', '2025-05-05 23:32:55', '2025-05-05 23:33:25');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (62, '4bd72658-757d-418c-ad0a-55371db6cb16', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome',
        '127.0.0.1:52998', '本机地址', '2025-05-05 23:45:11', NULL, '2025-05-05 23:45:11', '2025-05-05 23:45:11');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (63, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:65298', '本机地址',
        '2025-05-06 00:23:53', NULL, '2025-05-06 00:23:53', '2025-05-06 00:23:53');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (64, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:65501', '本机地址',
        '2025-05-06 00:24:24', '2025-05-06 00:41:01', '2025-05-06 00:24:24', '2025-05-06 00:41:02');
INSERT INTO `t_login_log` (`id`, `user_id`, `login_type`, `app_name`, `os`, `browser`, `ip_address`, `ip_source`,
                           `login_at`, `logout_at`, `created_at`, `updated_at`)
VALUES (65, 'root', 'email', 'blog-web', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:55472', '本机地址',
        '2025-05-06 00:43:29', NULL, '2025-05-06 00:43:29', '2025-05-06 00:43:29');
COMMIT;

-- ----------------------------
-- Table structure for t_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_menu`;
CREATE TABLE `t_menu`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `parent_id`   int           NOT NULL DEFAULT '0' COMMENT '父id',
    `path`        varchar(64)   NOT NULL DEFAULT '' COMMENT '路由路径',
    `name`        varchar(64)   NOT NULL DEFAULT '' COMMENT '路由名称',
    `component`   varchar(256)  NOT NULL DEFAULT '' COMMENT '路由组件',
    `redirect`    varchar(256)  NOT NULL DEFAULT '' COMMENT '路由重定向',
    `type`        varchar(64)   NOT NULL DEFAULT '0' COMMENT '菜单类型',
    `title`       varchar(64)   NOT NULL DEFAULT '' COMMENT '菜单标题',
    `icon`        varchar(64)   NOT NULL DEFAULT '' COMMENT '菜单图标',
    `rank`        int           NOT NULL DEFAULT '0' COMMENT '排序',
    `perm`        varchar(64)   NOT NULL DEFAULT '' COMMENT '权限标识',
    `params`      varchar(256)  NOT NULL DEFAULT '' COMMENT '路由参数',
    `keep_alive`  tinyint       NOT NULL DEFAULT '0' COMMENT '是否缓存',
    `always_show` tinyint       NOT NULL DEFAULT '0' COMMENT '是否一直显示菜单',
    `is_hidden`   tinyint       NOT NULL DEFAULT '0' COMMENT '是否隐藏',
    `is_disable`  tinyint       NOT NULL DEFAULT '0' COMMENT '是否禁用',
    `extra`       varchar(1024) NOT NULL DEFAULT '' COMMENT '菜单元数据',
    `created_at`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_path` (`path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Records of t_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (1, 0, '/article', '', '/src/layout/index', '/article/publish', 'CATALOG', '文章管理', 'el-icon-document', 1, '',
        'null', 0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"文章管理\",\"icon\":\"el-icon-document\",\"rank\":1,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (2, 1, '/article/publish', 'ArticlePublish', '/src/views/admin/blog/article/Write', '', 'MENU', '发布文章', '',
        1, '', 'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"发布文章\",\"rank\":1,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (3, 1, '/article/edit/:articleId', 'ArticleEdit', '/src/views/admin/blog/article/Write', '', 'MENU', '查看文章',
        '', 2, '', 'null', 1, 0, 1, 0,
        '{\"type\":\"MENU\",\"title\":\"查看文章\",\"rank\":2,\"params\":\"null\",\"keep_alive\":1,\"is_hidden\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (4, 1, '/article/list', 'ArticleList', '/src/views/admin/blog/article/Article', '', 'MENU', '文章列表', '', 3,
        '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"文章列表\",\"rank\":3,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (5, 1, '/article/category', 'Category', '/src/views/admin/blog/category/Category', '', 'MENU', '分类管理', '', 4,
        '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"分类管理\",\"rank\":4,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (6, 1, '/article/tag', 'Tag', '/src/views/admin/blog/tag/Tag', '', 'MENU', '标签管理', '', 5, '', 'null', 0, 0,
        0, 0, '{\"type\":\"MENU\",\"title\":\"标签管理\",\"rank\":5,\"params\":\"null\"}', '2025-04-29 23:25:41',
        '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (7, 0, '/message', '', '/src/layout/index', '/message/comment', 'CATALOG', '消息管理', 'el-icon-message', 2, '',
        'null', 0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"消息管理\",\"icon\":\"el-icon-message\",\"rank\":2,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (8, 7, '/message/comment', 'Comment', '/src/views/admin/message/comment/Comment', '', 'MENU', '评论管理', '', 1,
        '', 'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"评论管理\",\"rank\":1,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (9, 7, '/message/remark', 'Remark', '/src/views/admin/message/remark/Remark', '', 'MENU', '留言管理', '', 2, '',
        'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"留言管理\",\"rank\":2,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (10, 0, '/resource', '', '/src/layout/index', '/resource/file', 'CATALOG', '资源管理', 'el-icon-folder', 3, '',
        'null', 0, 1, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"资源管理\",\"icon\":\"el-icon-folder\",\"rank\":3,\"params\":\"null\",\"always_show\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (11, 10, '/resource/file', 'File', '/src/views/admin/resource/file/index', '', 'MENU', '文件管理', '', 1, '',
        'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"文件管理\",\"rank\":1,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (12, 10, '/picture', '', '', '/picture/albums', 'CATALOG', '图片管理', 'el-icon-picture', 2, '', 'null', 0, 0, 0,
        0, '{\"type\":\"CATALOG\",\"title\":\"图片管理\",\"icon\":\"el-icon-picture\",\"rank\":2,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (13, 12, '/picture/albums', 'Albums', '/src/views/admin/resource/picture/album/Album', '', 'MENU', '相册管理',
        '', 1, '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"相册管理\",\"rank\":1,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (14, 12, '/picture/albums/:id', 'Photo', '/src/views/admin/resource/picture/album/Photo', '', 'MENU', '相册详情',
        '', 2, '', 'null', 0, 0, 1, 0,
        '{\"type\":\"MENU\",\"title\":\"相册详情\",\"rank\":2,\"params\":\"null\",\"is_hidden\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (15, 12, '/picture/photo/delete', 'PhotoDelete', '/src/views/admin/resource/picture/album/Delete', '', 'MENU',
        '相片回收站', '', 3, '', 'null', 0, 0, 1, 0,
        '{\"type\":\"MENU\",\"title\":\"相片回收站\",\"rank\":3,\"params\":\"null\",\"is_hidden\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (16, 0, '/log', '', '/src/layout/index', '/log/operation', 'CATALOG', '日志管理', 'el-icon-edit', 4, '', 'null',
        0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"日志管理\",\"icon\":\"el-icon-edit\",\"rank\":4,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (17, 16, '/log/login', 'LogLogin', '/src/views/admin/log/login/Login', '', 'MENU', '登录日志', '', 1, '', 'null',
        1, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"登录日志\",\"rank\":1,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (18, 16, '/log/operation', 'LogOperation', '/src/views/admin/log/operation/Operation', '', 'MENU', '操作日志',
        '', 2, '', 'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"操作日志\",\"rank\":2,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (19, 16, '/log/visit', 'LogVisit', '/src/views/admin/log/visit/Visit', '', 'MENU', '浏览日志', '', 3, '', 'null',
        1, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"浏览日志\",\"rank\":3,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (20, 0, '/user', '', '/src/layout/index', '/monitor/online', 'CATALOG', '系统监控', 'el-icon-monitor', 5, '',
        'null', 0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"系统监控\",\"icon\":\"el-icon-monitor\",\"rank\":5,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (21, 20, '/monitor/online', 'Online', '/src/views/admin/monitor/online/Online', '', 'MENU', '在线用户', '', 1,
        '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"在线用户\",\"rank\":1,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (22, 20, '/monitor/state', 'State', '/src/views/admin/monitor/state/State', '', 'MENU', '服务器状态', '', 2, '',
        'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"服务器状态\",\"rank\":2,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (23, 0, '/system', '', '/src/layout/index', '/system/user', 'CATALOG', '系统管理', 'el-icon-setting', 6, '',
        'null', 0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"系统管理\",\"icon\":\"el-icon-setting\",\"rank\":6,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (24, 23, '/system/user', 'User', '/src/views/admin/system/user/User', '', 'MENU', '用户列表', '', 1, '', 'null',
        0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"用户列表\",\"rank\":1,\"params\":\"null\"}', '2025-04-29 23:25:41',
        '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (25, 23, '/system/role', 'Role', '/src/views/admin/system/role/Role', '', 'MENU', '角色管理', '', 2, '', 'null',
        0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"角色管理\",\"rank\":2,\"params\":\"null\"}', '2025-04-29 23:25:41',
        '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (26, 23, '/system/menu', 'Menu', '/src/views/admin/system/menu/Menu', '', 'MENU', '菜单管理', '', 3, '', 'null',
        0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"菜单管理\",\"rank\":3,\"params\":\"null\"}', '2025-04-29 23:25:41',
        '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (27, 23, '/system/api', 'Api', '/src/views/admin/system/api/Api', '', 'MENU', '接口管理', '', 4, '', 'null', 0,
        0, 0, 0, '{\"type\":\"MENU\",\"title\":\"接口管理\",\"rank\":4,\"params\":\"null\"}', '2025-04-29 23:25:41',
        '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (28, 0, '/website', '', '/src/layout/index', '/website/profile', 'CATALOG', '网站管理', 'el-icon-operation', 7,
        '', 'null', 0, 0, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"网站管理\",\"icon\":\"el-icon-operation\",\"rank\":7,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (29, 28, '/website/profile', 'Config', '/src/views/admin/website/profile/Profile', '', 'MENU', '网站设置', '', 1,
        '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"网站设置\",\"rank\":1,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (30, 28, '/website/page', 'Page', '/src/views/admin/website/page/Page', '', 'MENU', '页面管理', '', 2, '',
        'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"页面管理\",\"rank\":2,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (31, 28, '/website/talk', 'Talk', '/src/views/admin/website/talk/Talk', '', 'MENU', '说说管理', '', 3, '',
        'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"说说管理\",\"rank\":3,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (32, 28, '/website/friend', 'Friend', '/src/views/admin/website/friend/Friend', '', 'MENU', '友链管理', '', 4,
        '', 'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"友链管理\",\"rank\":4,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (33, 28, '/website/about', 'AboutMe', '/src/views/admin/website/about/About', '', 'MENU', '关于我', '', 5, '',
        'null', 0, 0, 0, 0, '{\"type\":\"MENU\",\"title\":\"关于我\",\"rank\":5,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (34, 0, '/document', '', '/src/layout/index', '/document/apifox', 'CATALOG', '接口文档', 'api', 8, '', 'null', 0,
        1, 0, 0,
        '{\"type\":\"CATALOG\",\"title\":\"接口文档\",\"icon\":\"api\",\"rank\":8,\"params\":\"null\",\"always_show\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (35, 34, '/document/apifox', 'Apifox', '/src/views/admin/document/api/apifox', '', 'MENU', 'Apifox', 'api', 1,
        '', 'null', 1, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"Apifox\",\"icon\":\"api\",\"rank\":1,\"params\":\"null\",\"keep_alive\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (36, 34, '/document/doc', 'InternalDoc', '/src/views/admin/document/doc/internal-doc', '', 'MENU', '博客前台',
        'document', 2, '', 'null', 0, 0, 0, 0,
        '{\"type\":\"MENU\",\"title\":\"博客前台\",\"icon\":\"document\",\"rank\":2,\"params\":\"null\"}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (37, 0, '/mine', '', '/src/layout/index', '/mine/info', 'CATALOG', '个人中心', 'el-icon-user', 9, '', 'null', 0,
        1, 1, 0,
        '{\"type\":\"CATALOG\",\"title\":\"个人中心\",\"icon\":\"el-icon-user\",\"rank\":9,\"params\":\"null\",\"always_show\":1,\"is_hidden\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
INSERT INTO `t_menu` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `type`, `title`, `icon`, `rank`,
                      `perm`, `params`, `keep_alive`, `always_show`, `is_hidden`, `is_disable`, `extra`, `created_at`,
                      `updated_at`)
VALUES (38, 37, '/mine/info', 'Mine', '/src/views/admin/mine/Mine', '', 'MENU', '个人信息', '', 1, '', 'null', 0, 0, 1,
        0, '{\"type\":\"MENU\",\"title\":\"个人信息\",\"rank\":1,\"params\":\"null\",\"is_hidden\":1}',
        '2025-04-29 23:25:41', '2025-04-29 23:25:41');
COMMIT;

-- ----------------------------
-- Table structure for t_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `t_operation_log`;
CREATE TABLE `t_operation_log`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`         varchar(64)   NOT NULL DEFAULT '' COMMENT '用户id',
    `terminal_id`     varchar(64)   NOT NULL DEFAULT '' COMMENT '设备id',
    `ip_address`      varchar(255)  NOT NULL DEFAULT '' COMMENT '操作ip',
    `ip_source`       varchar(255)  NOT NULL DEFAULT '' COMMENT '操作地址',
    `opt_module`      varchar(32)   NOT NULL DEFAULT '' COMMENT '操作模块',
    `opt_desc`        varchar(255)  NOT NULL DEFAULT '' COMMENT '操作描述',
    `request_uri`     varchar(255)  NOT NULL DEFAULT '' COMMENT '请求地址',
    `request_method`  varchar(32)   NOT NULL DEFAULT '' COMMENT '请求方式',
    `request_data`    varchar(4096) NOT NULL DEFAULT '' COMMENT '请求参数',
    `response_data`   varchar(4096) NOT NULL DEFAULT '' COMMENT '返回数据',
    `response_status` int           NOT NULL DEFAULT '0' COMMENT '响应状态码',
    `cost`            varchar(32)   NOT NULL DEFAULT '' COMMENT '耗时（ms）',
    `created_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='操作记录';

-- ----------------------------
-- Records of t_operation_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_page
-- ----------------------------
DROP TABLE IF EXISTS `t_page`;
CREATE TABLE `t_page`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT COMMENT '页面id',
    `page_name`       varchar(32)   NOT NULL DEFAULT '' COMMENT '页面名',
    `page_label`      varchar(32)   NOT NULL DEFAULT '' COMMENT '页面标签',
    `page_cover`      varchar(255)  NOT NULL DEFAULT '' COMMENT '页面封面',
    `is_carousel`     tinyint       NOT NULL DEFAULT '0' COMMENT '是否轮播',
    `carousel_covers` varchar(1024) NOT NULL DEFAULT '' COMMENT '轮播图片列表',
    `created_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面';

-- ----------------------------
-- Records of t_page
-- ----------------------------
BEGIN;
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (1, '首页', 'home', 'https://veport.oss-cn-beijing.aliyuncs.com/config/f9fa18da262910eb13f802b003147915.jpg', 1,
        '', '2021-08-07 10:32:36', '2024-11-21 14:42:49');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (2, '归档', 'archive', 'https://veport.oss-cn-beijing.aliyuncs.com/config/82fc9c41de3c511ca1532d978b36fec7.jpg',
        1, '', '2021-08-07 10:32:36', '2024-11-21 14:42:56');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (3, '分类', 'category', 'https://veport.oss-cn-beijing.aliyuncs.com/config/f9fa18da262910eb13f802b003147915.jpg',
        1, '', '2021-08-07 10:32:36', '2024-11-21 14:43:00');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (4, '标签', 'tag', 'https://static.veweiyi.cn/blog/page/remu-20241121141754.jpeg', 0, '', '2021-08-07 10:32:36',
        '2024-11-21 14:43:04');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (5, '相册', 'album', 'https://veport.oss-cn-beijing.aliyuncs.com/config/dd3678e409cab21ff2e5f875976058a6.jpg', 0,
        '', '2021-08-07 10:32:36', '2022-01-19 22:21:47');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (6, '友链', 'link', 'https://veport.oss-cn-beijing.aliyuncs.com/config/8b03884995623eab1a76772f23b58875.jpg', 0,
        '', '2021-08-07 10:32:36', '2022-01-20 21:36:35');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (7, '关于', 'about', 'https://veport.oss-cn-beijing.aliyuncs.com/config/3a4b4e40fb8aa5fcc016f0228938d321.jpg', 0,
        '', '2021-08-07 10:32:36', '2022-01-19 23:10:52');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (8, '留言', 'message', 'https://veport.oss-cn-beijing.aliyuncs.com/config/75e976f3364ba013d62e99ff3ab65d19.jpg',
        0, 'null', '2021-08-07 10:32:36', '2025-05-05 00:51:46');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (9, '个人中心', 'user', 'https://veport.oss-cn-beijing.aliyuncs.com/config/4e319068b295ca52080979d5653c334d.jpg',
        0, '', '2021-08-07 10:32:36', '2022-01-19 22:21:03');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (10, '文章列表', 'article',
        'https://veport.oss-cn-beijing.aliyuncs.com/config/3a4b4e40fb8aa5fcc016f0228938d321.jpg', 0, '',
        '2021-08-10 15:36:19', '2024-10-06 00:14:03');
INSERT INTO `t_page` (`id`, `page_name`, `page_label`, `page_cover`, `is_carousel`, `carousel_covers`, `created_at`,
                      `updated_at`)
VALUES (11, '说说', 'talk', 'https://veport.oss-cn-beijing.aliyuncs.com/config/f9fa18da262910eb13f802b003147915.jpg', 0,
        '', '2022-01-23 00:51:24', '2022-02-11 12:23:15');
COMMIT;

-- ----------------------------
-- Table structure for t_photo
-- ----------------------------
DROP TABLE IF EXISTS `t_photo`;
CREATE TABLE `t_photo`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `album_id`   int          NOT NULL DEFAULT '0' COMMENT '相册id',
    `photo_name` varchar(32)  NOT NULL DEFAULT '' COMMENT '照片名',
    `photo_desc` varchar(64)  NOT NULL DEFAULT '' COMMENT '照片描述',
    `photo_src`  varchar(255) NOT NULL DEFAULT '' COMMENT '照片地址',
    `is_delete`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='照片';

-- ----------------------------
-- Records of t_photo
-- ----------------------------
BEGIN;
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (1, 1, '', '', 'https://static.veweiyi.cn/blog/photo/xuanwu-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (2, 1, '', '', 'https://static.veweiyi.cn/blog/photo/zhuque-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (3, 1, '', '', 'https://static.veweiyi.cn/blog/photo/wusheng-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (4, 1, '', '', 'https://static.veweiyi.cn/blog/photo/qinglong-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (5, 1, '', '', 'https://static.veweiyi.cn/blog/photo/qilin-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
INSERT INTO `t_photo` (`id`, `album_id`, `photo_name`, `photo_desc`, `photo_src`, `is_delete`, `created_at`,
                       `updated_at`)
VALUES (6, 1, '', '', 'https://static.veweiyi.cn/blog/photo/baihu-20250502030854.jpg', 0, '2025-05-02 03:09:12',
        '2025-05-02 03:09:12');
COMMIT;

-- ----------------------------
-- Table structure for t_remark
-- ----------------------------
DROP TABLE IF EXISTS `t_remark`;
CREATE TABLE `t_remark`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`         varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
    `terminal_id`     varchar(255) NOT NULL DEFAULT '' COMMENT '终端id',
    `message_content` varchar(255) NOT NULL DEFAULT '' COMMENT '留言内容',
    `ip_address`      varchar(64)  NOT NULL DEFAULT '' COMMENT '用户ip 127.0.0.1',
    `ip_source`       varchar(255) NOT NULL DEFAULT '' COMMENT '用户地址 广东省深圳市',
    `status`          int          NOT NULL DEFAULT '0' COMMENT '状态:0正常 1编辑 2撤回 3删除',
    `is_review`       tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核通过',
    `created_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
    `updated_at`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='留言';

-- ----------------------------
-- Records of t_remark
-- ----------------------------
BEGIN;
INSERT INTO `t_remark` (`id`, `user_id`, `terminal_id`, `message_content`, `ip_address`, `ip_source`, `status`,
                        `is_review`, `created_at`, `updated_at`)
VALUES (1, '', 'test', '游客测试留言', '127.0.0.1:61776', '本机地址', 0, 1, '2025-05-05 00:53:25',
        '2025-05-05 02:02:43');
INSERT INTO `t_remark` (`id`, `user_id`, `terminal_id`, `message_content`, `ip_address`, `ip_source`, `status`,
                        `is_review`, `created_at`, `updated_at`)
VALUES (2, 'veweiyi', 'veweiyi', '用户测试留言', '127.0.0.1:64964', '本机地址', 0, 1, '2025-05-05 01:03:27',
        '2025-05-05 02:02:38');
COMMIT;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role`
(
    `id`           int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id`    int         NOT NULL DEFAULT '0' COMMENT '父角色id',
    `role_key`     varchar(64) NOT NULL DEFAULT '' COMMENT '角色标识',
    `role_label`   varchar(64) NOT NULL DEFAULT '' COMMENT '角色标签',
    `role_comment` varchar(64) NOT NULL DEFAULT '' COMMENT '角色备注',
    `is_disable`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用  0否 1是',
    `is_default`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认角色 0否 1是',
    `created_at`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- ----------------------------
-- Records of t_role
-- ----------------------------
BEGIN;
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_disable`, `is_default`,
                      `created_at`, `updated_at`)
VALUES (1, 0, 'super-admin', '超级管理员', '超级管理员拥有所有权限', 0, 0, '2021-03-22 14:10:21',
        '2025-05-02 04:07:32');
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_disable`, `is_default`,
                      `created_at`, `updated_at`)
VALUES (2, 0, 'admin', '系统管理员', '系统管理员拥有修改页面权限', 0, 0, '2025-05-02 04:07:16', '2025-05-02 04:25:45');
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_disable`, `is_default`,
                      `created_at`, `updated_at`)
VALUES (3, 0, 'user', '用户', '用户拥有浏览页面权限', 0, 1, '2025-05-02 04:08:21', '2025-05-02 04:08:21');
INSERT INTO `t_role` (`id`, `parent_id`, `role_key`, `role_label`, `role_comment`, `is_disable`, `is_default`,
                      `created_at`, `updated_at`)
VALUES (4, 0, 'test', '测试员', '测试角色', 1, 0, '2025-05-02 04:08:59', '2025-05-02 04:09:26');
COMMIT;

-- ----------------------------
-- Table structure for t_role_api
-- ----------------------------
DROP TABLE IF EXISTS `t_role_api`;
CREATE TABLE `t_role_api`
(
    `id`      int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
    `api_id`  int NOT NULL DEFAULT '0' COMMENT '接口id',
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
CREATE TABLE `t_role_menu`
(
    `id`      int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
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
CREATE TABLE `t_tag`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `tag_name`   varchar(32) NOT NULL DEFAULT '' COMMENT '标签名',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_name` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='标签';

-- ----------------------------
-- Records of t_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_tag` (`id`, `tag_name`, `created_at`, `updated_at`)
VALUES (1, '测试标签', '2024-11-15 17:46:29', '2024-11-15 17:46:29');
INSERT INTO `t_tag` (`id`, `tag_name`, `created_at`, `updated_at`)
VALUES (2, '网站搭建', '2025-05-05 00:20:03', '2025-05-05 00:20:03');
COMMIT;

-- ----------------------------
-- Table structure for t_talk
-- ----------------------------
DROP TABLE IF EXISTS `t_talk`;
CREATE TABLE `t_talk`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT '说说id',
    `user_id`    varchar(64)   NOT NULL DEFAULT '' COMMENT '用户id',
    `content`    varchar(2048) NOT NULL DEFAULT '' COMMENT '说说内容',
    `images`     varchar(2048) NOT NULL DEFAULT '' COMMENT '图片',
    `is_top`     tinyint       NOT NULL DEFAULT '0' COMMENT '是否置顶',
    `status`     tinyint       NOT NULL DEFAULT '1' COMMENT '状态 1.公开 2.私密',
    `like_count` int           NOT NULL DEFAULT '0' COMMENT '点赞数',
    `created_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='说说';

-- ----------------------------
-- Records of t_talk
-- ----------------------------
BEGIN;
INSERT INTO `t_talk` (`id`, `user_id`, `content`, `images`, `is_top`, `status`, `like_count`, `created_at`,
                      `updated_at`)
VALUES (1, '1',
        '测试说说<img src=\"https://static.veweiyi.cn/emoji/qq/14@2x.gif\" width=\"24\" height=\"24\" alt=\"[微笑]\" style=\"margin: 0 1px;display: inline;vertical-align: text-bottom\">',
        'null', 1, 1, 0, '2024-11-16 00:33:43', '2024-11-16 00:39:15');
INSERT INTO `t_talk` (`id`, `user_id`, `content`, `images`, `is_top`, `status`, `like_count`, `created_at`,
                      `updated_at`)
VALUES (2, 'admin', '快来搭建一个属于自己的个人博客吧~', 'null', 1, 1, 0, '2025-05-02 04:13:41', '2025-05-02 04:13:41');
COMMIT;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    varchar(64)   NOT NULL COMMENT '用户id',
    `username`   varchar(64)   NOT NULL DEFAULT '' COMMENT '用户名',
    `password`   varchar(128)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `nickname`   varchar(64)   NOT NULL DEFAULT '' COMMENT '用户昵称',
    `avatar`     varchar(255)  NOT NULL DEFAULT '' COMMENT '用户头像',
    `email`      varchar(64)   NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone`      varchar(64)   NOT NULL DEFAULT '' COMMENT '手机号',
    `info`       varchar(1024) NOT NULL DEFAULT '' COMMENT '用户信息',
    `status`     tinyint       NOT NULL DEFAULT '0' COMMENT '状态: -1删除 0正常 1禁用',
    `login_type` varchar(64)   NOT NULL DEFAULT '' COMMENT '注册方式',
    `ip_address` varchar(255)  NOT NULL DEFAULT '' COMMENT '注册ip',
    `ip_source`  varchar(255)  NOT NULL DEFAULT '' COMMENT '注册ip 源',
    `created_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_uid` (`user_id`) USING BTREE,
    UNIQUE KEY `uk_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录信息';

-- ----------------------------
-- Records of t_user
-- ----------------------------
BEGIN;
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (1, 'root', 'root', '$2a$10$2FQhHyejaB998v1GBVUQYu8MiLPdrgnDP1ozltfa1.LsWD6.P.A/.', '超级管理员',
        'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'root@qq.com', '',
        '{\"gender\":0,\"intro\":\"hello!\",\"website\":\"https://blog.veweiyi.cn\"}', 0, 'email', '127.0.0.1',
        '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2025-05-06 00:35:54');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (2, 'admin', 'admin', '$2a$10$M8EFxmvlxYa9FO2LQOTZtenZ/mZOe89.g3cN0u40LMUGo9yLy9RKa', '管理员',
        'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'admin@qq.com', '',
        '{\"gender\":1,\"intro\":\"hello!\",\"website\":\"https://blog.veweiyi.cn\"}', 0, 'email', '127.0.0.1',
        '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2025-05-03 16:19:13');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (3, 'test', 'test', '$2a$10$eocUk6R87VPE06/iUWFuw.8LOvlRqEl0D7pAS5G2F2H8N44fMgT/a', '测试用户',
        'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', 'test@qq.com', '',
        '{\"gender\":0,\"intro\":\"hello!\",\"website\":\"https://blog.veweiyi.cn\"}', 0, 'email', '127.0.0.1',
        '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2025-05-03 16:19:07');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (4, 'veweiyi', 'veweiyi', '$2a$10$3L1.S0Ja3Oc7QEy1vznuRuYb3yj8WnTUjo2pQZkWvjAreF5ggVq3S', '与梦',
        'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', '791422171@qq.com', '',
        '{\"gender\":0,\"intro\":\"hello!\",\"website\":\"https://blog.veweiyi.cn\"}', 0, 'email', '127.0.0.1',
        '广西壮族自治区梧州市 移动', '2024-07-10 16:24:50', '2025-05-05 03:39:58');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (5, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0',
        '$2a$10$pVvAu.0QFv.dF.bvpEr6mO9gzptzHjesFpD93uNfp/nGJVaiQE8NO', 've-weiyi',
        'https://gitee.com/assets/no_portrait.png', '', '', '', 0, 'gitee', '127.0.0.1:55220', '本机地址',
        '2025-05-05 06:17:24', '2025-05-05 06:17:24');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (6, '582d1d06-e496-4d87-888f-9910c09a9149', '582d1d06-e496-4d87-888f-9910c09a9149',
        '$2a$10$5IK16qTT3LQH7DyfHsfBvO7t26afrdjn/yPmWPl8SdFaQnbcH7dFK', 've-weiyi',
        'https://avatars.githubusercontent.com/u/67481255?v=4', '', '', '', 0, 'github', '127.0.0.1:55844', '本机地址',
        '2025-05-05 06:18:49', '2025-05-05 06:18:49');
INSERT INTO `t_user` (`id`, `user_id`, `username`, `password`, `nickname`, `avatar`, `email`, `phone`, `info`, `status`,
                      `login_type`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (7, '4bd72658-757d-418c-ad0a-55371db6cb16', '919390162@qq.com',
        '$2a$10$Ra1K3zunhmnbLaHuNGEAI.OPZkOP1K6KdFQTUzPfCd1E1LjzrJ8YK', '919390162@qq.com',
        'https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG', '919390162@qq.com', '', '', 0,
        'email', '127.0.0.1:52934', '本机地址', '2025-05-05 23:45:00', '2025-05-05 23:45:00');
COMMIT;

-- ----------------------------
-- Table structure for t_user_oauth
-- ----------------------------
DROP TABLE IF EXISTS `t_user_oauth`;
CREATE TABLE `t_user_oauth`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    varchar(64)                                                   NOT NULL DEFAULT '' COMMENT '用户id',
    `platform`   varchar(64)                                                   NOT NULL DEFAULT '' COMMENT '平台:手机号、邮箱、微信、飞书',
    `open_id`    varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方平台id，标识唯一用户',
    `nickname`   varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '第三方平台昵称',
    `avatar`     varchar(256)                                                  NOT NULL DEFAULT '' COMMENT '第三方平台头像',
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_uid_plat` (`user_id`,`platform`) USING BTREE,
    UNIQUE KEY `uk_oid_plat` (`open_id`,`platform`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='第三方登录信息';

-- ----------------------------
-- Records of t_user_oauth
-- ----------------------------
BEGIN;
INSERT INTO `t_user_oauth` (`id`, `user_id`, `platform`, `open_id`, `nickname`, `avatar`, `created_at`, `updated_at`)
VALUES (7, 'root', 'gitee', '7705905', 've-weiyi',
        'https://foruda.gitee.com/avatar/1746453514161082773/7705905_wy791422171_1746453514.png', '2025-05-06 00:34:26',
        '2025-05-06 00:34:26');
COMMIT;

-- ----------------------------
-- Table structure for t_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_user_role`;
CREATE TABLE `t_user_role`
(
    `id`      int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
    `role_id` int         NOT NULL DEFAULT '0' COMMENT '角色id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户-角色关联';

-- ----------------------------
-- Records of t_user_role
-- ----------------------------
BEGIN;
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (2, 'root', 1);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (3, 'admin', 2);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (4, 'test', 4);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (5, 'veweiyi', 1);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (6, '95a64673-4bd5-4a4b-83d1-3095f4a0e1c0', 3);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (7, '582d1d06-e496-4d87-888f-9910c09a9149', 3);
INSERT INTO `t_user_role` (`id`, `user_id`, `role_id`)
VALUES (8, '4bd72658-757d-418c-ad0a-55371db6cb16', 3);
COMMIT;

-- ----------------------------
-- Table structure for t_visit_daily_stats
-- ----------------------------
DROP TABLE IF EXISTS `t_visit_daily_stats`;
CREATE TABLE `t_visit_daily_stats`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `date`       varchar(10) NOT NULL DEFAULT '' COMMENT '日期',
    `view_count` int         NOT NULL DEFAULT '0' COMMENT '访问量',
    `visit_type` tinyint     NOT NULL DEFAULT '1' COMMENT '1 访客数 2 浏览数',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_date_type` (`date`,`visit_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面访问数量';

-- ----------------------------
-- Records of t_visit_daily_stats
-- ----------------------------
BEGIN;
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (1, '2024-12-04', 6, 1, '2024-12-04 12:20:22', '2024-12-04 17:36:06');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (2, '2024-12-03', 1, 1, '2024-12-04 14:48:04', '2024-12-04 14:48:04');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (3, '2024-12-05', 36, 1, '2024-12-05 11:16:07', '2024-12-05 15:31:26');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (4, '2024-12-06', 3, 1, '2024-12-06 15:04:42', '2024-12-06 15:52:43');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (5, '2024-12-11', 1, 1, '2024-12-11 12:04:01', '2024-12-11 12:04:01');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (6, '2024-12-13', 4, 1, '2024-12-13 14:34:20', '2024-12-13 18:33:07');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (7, '2024-12-14', 3, 1, '2024-12-14 02:34:07', '2024-12-14 20:52:43');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (8, '2024-12-16', 16, 1, '2024-12-16 09:54:06', '2024-12-16 18:16:28');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (9, '2024-12-17', 5, 1, '2024-12-17 13:36:39', '2024-12-17 20:35:02');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (10, '2024-12-18', 4, 1, '2024-12-18 09:47:51', '2024-12-18 15:48:09');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (11, '2024-12-19', 3, 1, '2024-12-19 14:57:36', '2024-12-19 20:54:32');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (12, '2024-12-21', 2, 1, '2024-12-21 07:34:12', '2024-12-21 10:16:14');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (13, '2024-12-22', 1, 1, '2024-12-22 16:42:20', '2024-12-22 16:42:20');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (14, '2024-12-24', 1, 1, '2024-12-24 20:43:21', '2024-12-24 20:43:21');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (15, '2024-12-25', 3, 1, '2024-12-25 00:14:34', '2024-12-25 18:13:48');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (16, '2024-12-27', 1, 1, '2024-12-27 14:11:58', '2024-12-27 14:11:58');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (17, '2024-12-28', 1, 1, '2024-12-28 20:04:04', '2024-12-28 20:04:04');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (18, '2024-12-29', 1, 1, '2024-12-29 11:10:15', '2024-12-29 11:10:15');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (19, '2024-12-30', 6, 1, '2024-12-30 06:21:08', '2024-12-30 22:09:33');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (20, '2025-01-02', 3, 1, '2025-01-02 06:44:14', '2025-01-02 10:19:25');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (21, '2025-01-03', 2, 1, '2025-01-03 09:25:20', '2025-01-03 09:25:39');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (22, '2025-01-04', 1, 1, '2025-01-04 00:51:14', '2025-01-04 00:51:14');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (23, '2025-01-05', 6, 1, '2025-01-05 00:03:19', '2025-01-05 18:33:54');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (24, '2025-01-06', 1, 1, '2025-01-06 09:09:39', '2025-01-06 09:09:39');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (25, '2025-01-07', 1, 1, '2025-01-07 12:46:20', '2025-01-07 12:46:20');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (26, '2025-02-07', 2, 1, '2025-02-07 12:18:33', '2025-02-07 21:00:01');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (27, '2025-02-08', 3, 1, '2025-02-08 09:14:24', '2025-02-08 15:15:12');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (28, '2025-02-09', 4, 1, '2025-02-09 03:24:04', '2025-02-09 22:05:03');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (29, '2025-02-10', 7, 1, '2025-02-10 09:39:28', '2025-02-10 19:02:00');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (30, '2025-02-11', 3, 1, '2025-02-11 08:55:51', '2025-02-11 21:21:17');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (31, '2025-02-12', 1, 1, '2025-02-12 14:24:30', '2025-02-12 14:24:30');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (32, '2025-02-13', 3, 1, '2025-02-13 17:37:03', '2025-02-13 17:45:09');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (33, '2025-02-14', 2, 1, '2025-02-14 14:04:24', '2025-02-14 14:19:40');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (34, '2025-02-15', 7, 1, '2025-02-15 00:43:18', '2025-02-15 20:00:26');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (35, '2025-02-16', 4, 1, '2025-02-16 02:52:39', '2025-02-16 22:16:50');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (36, '2025-02-17', 4, 1, '2025-02-17 03:03:05', '2025-02-17 23:34:32');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (37, '2025-02-18', 14, 1, '2025-02-18 00:03:08', '2025-02-18 17:07:35');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (38, '2025-02-19', 5, 1, '2025-02-19 01:04:21', '2025-02-19 22:24:29');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (39, '2025-02-20', 4, 1, '2025-02-20 07:16:09', '2025-02-20 18:29:43');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (40, '2025-02-21', 6, 1, '2025-02-21 08:15:20', '2025-02-21 22:29:37');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (41, '2025-02-22', 4, 1, '2025-02-22 03:24:08', '2025-02-22 20:08:15');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (42, '2025-02-23', 3, 1, '2025-02-23 03:56:23', '2025-02-23 15:46:31');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (43, '2025-02-24', 4, 1, '2025-02-24 12:24:44', '2025-02-24 12:49:48');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (44, '2025-02-25', 3, 1, '2025-02-25 00:44:27', '2025-02-25 21:45:03');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (45, '2025-02-26', 4, 1, '2025-02-26 03:05:27', '2025-02-26 23:17:18');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (46, '2025-02-27', 1, 1, '2025-02-27 20:36:43', '2025-02-27 20:36:43');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (47, '2025-02-28', 2, 1, '2025-02-28 10:06:54', '2025-02-28 13:38:13');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (48, '2025-03-01', 6, 1, '2025-03-01 03:25:33', '2025-03-01 19:37:36');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (49, '2025-03-02', 3, 1, '2025-03-02 07:43:00', '2025-03-02 15:50:21');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (50, '2025-03-03', 4, 1, '2025-03-03 01:52:51', '2025-03-03 23:53:15');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (51, '2025-03-04', 7, 1, '2025-03-04 07:17:39', '2025-03-04 20:06:38');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (52, '2025-03-05', 1, 1, '2025-03-05 13:43:02', '2025-03-05 13:43:02');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (53, '2025-03-10', 3, 1, '2025-03-10 17:19:28', '2025-03-10 17:30:52');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (54, '2025-03-11', 4, 1, '2025-03-11 02:06:25', '2025-03-11 21:42:51');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (55, '2025-03-12', 5, 1, '2025-03-12 00:57:41', '2025-03-12 22:18:05');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (56, '2025-03-13', 4, 1, '2025-03-13 07:44:17', '2025-03-13 10:30:29');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (57, '2025-03-14', 5, 1, '2025-03-14 00:50:46', '2025-03-14 21:41:54');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (58, '2025-03-15', 4, 1, '2025-03-15 03:34:28', '2025-03-15 10:05:54');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (59, '2025-03-16', 1, 1, '2025-03-16 14:38:40', '2025-03-16 14:38:40');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (60, '2025-03-17', 2, 1, '2025-03-17 10:42:48', '2025-03-17 23:36:47');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (61, '2025-03-18', 7, 1, '2025-03-18 05:24:03', '2025-03-18 16:13:27');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (62, '2025-03-19', 1, 1, '2025-03-19 17:11:59', '2025-03-19 17:11:59');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (63, '2025-03-20', 2, 1, '2025-03-20 19:55:49', '2025-03-20 22:01:10');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (64, '2025-03-22', 7, 1, '2025-03-22 07:12:42', '2025-03-22 23:27:20');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (65, '2025-03-23', 1, 1, '2025-03-23 03:48:13', '2025-03-23 03:48:13');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (66, '2025-03-24', 3, 1, '2025-03-24 05:33:25', '2025-03-24 15:18:16');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (67, '2025-03-25', 6, 1, '2025-03-25 02:15:58', '2025-03-25 23:05:01');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (68, '2025-03-26', 1, 1, '2025-03-26 02:52:13', '2025-03-26 02:52:13');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (69, '2025-03-27', 18, 1, '2025-03-27 07:10:37', '2025-03-27 17:44:50');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (70, '2025-03-28', 6, 1, '2025-03-28 00:11:24', '2025-03-28 19:14:32');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (71, '2025-03-29', 6, 1, '2025-03-29 00:44:17', '2025-03-29 09:14:09');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (72, '2025-03-30', 1, 1, '2025-03-30 07:24:26', '2025-03-30 07:24:26');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (73, '2025-03-31', 2, 1, '2025-03-31 04:06:56', '2025-03-31 17:41:30');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (74, '2025-04-02', 5, 1, '2025-04-02 07:41:17', '2025-04-02 12:34:38');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (75, '2025-04-03', 2, 1, '2025-04-03 00:10:21', '2025-04-03 10:59:45');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (76, '2025-04-04', 2, 1, '2025-04-04 03:59:19', '2025-04-04 23:06:39');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (77, '2025-04-05', 3, 1, '2025-04-05 02:32:17', '2025-04-05 12:04:54');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (78, '2025-04-06', 2, 1, '2025-04-06 06:54:54', '2025-04-06 16:48:56');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (79, '2025-04-07', 5, 1, '2025-04-07 10:13:43', '2025-04-07 18:49:53');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (80, '2025-04-08', 4, 1, '2025-04-08 01:01:27', '2025-04-08 11:45:03');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (81, '2025-04-09', 14, 1, '2025-04-09 06:20:11', '2025-04-09 23:59:00');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (82, '2025-04-10', 6, 1, '2025-04-10 08:06:03', '2025-04-10 21:12:48');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (83, '2025-04-11', 3, 1, '2025-04-11 13:02:46', '2025-04-11 23:57:07');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (84, '2025-04-12', 5, 1, '2025-04-12 04:02:24', '2025-04-12 18:03:16');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (85, '2025-04-13', 2, 1, '2025-04-13 01:01:52', '2025-04-13 22:01:51');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (86, '2025-04-14', 4, 1, '2025-04-14 03:02:22', '2025-04-14 23:09:17');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (87, '2025-04-15', 12, 1, '2025-04-15 12:04:58', '2025-04-15 19:13:46');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (88, '2025-04-16', 18, 1, '2025-04-16 01:16:56', '2025-04-16 17:44:34');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (89, '2025-04-17', 16, 1, '2025-04-17 01:48:31', '2025-04-17 23:33:40');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (90, '2025-04-18', 8, 1, '2025-04-18 00:33:26', '2025-04-18 19:52:45');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (91, '2025-04-19', 21, 1, '2025-04-19 00:43:26', '2025-04-19 23:24:55');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (92, '2025-04-20', 5, 1, '2025-04-20 04:03:20', '2025-04-20 23:42:45');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (93, '2025-04-21', 24, 1, '2025-04-21 01:43:13', '2025-04-21 23:16:19');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (94, '2025-04-22', 14, 1, '2025-04-22 02:13:00', '2025-04-22 20:56:29');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (95, '2025-04-23', 12, 1, '2025-04-23 00:35:42', '2025-04-23 20:43:47');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (96, '2025-04-24', 4, 1, '2025-04-24 00:04:25', '2025-04-24 23:41:24');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (97, '2025-04-25', 21, 1, '2025-04-25 06:23:03', '2025-04-25 23:54:39');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (98, '2025-04-26', 14, 1, '2025-04-26 02:03:32', '2025-04-26 22:23:07');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (99, '2025-04-27', 14, 1, '2025-04-27 00:00:06', '2025-04-27 14:11:23');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (100, '2025-04-29', 5, 1, '2025-04-29 17:14:05', '2025-04-29 22:05:22');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (101, '2025-04-30', 12, 1, '2025-04-30 03:43:02', '2025-04-30 17:40:55');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (102, '2025-05-01', 32, 1, '2025-05-01 02:33:07', '2025-05-01 23:52:45');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (111, '2025-05-01', 3, 2, '2025-05-01 23:52:45', '2025-05-01 23:53:30');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (112, '2025-05-02', 1, 1, '2025-05-02 00:28:31', '2025-05-02 00:28:31');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (113, '2025-05-02', 82, 2, '2025-05-02 00:28:31', '2025-05-02 04:42:36');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (114, '2025-05-05', 4, 1, '2025-05-05 00:16:38', '2025-05-05 06:17:23');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (115, '2025-05-05', 82, 2, '2025-05-05 00:16:38', '2025-05-05 23:57:57');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (116, '2025-05-06', 1, 1, '2025-05-06 00:08:58', '2025-05-06 00:08:58');
INSERT INTO `t_visit_daily_stats` (`id`, `date`, `view_count`, `visit_type`, `created_at`, `updated_at`)
VALUES (117, '2025-05-06', 33, 2, '2025-05-06 00:08:58', '2025-05-06 00:43:28');
COMMIT;

-- ----------------------------
-- Table structure for t_visit_log
-- ----------------------------
DROP TABLE IF EXISTS `t_visit_log`;
CREATE TABLE `t_visit_log`
(
    `id`          int                                                          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`     varchar(64)                                                  NOT NULL DEFAULT '' COMMENT '用户id',
    `terminal_id` varchar(64)                                                  NOT NULL DEFAULT '' COMMENT '设备id',
    `page_name`   varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '访问页面',
    `ip_address`  varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '操作ip',
    `ip_source`   varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '操作地址',
    `os`          varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作系统',
    `browser`     varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '浏览器',
    `created_at`  datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of t_visit_log
-- ----------------------------
BEGIN;
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (1, '', '8c4a5cc5bc9976096cd96f7cf2273b84', 'category', '127.0.0.1', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-01 21:57:20', '2025-05-01 21:57:20');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (2, '', '8c4a5cc5bc9976096cd96f7cf2273b84', 'tag', '127.0.0.1', '本机地址', 'Intel Mac OS X 10_15_7', 'Chrome',
        '2025-05-01 21:57:25', '2025-05-01 21:57:25');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (3, '', '8c4a5cc5bc9976096cd96f7cf2273b84', 'friend', '127.0.0.1', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-01 21:57:33', '2025-05-01 21:57:33');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (4, '', '8c4a5cc5bc9976096cd96f7cf2273b84', 'remark', '127.0.0.1', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-01 21:57:44', '2025-05-01 21:57:44');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (5, '', '8c4a5cc5bc9976096cd96f7cf2273b84', 'remark', '127.0.0.1', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-01 22:03:59', '2025-05-01 22:03:59');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (6, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:65035', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 01:03:33', '2025-05-05 01:03:33');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (7, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:56240', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 01:25:58', '2025-05-05 01:25:58');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (8, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:51866', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:02:14', '2025-05-05 02:02:14');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (9, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:52063', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:02:47', '2025-05-05 02:02:47');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (10, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:54234', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:09:10', '2025-05-05 02:09:10');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (11, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:54867', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:11:00', '2025-05-05 02:11:00');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (12, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:55016', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:11:21', '2025-05-05 02:11:21');
INSERT INTO `t_visit_log` (`id`, `user_id`, `terminal_id`, `page_name`, `ip_address`, `ip_source`, `os`, `browser`,
                           `created_at`, `updated_at`)
VALUES (13, 'veweiyi', '57f62293a7c250c4c4c27542eb530a05', '', '127.0.0.1:55162', '本机地址', 'Intel Mac OS X 10_15_7',
        'Chrome', '2025-05-05 02:11:46', '2025-05-05 02:11:46');
COMMIT;

-- ----------------------------
-- Table structure for t_visitor
-- ----------------------------
DROP TABLE IF EXISTS `t_visitor`;
CREATE TABLE `t_visitor`
(
    `id`          int                                                          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `terminal_id` varchar(64)                                                  NOT NULL DEFAULT '' COMMENT '设备id',
    `os`          varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作系统',
    `browser`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '浏览器',
    `ip_address`  varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '操作ip',
    `ip_source`   varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '操作地址',
    `created_at`  datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_tid` (`terminal_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of t_visitor
-- ----------------------------
BEGIN;
INSERT INTO `t_visitor` (`id`, `terminal_id`, `os`, `browser`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (1, '57f62293a7c250c4c4c27542eb530a05', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1',
        '广西壮族自治区梧州市 移动', '2025-05-02 00:28:31', '2025-05-02 03:01:45');
INSERT INTO `t_visitor` (`id`, `terminal_id`, `os`, `browser`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (9, '7efe2cbad1787f5efd1c2089c588186c', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:51753', '本机地址',
        '2025-05-05 06:07:42', '2025-05-05 06:07:42');
INSERT INTO `t_visitor` (`id`, `terminal_id`, `os`, `browser`, `ip_address`, `ip_source`, `created_at`, `updated_at`)
VALUES (10, 'eefa656b1f6ca0d7dd911fec6b42e929', 'Intel Mac OS X 10_15_7', 'Chrome', '127.0.0.1:55218', '本机地址',
        '2025-05-05 06:17:23', '2025-05-05 06:17:23');
COMMIT;

-- ----------------------------
-- Table structure for t_website_config
-- ----------------------------
DROP TABLE IF EXISTS `t_website_config`;
CREATE TABLE `t_website_config`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `key`        varchar(32)   NOT NULL DEFAULT '' COMMENT '关键词',
    `config`     varchar(2048) NOT NULL DEFAULT '' COMMENT '配置信息',
    `created_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='网站配置表';

-- ----------------------------
-- Records of t_website_config
-- ----------------------------
BEGIN;
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`)
VALUES (1, 'website_config',
        '{\"admin_url\":\"\",\"alipay_qr_code\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/17f234dc487c1bb5bbb732869be0eb53.jpg\",\"gitee\":\"https://gitee.com/wy791422171\",\"github\":\"https://github.com/ve-weiyi\",\"is_chat_room\":1,\"is_comment_review\":1,\"is_email_notice\":1,\"is_message_review\":0,\"is_music_player\":1,\"is_reward\":0,\"qq\":\"791422171\",\"social_login_list\":[\"github\",\"gitee\",\"qq\",\"wechat\",\"weibo\",\"feishu\"],\"social_url_list\":[\"qq\",\"github\",\"gitee\"],\"tourist_avatar\":\"https://veport.oss-cn-beijing.aliyuncs.com/config/5bfb96809bee5ba80a36811f0bf1d1ea.gif\",\"user_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175820.jpg\",\"website_author\":\"与梦\",\"website_avatar\":\"https://static.veweiyi.cn/blog/website/tiger-20241115175746.jpg\",\"website_create_time\":\"2022-01-17\",\"website_intro\":\"你能做的，岂止如此。\",\"website_name\":\"与梦\",\"website_notice\":\"网站搭建问题请联系QQ 791422171。\",\"website_record_no\":\"桂ICP备2023013735号-1\",\"websocket_url\":\"wss://blog.veweiyi.cn/api/websocket\",\"weixin_qr_code\":\"\"}',
        '2021-08-09 19:37:30', '2025-05-05 23:30:22');
INSERT INTO `t_website_config` (`id`, `key`, `config`, `created_at`, `updated_at`)
VALUES (2, 'about_me', '{\"content\":\"welcome to my blog!\"}', '2024-11-15 17:57:20', '2024-11-15 17:57:20');
COMMIT;

SET
FOREIGN_KEY_CHECKS = 1;
