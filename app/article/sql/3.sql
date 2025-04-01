CREATE TABLE `article_comment` (
       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论ID',
       `article_id` BIGINT UNSIGNED NOT NULL COMMENT '所属文章ID',
       `user_id` BIGINT UNSIGNED NOT NULL COMMENT '评论用户ID',
       `content` TEXT NOT NULL COMMENT '评论内容',
       `like_count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
       `reply_count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复数',
       `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父评论ID，NULL 表示一级评论',
       `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
       PRIMARY KEY (`id`),
       KEY `idx_article_id` (`article_id`),
       KEY `idx_user_id` (`user_id`),
       KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章评论表';
