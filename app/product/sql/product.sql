CREATE TABLE `product` (
   `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键，商品ID',
   `title` VARCHAR(255) NOT NULL COMMENT '商品标题',
   `description` TEXT NOT NULL COMMENT '商品描述',
   `price` DECIMAL(10,2) NOT NULL COMMENT '售价，单位元，可用分来存储（如 price*100 作为整数）',
   `original_price` DECIMAL(10,2) DEFAULT NULL COMMENT '原价，允许为空',
   `category_id` BIGINT UNSIGNED NOT NULL COMMENT '商品分类ID',
   `images` JSON NOT NULL COMMENT '商品图片，建议用 JSON 数组存储多个图片 URL',
   `condition` ENUM('全新','9成新','8成新','7成新','其他') DEFAULT '其他' COMMENT '商品成色',
   `contact_info` VARCHAR(255) DEFAULT NULL COMMENT '联系信息，如QQ、微信或其他',
   `contact_way` VARCHAR(100) DEFAULT NULL COMMENT '联系方式，如电话号码',
   `location` VARCHAR(255) DEFAULT NULL COMMENT '大致位置，如校区、城市',
   `location_detail` JSON DEFAULT NULL COMMENT '详细位置，例如存储坐标、具体地址',
   `tags` JSON DEFAULT NULL COMMENT '标签，存储 JSON 数组',
   `status` ENUM('active','sold','pending','deleted') NOT NULL DEFAULT 'active' COMMENT '状态：在售、已售出、待审核、已下架/删除',
   `seller_id` BIGINT UNSIGNED NOT NULL COMMENT '卖家ID',
   `seller_name` VARCHAR(100) NOT NULL COMMENT '卖家名称',
   `seller_avatar` VARCHAR(255) DEFAULT NULL COMMENT '卖家头像URL',
   `view_count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览数',
   `like_count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
   `comment_count` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数',
   `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`),
   KEY `idx_category_id` (`category_id`),
   KEY `idx_seller_id` (`seller_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';
