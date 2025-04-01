CREATE TABLE `comment` (
       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论ID',
       `product_id` BIGINT UNSIGNED NOT NULL COMMENT '所属商品ID',
       `user_id` BIGINT UNSIGNED NOT NULL COMMENT '评论用户ID',
       `user_name` VARCHAR(100) NOT NULL COMMENT '用户名称',
       `user_avatar` VARCHAR(255) DEFAULT NULL COMMENT '用户头像URL',
       `content` TEXT NOT NULL COMMENT '评论内容',
       `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
       PRIMARY KEY (`id`),
       KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品评论表';
