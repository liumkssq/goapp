# CREATE TABLE `category` (
#     `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类ID',
#     `name` VARCHAR(100) NOT NULL COMMENT '分类名称',
#     `icon` VARCHAR(255) DEFAULT NULL COMMENT '分类图标URL',
#     `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父分类ID，NULL 表示顶级分类',
#     `level` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '分类层级',
#     `sort` INT UNSIGNED DEFAULT 0 COMMENT '排序值，越大越靠前',
#     PRIMARY KEY (`id`),
#     KEY `idx_parent_id` (`parent_id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章分类表';
