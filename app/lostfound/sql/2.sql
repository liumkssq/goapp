CREATE TABLE `lost_found_comment` (
          `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论ID',
          `lost_found_item_id` BIGINT UNSIGNED NOT NULL COMMENT '所属失物招领记录ID',
          `user_id` BIGINT UNSIGNED NOT NULL COMMENT '评论用户ID',
          `user_name` VARCHAR(100) NOT NULL COMMENT '评论者名称',
          `user_avatar` VARCHAR(255) DEFAULT NULL COMMENT '评论者头像',
          `content` TEXT NOT NULL COMMENT '评论内容',
          `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
          PRIMARY KEY (`id`),
          KEY `idx_lost_found_item_id` (`lost_found_item_id`),
          KEY `idx_user_id` (`user_id`),
          CONSTRAINT `fk_comment_item` FOREIGN KEY (`lost_found_item_id`) REFERENCES `lost_found_item` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='失物招领评论表';
