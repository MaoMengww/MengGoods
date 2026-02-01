CREATE DATABASE product;

CREATE TABLE `spu` (
    `spu_id` BIGINT NOT NULL  PRIMARY KEY COMMENT '主键, 使用雪花算法',
    `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
    `price` BIGINT NOT NULL COMMENT 'sku的最低价格',
    `description` VARCHAR(255) NOT NULL COMMENT '商品描述',
    `main_image_url` VARCHAR(255) NOT NULL COMMENT '商品主图url',
    `slider_image_urls` VARCHAR(255) NOT NULL COMMENT '商品轮播图urls',
    `creator` BIGINT NOT NULL COMMENT '创建者 id',
    `category` INT NOT NULL  COMMENT '商品类别',
    `status` INT NOT NULL DEFAULT 0 COMMENT '商品状态',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_category` (`category`),
    INDEX `idx_creator` (`creator`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

CREATE TABLE `sku` (
    `sku_id` BIGINT NOT NULL  PRIMARY KEY COMMENT '主键, 使用雪花算法',
    `name` VARCHAR(255) NOT NULL COMMENT '商品sku名称',
    `price` BIGINT NOT NULL COMMENT '商品sku价格',
    `description` VARCHAR(255) NOT NULL COMMENT '商品sku描述',
    `image_url` VARCHAR(255) NOT NULL COMMENT '商品sku图片url',
    `spu_id` BIGINT NOT NULL COMMENT '商品id',
    `properties` VARCHAR(255) NOT NULL COMMENT '商品sku属性',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_spu_id` (`spu_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品sku表';

CREATE TABLE `category` (
    `category_id` BIGINT NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `name` VARCHAR(255) NOT NULL COMMENT '商品类别名称',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';
