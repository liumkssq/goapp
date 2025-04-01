CREATE TABLE `article_favorite` (
        `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
        `article_id` BIGINT UNSIGNED NOT NULL COMMENT '文章ID',
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
        PRIMARY KEY (`user_id`, `article_id`),
        KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章收藏表';
