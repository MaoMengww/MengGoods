CREATE DATABASE stock;

CREATE TABLE `stock` (
    `sku_id` BIGINT NOT NULL  PRIMARY KEY COMMENT 'sku id',
    `stock` BIGINT NOT NULL DEFAULT 0 COMMENT '库存数量, 总数',
    `locked_stock` BIGINT NOT NULL DEFAULT 0 COMMENT '锁定库存数量, 锁定的总数',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存表';



CREATE TABLE `stock_journal` (
    `journal_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `sku_id` BIGINT NOT NULL COMMENT 'sku id',
    `order_id` BIGINT NOT NULL COMMENT '订单 id',
    `change_num` BIGINT NOT NULL COMMENT '库存变化量',
    `change_type` TINYINT NOT NULL COMMENT '库存变化类型, 1: 锁定库存, 2: 释放内存 3: 扣减库存 4: 恢复库存',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE INDEX `uk_id` (`order_id`, `sku_id`, `change_type`),
    INDEX `sku_id` (`sku_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存流水表';