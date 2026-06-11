-- ============================================
-- clean.sql
-- 清理所有日志数据和角色权限数据
-- 执行前请确认已备份数据库！
-- ============================================

SET FOREIGN_KEY_CHECKS = 0;

-- ============================================
-- 日志表
-- ============================================
TRUNCATE TABLE `t_login_log`;
TRUNCATE TABLE `t_operation_log`;
TRUNCATE TABLE `t_upload_log`;
TRUNCATE TABLE `t_visit_log`;

-- ============================================
-- 角色权限表
-- ============================================
TRUNCATE TABLE `t_role_api`;
TRUNCATE TABLE `t_role_menu`;
TRUNCATE TABLE `t_role`;
TRUNCATE TABLE `t_menu`;
TRUNCATE TABLE `t_api`;

SET FOREIGN_KEY_CHECKS = 1;
