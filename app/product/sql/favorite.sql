CREATE TABLE `favorite` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键，收藏ID',
        `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
        `product_id` BIGINT UNSIGNED NOT NULL COMMENT '商品ID',
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uniq_user_product` (`user_id`, `product_id`),  -- 保证同一个用户不能重复收藏同个商品
        KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏商品关系表';
