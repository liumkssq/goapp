CREATE TABLE `report` (
          `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '举报ID',
          `product_id` BIGINT UNSIGNED NOT NULL COMMENT '被举报商品ID',
          `user_id` BIGINT UNSIGNED NOT NULL COMMENT '举报用户ID',
          `reason` VARCHAR(255) NOT NULL COMMENT '举报原因',
          `description` TEXT DEFAULT NULL COMMENT '举报详情',
          `images` JSON DEFAULT NULL COMMENT '举报证据图片，存储图片 URL 列表',
          `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '举报时间',
          PRIMARY KEY (`id`),
          KEY `idx_product_id` (`product_id`),
          KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品举报表';
