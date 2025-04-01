CREATE TABLE `lost_found_like` (
       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
       `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
       `lost_found_item_id` BIGINT UNSIGNED NOT NULL COMMENT '失物招领记录ID',
       `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间',
       PRIMARY KEY (`id`),
       UNIQUE KEY `uniq_user_item` (`user_id`, `lost_found_item_id`), -- 确保用户对同一记录只能点赞一次
       KEY `idx_lost_found_item_id` (`lost_found_item_id`),
       CONSTRAINT `fk_like_item` FOREIGN KEY (`lost_found_item_id`) REFERENCES `lost_found_item` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='失物招领点赞记录表';
