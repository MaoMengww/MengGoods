CREATE DATABASE payment;

CREATE TABLE `payment_order` (
    `payment_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `payment_no` VARCHAR(255) NOT NULL COMMENT '支付订单号',
    `order_id` BIGINT NOT NULL COMMENT '订单 id',
    `user_id` BIGINT NOT NULL COMMENT '用户 id',
    `payment_amount` BIGINT NOT NULL COMMENT '支付金额(单位:分)',
    `payment_method` INT NOT NULL COMMENT '支付方式, 0: 微信支付, 1: 支付宝支付',
    `payment_status` INT NOT NULL DEFAULT 0 COMMENT '支付状态, 0: 未支付, 1: 处理中, 2: 支付成功, 3: 支付失败',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_order_id` (`order_id`),
    INDEX `idx_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付订单表';

CREATE TABLE `payment_refund` (
    `refund_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `refund_no` VARCHAR(255) NOT NULL COMMENT '退款订单号',
    `order_item_id` BIGINT NOT NULL COMMENT '订单商品 id',
    `payment_no` VARCHAR(255) NOT NULL COMMENT '支付订单号',
    `seller_id` BIGINT NOT NULL COMMENT '卖家 id',
    `user_id` BIGINT NOT NULL COMMENT '用户 id',
    `refund_amount` BIGINT NOT NULL COMMENT '退款金额(单位:分)',
    `refund_reason` VARCHAR(255) NOT NULL COMMENT '退款原因',
    `refund_status` INT NOT NULL DEFAULT 0 COMMENT '退款状态, 0: 未退款, 1: 处理中, 2: 退款成功, 3: 退款失败',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_payment_id` (`payment_no`),
    INDEX `idx_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='退款订单表';

CREATE TABLE `payment_transaction` (
    `transaction_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `transaction_no` VARCHAR(255) NOT NULL COMMENT '支付交易号',
    `order_id` BIGINT NOT NULL COMMENT '支付订单 id',
    `seller_id` BIGINT NOT NULL COMMENT '卖家 id',
    `user_id` BIGINT NOT NULL COMMENT '用户 id',
    `transaction_amount` BIGINT NOT NULL COMMENT '支付交易金额(单位:分)',
    `transaction_type` INT NOT NULL DEFAULT 0 COMMENT '支付交易类型, 0: 支付, 1: 退款',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX `idx_order_id` (`order_id`),
    INDEX `idx_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付流水表';
