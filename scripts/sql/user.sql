CREATE TABLE users (
       id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
       username VARCHAR(32) NOT NULL COMMENT '用户名',
       phone VARCHAR(20) NOT NULL COMMENT '手机号',
       password_hash VARCHAR(128) NOT NULL COMMENT '密码哈希',
       nickname VARCHAR(32) DEFAULT NULL COMMENT '用户昵称',
       avatar_url VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
    version BIGINT UNSIGNED DEFAULT 0 COMMENT '版本号',
    sex TINYINT UNSIGNED DEFAULT 1 COMMENT '性别: 0-女,1-男',
    bio VARCHAR(255) DEFAULT NULL COMMENT '简介',
       user_status TINYINT UNSIGNED DEFAULT 1 COMMENT '状态: 0-禁用,1-正常,2-待验证',
       last_login_at TIMESTAMP NULL DEFAULT NULL COMMENT '最后登录时间',
       last_login_ip VARCHAR(64) DEFAULT NULL COMMENT '最后登录IP',
       `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
       `del_state` tinyint NOT NULL DEFAULT '0',
       UNIQUE KEY uk_username (username),
       UNIQUE KEY uk_phone (phone),
       KEY idx_status (user_status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


