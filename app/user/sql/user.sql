CREATE TABLE `user` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
     `username` varchar(32) NOT NULL COMMENT '用户名',
     `phone` varchar(15) NOT NULL COMMENT '手机号',
     `password_hash` varchar(128) NOT NULL COMMENT '密码哈希',
     `nickname` varchar(32) DEFAULT NULL COMMENT '用户昵称',
     `avatar_url` varchar(256) DEFAULT NULL COMMENT '头像URL',
     `user_status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-正常, 2-待验证',
     `version` int unsigned NOT NULL DEFAULT '0' COMMENT '版本号',
     `sex` tinyint DEFAULT '0' COMMENT '性别: 0-未知, 1-男, 2-女',
     `bio` varchar(256) DEFAULT NULL COMMENT '个人简介',
     `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
     `last_login_ip` varchar(45) DEFAULT NULL COMMENT '最后登录IP',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
     `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_username` (`username`),
     UNIQUE KEY `idx_phone` (`phone`),
     KEY `idx_create_time` (`create_time`),
     KEY `idx_user_status` (`user_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `user_follows` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '关注记录ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `follow_user_id` bigint unsigned NOT NULL COMMENT '被关注用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_follow` (`user_id`,`follow_user_id`),
    KEY `idx_follow_user_id` (`follow_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户关注表';

CREATE TABLE `user_notifications` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '通知ID',
      `user_id` bigint unsigned NOT NULL COMMENT '接收用户ID',
      `type` varchar(32) NOT NULL COMMENT '通知类型',
      `content` text NOT NULL COMMENT '通知内容',
      `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读: 0-未读, 1-已读',
      `related_id` bigint unsigned DEFAULT NULL COMMENT '相关ID',
      `related_type` varchar(32) DEFAULT NULL COMMENT '相关类型',
      `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
      `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
      `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
      PRIMARY KEY (`id`),
      KEY `idx_user_id` (`user_id`),
      KEY `idx_is_read` (`is_read`),
      KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户通知表';