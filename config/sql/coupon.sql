CREATE TABLE `coupon_batch`(
    `batch_id` BIGINT NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT '优惠券批次 id',
    `batch_name` VARCHAR(255) NOT NULL COMMENT '优惠券批次名称',
    `remark` VARCHAR(255) NOT NULL COMMENT '优惠券模板描述',
    `type` TINYINT NOT NULL DEFAULT 0 COMMENT '优惠券类型, 0: 满减券, 1: 折扣券',
    `threshold` BIGINT NOT NULL DEFAULT 0 COMMENT '优惠券使用门槛(分)',
    `discount_amount` BIGINT NOT NULL DEFAULT 0 COMMENT '优惠券折扣金额(分)',
    `discount_rate` INT NOT NULL DEFAULT 0 COMMENT '优惠券折扣比例(%)',   
    `total_num` INT NOT NULL DEFAULT 0 COMMENT '优惠券总量',
    `used_num` INT NOT NULL DEFAULT 0 COMMENT '优惠券已发放数量',
    `start_time` DATETIME NOT NULL COMMENT '优惠券开始时间',
    `end_time` DATETIME NOT NULL COMMENT '优惠券结束时间',
    `duration` INT NOT NULL DEFAULT 0 COMMENT '优惠券持续时间(天)',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券批次表';

CREATE TABLE `coupon`(
    `coupon_id` BIGINT NOT NULL  PRIMARY KEY COMMENT '优惠券 id, 使用雪花算法生成',
    `batch_id` BIGINT NOT NULL COMMENT '优惠券批次 id',
    `used_order_id` BIGINT NULL DEFAULT NULL COMMENT '使用订单 id',
    `user_id` BIGINT NOT NULL COMMENT '用户 id',
    `type` TINYINT NOT NULL DEFAULT 0 COMMENT '优惠券类型, 0: 满减券, 1: 折扣券',
    `threshold` BIGINT NOT NULL DEFAULT 0 COMMENT '优惠券使用门槛(分)',
    `discount_amount` BIGINT NOT NULL DEFAULT 0 COMMENT '优惠券折扣金额(分)',
    `discount_rate` DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '优惠券折扣比例',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '优惠券状态, 0: 未使用, 1: 已冻结, 2: 已使用, 3: 已过期',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '领取时间',
    `expired_at` DATETIME NULL DEFAULT NULL COMMENT '过期时间',
    `used_at` DATETIME NULL DEFAULT NULL COMMENT '使用时间',
    INDEX `idx_user_status` (`user_id`, `status`),
    INDEX `idx_batch_id` (`batch_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优惠券表';

