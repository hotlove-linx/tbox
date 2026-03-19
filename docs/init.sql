-- TBox 数据库初始化脚本
-- 不使用外键约束，通过应用层保证数据一致性

CREATE DATABASE IF NOT EXISTS `tbox` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `tbox`;

-- ============================================================
-- 用户端表
-- ============================================================

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `nickname` VARCHAR(50) DEFAULT NULL,
  `avatar_url` VARCHAR(500) DEFAULT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `disabled_reason` VARCHAR(500) DEFAULT NULL,
  `disabled_at` DATETIME DEFAULT NULL,
  `last_login_at` DATETIME DEFAULT NULL,
  `last_login_ip` VARCHAR(50) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 分类表
CREATE TABLE IF NOT EXISTS `categories` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `icon` VARCHAR(100) DEFAULT NULL,
  `sort_order` INT NOT NULL DEFAULT 0,
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 标签表
CREATE TABLE IF NOT EXISTS `tags` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `color` VARCHAR(20) DEFAULT NULL,
  `type` VARCHAR(20) NOT NULL DEFAULT 'manual',
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 软件表
CREATE TABLE IF NOT EXISTS `software` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `description` TEXT DEFAULT NULL,
  `detail` TEXT DEFAULT NULL,
  `icon_url` VARCHAR(500) DEFAULT NULL,
  `category_id` BIGINT UNSIGNED DEFAULT NULL,
  `version` VARCHAR(20) DEFAULT NULL,
  `size` BIGINT DEFAULT NULL,
  `file_name` VARCHAR(255) DEFAULT NULL,
  `download_url` VARCHAR(500) DEFAULT NULL,
  `file_hash` VARCHAR(64) DEFAULT NULL,
  `developer` VARCHAR(100) DEFAULT NULL,
  `uploader_id` BIGINT UNSIGNED DEFAULT NULL,
  `visibility` VARCHAR(20) NOT NULL DEFAULT 'public',
  `audit_status` VARCHAR(20) NOT NULL DEFAULT 'pending',
  `audit_remark` VARCHAR(500) DEFAULT NULL,
  `is_recommended` TINYINT(1) NOT NULL DEFAULT 0,
  `is_on_shelf` TINYINT(1) NOT NULL DEFAULT 0,
  `off_shelf_reason` VARCHAR(500) DEFAULT NULL,
  `scan_status` VARCHAR(20) DEFAULT NULL,
  `license` VARCHAR(100) DEFAULT NULL,
  `system_requirement` VARCHAR(100) DEFAULT NULL,
  `rating` DECIMAL(2,1) NOT NULL DEFAULT 0.0,
  `download_count` BIGINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_software_category_id` (`category_id`),
  INDEX `idx_software_uploader_id` (`uploader_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 软件-标签关联表
CREATE TABLE IF NOT EXISTS `software_tags` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `software_id` BIGINT UNSIGNED NOT NULL,
  `tag_id` BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_software_tags_software_id` (`software_id`),
  INDEX `idx_software_tags_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 软件截图表
CREATE TABLE IF NOT EXISTS `software_screenshots` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `software_id` BIGINT UNSIGNED NOT NULL,
  `image_url` VARCHAR(500) DEFAULT NULL,
  `sort_order` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_software_screenshots_software_id` (`software_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 软件版本表
CREATE TABLE IF NOT EXISTS `software_versions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `software_id` BIGINT UNSIGNED NOT NULL,
  `version` VARCHAR(20) DEFAULT NULL,
  `size` BIGINT DEFAULT NULL,
  `file_name` VARCHAR(255) DEFAULT NULL,
  `download_url` VARCHAR(500) DEFAULT NULL,
  `changelog` TEXT DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_software_versions_software_id` (`software_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 软件评价表
CREATE TABLE IF NOT EXISTS `software_reviews` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `software_id` BIGINT UNSIGNED NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `rating` INT NOT NULL,
  `content` TEXT DEFAULT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'normal',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_software_reviews_software_id` (`software_id`),
  INDEX `idx_software_reviews_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户空间表
CREATE TABLE IF NOT EXISTS `user_spaces` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `total_capacity` BIGINT NOT NULL DEFAULT 5368709120 COMMENT '默认5GB',
  `used_capacity` BIGINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_user_spaces_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 管理后台表
-- ============================================================

-- 管理员表
CREATE TABLE IF NOT EXISTS `admins` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `avatar_url` VARCHAR(500) DEFAULT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `last_login_at` DATETIME DEFAULT NULL,
  `last_login_ip` VARCHAR(50) DEFAULT NULL,
  `password_changed_at` DATETIME DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_admins_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 角色表
CREATE TABLE IF NOT EXISTS `roles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `code` VARCHAR(50) NOT NULL,
  `description` VARCHAR(200) DEFAULT NULL,
  `is_system` TINYINT(1) NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_roles_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 管理员-角色关联表
CREATE TABLE IF NOT EXISTS `admin_roles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` BIGINT UNSIGNED NOT NULL,
  `role_id` BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_admin_roles_admin_id` (`admin_id`),
  INDEX `idx_admin_roles_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 权限表
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `module` VARCHAR(50) NOT NULL,
  `action` VARCHAR(50) NOT NULL,
  `code` VARCHAR(100) NOT NULL,
  `description` VARCHAR(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_permissions_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 角色-权限关联表
CREATE TABLE IF NOT EXISTS `role_permissions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` BIGINT UNSIGNED NOT NULL,
  `permission_id` BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_role_permissions_role_id` (`role_id`),
  INDEX `idx_role_permissions_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 操作日志表
CREATE TABLE IF NOT EXISTS `operation_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` BIGINT UNSIGNED NOT NULL,
  `module` VARCHAR(50) DEFAULT NULL,
  `action` VARCHAR(50) DEFAULT NULL,
  `target_type` VARCHAR(50) DEFAULT NULL,
  `target_id` BIGINT UNSIGNED DEFAULT NULL,
  `target_name` VARCHAR(200) DEFAULT NULL,
  `detail` TEXT DEFAULT NULL,
  `ip` VARCHAR(50) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_operation_logs_admin_id` (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Banner 表
CREATE TABLE IF NOT EXISTS `banners` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL,
  `subtitle` VARCHAR(200) DEFAULT NULL,
  `image_url` VARCHAR(500) DEFAULT NULL,
  `link_type` VARCHAR(20) DEFAULT NULL,
  `link_target` VARCHAR(500) DEFAULT NULL,
  `sort_order` INT NOT NULL DEFAULT 0,
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `start_time` DATETIME DEFAULT NULL,
  `end_time` DATETIME DEFAULT NULL,
  `created_by` BIGINT UNSIGNED DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 公告表
CREATE TABLE IF NOT EXISTS `announcements` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(200) NOT NULL,
  `content` TEXT DEFAULT NULL,
  `type` VARCHAR(20) DEFAULT NULL,
  `push_method` VARCHAR(20) DEFAULT NULL,
  `target_scope` VARCHAR(20) NOT NULL DEFAULT 'all',
  `status` VARCHAR(20) NOT NULL DEFAULT 'draft',
  `publish_time` DATETIME DEFAULT NULL,
  `created_by` BIGINT UNSIGNED DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 推荐表
CREATE TABLE IF NOT EXISTS `recommendations` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `software_id` BIGINT UNSIGNED NOT NULL,
  `type` VARCHAR(20) DEFAULT NULL,
  `sort_order` INT NOT NULL DEFAULT 0,
  `start_time` DATETIME DEFAULT NULL,
  `end_time` DATETIME DEFAULT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'active',
  `created_by` BIGINT UNSIGNED DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_recommendations_software_id` (`software_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 系统配置表
CREATE TABLE IF NOT EXISTS `system_configs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `config_group` VARCHAR(50) DEFAULT NULL,
  `config_key` VARCHAR(100) NOT NULL,
  `config_value` TEXT DEFAULT NULL,
  `description` VARCHAR(200) DEFAULT NULL,
  `updated_by` BIGINT UNSIGNED DEFAULT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_system_configs_config_key` (`config_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 举报表
CREATE TABLE IF NOT EXISTS `reports` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `reporter_id` BIGINT UNSIGNED NOT NULL,
  `target_type` VARCHAR(20) DEFAULT NULL,
  `target_id` BIGINT UNSIGNED DEFAULT NULL,
  `reason` VARCHAR(50) DEFAULT NULL,
  `description` TEXT DEFAULT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending',
  `result` VARCHAR(20) DEFAULT NULL,
  `handler_id` BIGINT UNSIGNED DEFAULT NULL,
  `handler_remark` TEXT DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `resolved_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_reports_reporter_id` (`reporter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户反馈表
CREATE TABLE IF NOT EXISTS `feedbacks` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `type` VARCHAR(20) DEFAULT NULL,
  `title` VARCHAR(200) DEFAULT NULL,
  `content` TEXT DEFAULT NULL,
  `screenshots` TEXT DEFAULT NULL COMMENT 'JSON序列化的截图URL列表',
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending',
  `handler_id` BIGINT UNSIGNED DEFAULT NULL,
  `reply` TEXT DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `replied_at` DATETIME DEFAULT NULL,
  `closed_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_feedbacks_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================================
-- 初始数据
-- ============================================================

-- 默认管理员（密码: admin123，需在应用中使用 bcrypt 生成）
INSERT INTO `admins` (`name`, `email`, `password_hash`, `status`) VALUES
('超级管理员', 'admin@tbox.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'active');

-- 默认角色
INSERT INTO `roles` (`name`, `code`, `description`, `is_system`) VALUES
('超级管理员', 'super_admin', '拥有所有权限', 1),
('运营管理员', 'operator', '内容运营与审核管理', 0),
('审计员', 'auditor', '查看日志与审计信息', 0);

-- 绑定超级管理员角色
INSERT INTO `admin_roles` (`admin_id`, `role_id`) VALUES (1, 1);

-- 默认权限
INSERT INTO `permissions` (`module`, `action`, `code`, `description`) VALUES
('dashboard', 'view', 'dashboard:view', '查看仪表盘'),
('software', 'view', 'software:view', '查看软件列表'),
('software', 'edit', 'software:edit', '编辑软件信息'),
('software', 'audit', 'software:audit', '审核软件'),
('software', 'shelf', 'software:shelf', '上下架软件'),
('user', 'view', 'user:view', '查看用户列表'),
('user', 'edit', 'user:edit', '编辑用户信息'),
('user', 'disable', 'user:disable', '禁用/启用用户'),
('category', 'view', 'category:view', '查看分类'),
('category', 'edit', 'category:edit', '编辑分类'),
('tag', 'view', 'tag:view', '查看标签'),
('tag', 'edit', 'tag:edit', '编辑标签'),
('review', 'view', 'review:view', '查看评价'),
('review', 'moderate', 'review:moderate', '审核评价'),
('content', 'view', 'content:view', '查看内容管理'),
('content', 'edit', 'content:edit', '编辑内容'),
('report', 'view', 'report:view', '查看举报'),
('report', 'handle', 'report:handle', '处理举报'),
('feedback', 'view', 'feedback:view', '查看反馈'),
('feedback', 'reply', 'feedback:reply', '回复反馈'),
('system', 'view', 'system:view', '查看系统配置'),
('system', 'edit', 'system:edit', '编辑系统配置'),
('admin', 'view', 'admin:view', '查看管理员'),
('admin', 'edit', 'admin:edit', '编辑管理员'),
('role', 'view', 'role:view', '查看角色'),
('role', 'edit', 'role:edit', '编辑角色'),
('log', 'view', 'log:view', '查看操作日志');

-- 超级管理员角色绑定所有权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`)
SELECT 1, `id` FROM `permissions`;

-- 默认软件分类
INSERT INTO `categories` (`name`, `icon`, `sort_order`, `status`) VALUES
('开发工具', 'code', 1, 'active'),
('办公软件', 'briefcase', 2, 'active'),
('图形设计', 'image', 3, 'active'),
('系统工具', 'settings', 4, 'active'),
('网络工具', 'wifi', 5, 'active'),
('安全防护', 'shield', 6, 'active'),
('媒体播放', 'play-circle', 7, 'active'),
('教育学习', 'book', 8, 'active');
