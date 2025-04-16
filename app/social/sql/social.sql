CREATE TABLE `friends` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '好友关系ID',
   `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
   `friend_id` bigint unsigned NOT NULL COMMENT '好友用户ID',
   `remark` varchar(255) DEFAULT NULL COMMENT '好友备注',
   `add_source` tinyint DEFAULT NULL COMMENT '添加来源: 0-搜索, 1-群聊, 2-附近的人',
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
   `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
   PRIMARY KEY (`id`),
   KEY `idx_user_id` (`user_id`),
   KEY `idx_friend_id` (`friend_id`),
   UNIQUE KEY `idx_user_friend` (`user_id`,`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友关系表';

CREATE TABLE `friend_requests` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '好友请求ID',
   `user_id` bigint unsigned NOT NULL COMMENT '发起用户ID',
   `req_user_id` bigint unsigned NOT NULL COMMENT '接收请求的用户ID',
   `req_msg` varchar(255) DEFAULT NULL COMMENT '请求消息',
   `req_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '请求时间',
   `handle_result` tinyint DEFAULT NULL COMMENT '处理结果: 0-拒绝, 1-接受, 2-忽略',
   `handle_msg` varchar(255) DEFAULT NULL COMMENT '处理消息',
   `handle_time` timestamp NULL DEFAULT NULL COMMENT '处理时间',
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
   `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
   PRIMARY KEY (`id`),
   KEY `idx_user_id` (`user_id`),
   KEY `idx_req_user_id` (`req_user_id`),
   KEY `idx_req_time` (`req_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友请求表';

CREATE TABLE `groups` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '群组ID',
  `group_code` varchar(24) NOT NULL COMMENT '群号',
  `name` varchar(255) NOT NULL COMMENT '群名称',
  `icon` varchar(255) DEFAULT NULL COMMENT '群图标',
  `status` tinyint DEFAULT '1' COMMENT '状态: 0-禁用, 1-正常',
  `creator_id` bigint unsigned NOT NULL COMMENT '创建者用户ID',
  `group_type` tinyint NOT NULL DEFAULT '0' COMMENT '群类型: 0-普通群, 1-临时群, 2-系统群',
  `is_verify` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否需要验证: 0-不需要, 1-需要',
  `notification` varchar(255) DEFAULT NULL COMMENT '群公告',
  `notification_user_id` bigint unsigned DEFAULT NULL COMMENT '公告发布人ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_group_code` (`group_code`),
  KEY `idx_creator_id` (`creator_id`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群组表';

CREATE TABLE `group_members` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '群成员ID',
 `group_id` bigint unsigned NOT NULL COMMENT '群组ID',
 `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
 `role_level` tinyint NOT NULL DEFAULT '0' COMMENT '角色等级: 0-普通成员, 1-管理员, 2-群主',
 `join_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
 `join_source` tinyint DEFAULT NULL COMMENT '加入来源: 0-搜索, 1-邀请, 2-扫码',
 `inviter_id` bigint unsigned DEFAULT NULL COMMENT '邀请人ID',
 `operator_id` bigint unsigned DEFAULT NULL COMMENT '操作人ID',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
 `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
 `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_group_user` (`group_id`,`user_id`),
 KEY `idx_user_id` (`user_id`),
 KEY `idx_join_time` (`join_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群成员表';

CREATE TABLE `group_requests` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '入群请求ID',
  `req_user_id` bigint unsigned NOT NULL COMMENT '请求用户ID',
  `group_id` bigint unsigned NOT NULL COMMENT '群组ID',
  `req_msg` varchar(255) DEFAULT NULL COMMENT '请求消息',
  `req_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '请求时间',
  `join_source` tinyint DEFAULT NULL COMMENT '来源: 0-搜索, 1-邀请, 2-扫码',
  `inviter_id` bigint unsigned DEFAULT NULL COMMENT '邀请人ID',
  `handle_user_id` bigint unsigned DEFAULT NULL COMMENT '处理人ID',
  `handle_time` timestamp NULL DEFAULT NULL COMMENT '处理时间',
  `handle_result` tinyint DEFAULT NULL COMMENT '处理结果: 0-拒绝, 1-接受, 2-忽略',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态: 0-正常, 1-已删除',
  PRIMARY KEY (`id`),
  KEY `idx_req_user_id` (`req_user_id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_req_time` (`req_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='入群请求表';

