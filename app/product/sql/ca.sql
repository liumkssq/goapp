CREATE TABLE `category` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类ID',
    `name` VARCHAR(100) NOT NULL COMMENT '分类名称',
    `icon` VARCHAR(255) DEFAULT NULL COMMENT '分类图标URL',
    `parent_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '父分类ID，0 表示顶级分类',
    `level` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '分类层级',
    `sort` INT UNSIGNED DEFAULT 0 COMMENT '排序值，越大越靠前',
    `description` VARCHAR(255) DEFAULT NULL COMMENT '分类描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品分类表';
