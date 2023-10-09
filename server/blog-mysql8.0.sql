/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : blog-v2

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 09/10/2023 14:37:01
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`       varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'api名称',
    `path`       varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'api路径',
    `method`     varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT 'api请求方法',
    `parent_id`  int                                                           NOT NULL COMMENT '分组id',
    `traceable`  tinyint                                                       NOT NULL COMMENT '是否追溯操作记录 0需要，1是',
    `status`     tinyint                                                       NOT NULL DEFAULT '1' COMMENT '状态 1开，2关',
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_api_path_method` (`path`,`method`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=502 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='api路由';

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`              int                                                          NOT NULL AUTO_INCREMENT,
    `user_id`         int                                                          NOT NULL COMMENT '作者',
    `category_id`     int                                                                   DEFAULT NULL COMMENT '文章分类',
    `article_cover`   varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci        DEFAULT NULL COMMENT '文章缩略图',
    `article_title`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
    `article_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
    `type`            tinyint                                                      NOT NULL DEFAULT '0' COMMENT '文章类型 1原创 2转载 3翻译',
    `original_url`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         DEFAULT NULL COMMENT '原文链接',
    `is_top`          tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶 0否 1是',
    `is_delete`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
    `status`          tinyint                                                      NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密 3评论可见',
    `created_at`      datetime                                                     NOT NULL COMMENT '发表时间',
    `updated_at`      datetime                                                              DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`
(
    `id`         int NOT NULL AUTO_INCREMENT,
    `article_id` int NOT NULL COMMENT '文章id',
    `tag_id`     int NOT NULL COMMENT '标签id',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `fk_article_tag_1` (`article_id`) USING BTREE,
    KEY          `fk_article_tag_2` (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1060 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) DEFAULT NULL,
    `v0`    varchar(100) DEFAULT NULL,
    `v1`    varchar(100) DEFAULT NULL,
    `v2`    varchar(100) DEFAULT NULL,
    `v3`    varchar(100) DEFAULT NULL,
    `v4`    varchar(100) DEFAULT NULL,
    `v5`    varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`
(
    `id`            int                                                          NOT NULL AUTO_INCREMENT,
    `category_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名',
    `created_at`    datetime                                                     NOT NULL COMMENT '创建时间',
    `updated_at`    datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=194 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record`
(
    `id`         int                                                            NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`    int      DEFAULT NULL COMMENT '用户id',
    `nickname`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci   NOT NULL COMMENT '昵称',
    `avatar`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '头像',
    `content`    varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '聊天内容',
    `ip_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci   NOT NULL COMMENT 'ip地址',
    `ip_source`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT 'ip来源',
    `type`       int                                                            NOT NULL COMMENT '类型',
    `created_at` datetime                                                       NOT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2929 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`              int      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`         int      NOT NULL COMMENT '评论用户Id',
    `topic_id`        int               DEFAULT NULL COMMENT '评论主题id',
    `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
    `reply_user_id`   int               DEFAULT NULL COMMENT '回复用户id',
    `parent_id`       int               DEFAULT NULL COMMENT '父评论id',
    `type`            tinyint  NOT NULL COMMENT '评论类型 1.文章 2.友链 3.说说',
    `is_delete`       tinyint  NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
    `is_review`       tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核',
    `created_at`      datetime NOT NULL COMMENT '评论时间',
    `updated_at`      datetime          DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY               `fk_comment_user` (`user_id`) USING BTREE,
    KEY               `fk_comment_parent` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=754 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for friend_link
-- ----------------------------
DROP TABLE IF EXISTS `friend_link`;
CREATE TABLE `friend_link`
(
    `id`           int                                                           NOT NULL AUTO_INCREMENT,
    `link_name`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '链接名',
    `link_avatar`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接头像',
    `link_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '链接地址',
    `link_intro`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接介绍',
    `created_at`   datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at`   datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `fk_friend_link_user` (`link_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`
(
    `id`         int                                                          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单名',
    `path`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
    `component`  varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '组件',
    `icon`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单icon',
    `rank`       tinyint                                                      NOT NULL DEFAULT '0' COMMENT '排序',
    `parent_id`  int                                                                   DEFAULT '0' COMMENT '父id',
    `is_hidden`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏  0否1是',
    `created_at` datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=242 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`
(
    `id`             int                                                           NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `opt_module`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '操作模块',
    `opt_type`       varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '操作类型',
    `opt_method`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作方法',
    `opt_desc`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作描述',
    `cost`           varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci   DEFAULT NULL COMMENT '耗时（ms）',
    `status`         int                                                            DEFAULT '0' COMMENT '响应状态码',
    `request_url`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作url',
    `request_method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '请求方式',
    `request_header` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '请求头',
    `request_param`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求参数',
    `response_data`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '返回数据',
    `user_id`        int                                                           NOT NULL COMMENT '用户id',
    `nickname`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户昵称',
    `ip_address`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作ip',
    `ip_source`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作地址',
    `created_at`     datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at`     datetime                                                       DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1269 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for page
-- ----------------------------
DROP TABLE IF EXISTS `page`;
CREATE TABLE `page`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT COMMENT '页面id',
    `page_name`  varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '页面名',
    `page_label` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '页面标签',
    `page_cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '页面封面',
    `created_at` datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=905 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='页面';

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT COMMENT '主键',
    `album_id`   int                                                           NOT NULL COMMENT '相册id',
    `photo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '照片名',
    `photo_desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '照片描述',
    `photo_src`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '照片地址',
    `is_delete`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
    `created_at` datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='照片';

-- ----------------------------
-- Table structure for photo_album
-- ----------------------------
DROP TABLE IF EXISTS `photo_album`;
CREATE TABLE `photo_album`
(
    `id`          int                                                           NOT NULL AUTO_INCREMENT COMMENT '主键',
    `album_name`  varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '相册名',
    `album_desc`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '相册描述',
    `album_cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '相册封面',
    `is_delete`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
    `status`      tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密',
    `created_at`  datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at`  datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='相册';

-- ----------------------------
-- Table structure for remark
-- ----------------------------
DROP TABLE IF EXISTS `remark`;
CREATE TABLE `remark`
(
    `id`              int                                                           NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `nickname`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '昵称',
    `avatar`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
    `message_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '留言内容',
    `ip_address`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户ip',
    `ip_source`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户地址',
    `time`            int      DEFAULT NULL COMMENT '弹幕速度',
    `is_review`       tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核',
    `created_at`      datetime                                                      NOT NULL COMMENT '发布时间',
    `updated_at`      datetime DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3918 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`           int                                                          NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_pid`     int                                                          NOT NULL DEFAULT '0' COMMENT '父角色id',
    `role_domain`  varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '角色域',
    `role_name`    varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色名',
    `role_comment` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色备注',
    `is_disable`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用  0否 1是',
    `is_default`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认角色 0否 1是',
    `created_at`   datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- ----------------------------
-- Table structure for role_api
-- ----------------------------
DROP TABLE IF EXISTS `role_api`;
CREATE TABLE `role_api`
(
    `id`      int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
    `api_id`  int NOT NULL DEFAULT '0' COMMENT '接口id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=470 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-api关联';

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`
(
    `id`      int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
    `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2980 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色-菜单关联';

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`
(
    `id`         int                                                          NOT NULL AUTO_INCREMENT,
    `tag_name`   varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签名',
    `created_at` datetime                                                     NOT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for talk
-- ----------------------------
DROP TABLE IF EXISTS `talk`;
CREATE TABLE `talk`
(
    `id`         int                                                            NOT NULL AUTO_INCREMENT COMMENT '说说id',
    `user_id`    int                                                            NOT NULL COMMENT '用户id',
    `content`    varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '说说内容',
    `images`     varchar(2500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci          DEFAULT NULL COMMENT '图片',
    `is_top`     tinyint                                                        NOT NULL DEFAULT '0' COMMENT '是否置顶',
    `status`     tinyint                                                        NOT NULL DEFAULT '1' COMMENT '状态 1.公开 2.私密',
    `created_at` datetime                                                       NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                                DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for unique_view
-- ----------------------------
DROP TABLE IF EXISTS `unique_view`;
CREATE TABLE `unique_view`
(
    `id`          int      NOT NULL AUTO_INCREMENT,
    `views_count` int      NOT NULL COMMENT '访问量',
    `created_at`  datetime NOT NULL COMMENT '创建时间',
    `updated_at`  datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=703 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for upload
-- ----------------------------
DROP TABLE IF EXISTS `upload`;
CREATE TABLE `upload`
(
    `id`         bigint                                                        NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    int                                                           NOT NULL DEFAULT '0' COMMENT '用户id',
    `label`      varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '标签',
    `file_name`  varchar(64)                                                   NOT NULL DEFAULT '' COMMENT '文件名称',
    `file_size`  int                                                           NOT NULL DEFAULT '0' COMMENT '文件大小',
    `file_md5`   varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '文件md5值',
    `file_url`   varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '上传路径',
    `created_at` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户灯效信息';

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account`
(
    `id`            int                                                          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`      varchar(64)                                                  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`      varchar(128)                                                 NOT NULL DEFAULT '' COMMENT '密码',
    `status`        tinyint                                                      NOT NULL DEFAULT '0' COMMENT '状态: -1删除 0正常 1禁用',
    `register_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '注册方式',
    `ip_address`    varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '注册ip',
    `ip_source`     varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '注册ip 源',
    `created_at`    datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录信息';

-- ----------------------------
-- Table structure for user_information
-- ----------------------------
DROP TABLE IF EXISTS `user_information`;
CREATE TABLE `user_information`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    int                                                           NOT NULL DEFAULT '0' COMMENT '用户id',
    `email`      varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `nickname`   varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '用户昵称',
    `avatar`     varchar(1024)                                                 NOT NULL DEFAULT '' COMMENT '用户头像',
    `phone`      varchar(32)                                                   NOT NULL DEFAULT '' COMMENT '用户手机号',
    `intro`      varchar(255)                                                  NOT NULL DEFAULT '' COMMENT '个人简介',
    `website`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个人网站',
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';

-- ----------------------------
-- Table structure for user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `user_login_history`;
CREATE TABLE `user_login_history`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    int                                                           NOT NULL DEFAULT '0' COMMENT '用户id',
    `login_type` varchar(64)                                                   NOT NULL DEFAULT '0' COMMENT '登录类型',
    `agent`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '代理',
    `ip_address` varchar(255)                                                  NOT NULL DEFAULT '' COMMENT 'ip host',
    `ip_source`  varchar(255)                                                  NOT NULL DEFAULT '' COMMENT 'ip 源',
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `uk_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=297 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录历史';

-- ----------------------------
-- Table structure for user_oauth
-- ----------------------------
DROP TABLE IF EXISTS `user_oauth`;
CREATE TABLE `user_oauth`
(
    `id`         int                                                          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    int                                                          NOT NULL DEFAULT '0' COMMENT '用户id',
    `open_id`    varchar(128)                                                 NOT NULL COMMENT '开发平台id，标识唯一用户',
    `platform`   varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '平台:手机号、邮箱、微信、飞书',
    `created_at` datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_oid_plat` (`open_id`,`platform`) USING BTREE,
    KEY          `idx_uuid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='第三方登录信息';

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `id`      int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
    `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户-角色关联';

-- ----------------------------
-- Table structure for website_config
-- ----------------------------
DROP TABLE IF EXISTS `website_config`;
CREATE TABLE `website_config`
(
    `id`         int         NOT NULL AUTO_INCREMENT,
    `key`        varchar(20) NOT NULL COMMENT '关键词',
    `config`     varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '配置信息',
    `created_at` datetime    NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                       DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

SET
FOREIGN_KEY_CHECKS = 1;
