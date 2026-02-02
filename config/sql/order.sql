CREATE DATABASE orders;

CREATE TABLE `orders` (
    `order_id` BIGINT NOT NULL  PRIMARY KEY COMMENT '使用订单 id, ',
    `user_id` BIGINT NOT NULL COMMENT '用户 id',
    `order_status` TINYINT NOT NULL DEFAULT 0 COMMENT '订单状态, 0: 待支付, 1: 已支付, 2: 已发货, 3: 已确认, 4: 已取消',

    `total_price` BIGINT NOT NULL COMMENT '订单总金额(单位:分)',
    `pay_price` BIGINT NOT NULL COMMENT '实付金额(单位:分)',
    `receiver_name` VARCHAR(30) NOT NULL COMMENT '收货人姓名',
    `receiver_email` VARCHAR(50) NOT NULL COMMENT '收货人邮箱',
    `receiver_province` VARCHAR(30) NOT NULL COMMENT '省',
    `receiver_city` VARCHAR(30) NOT NULL COMMENT '市',
    `receiver_detail` VARCHAR(255) NOT NULL COMMENT '详细地址',

    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单更新时间',
    `expire_time` DATETIME NOT NULL COMMENT '订单过期时间',
    `pay_time` DATETIME NULL DEFAULT NULL COMMENT '订单支付时间',
    `cancel_time` DATETIME NULL DEFAULT NULL COMMENT '订单取消时间',
    `cancel_reason` VARCHAR(255) NULL DEFAULT NULL COMMENT '订单取消原因',
    INDEX `idx_user_status` (`user_id`, `order_status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE `order_item` (
    `order_item_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `order_id` BIGINT NOT NULL COMMENT '订单 id',
    `seller_id` BIGINT NOT NULL COMMENT '卖家 id',
    `product_id` BIGINT NOT NULL COMMENT '商品 id',
    `product_name` VARCHAR(255) NOT NULL COMMENT '商品名称',
    `product_price` BIGINT NOT NULL COMMENT '商品价格(单位:分)',
    `product_img` VARCHAR(255) NOT NULL COMMENT '商品图片',
    `product_num` INT NOT NULL COMMENT '商品数量',
    `product_total_price` BIGINT NOT NULL COMMENT '商品总价(单位:分)',
    `product_properties` VARCHAR(255) NULL DEFAULT NULL COMMENT '商品属性',
    INDEX `idx_order_id` (`order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单商品表';

-- 本表用于解决双写不一致问题(不会rocketmq的坏处)
CREATE TABLE `mq_msg` (
    `msg_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
    `order_id` BIGINT NOT NULL COMMENT '订单 id',
    `coupon_id` BIGINT NULL DEFAULT NULL COMMENT '优惠券 id',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息创建时间',
    `msg_status` TINYINT NOT NULL DEFAULT 0 COMMENT '消息状态, 0: 未处理, 1: 已处理'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单消息表';
