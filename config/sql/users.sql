CREATE DATABASE user;

-- 用户表
CREATE TABLE `user` (
                        `user_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
                        `username` VARCHAR(30) NOT NULL COMMENT '用户名最多 10 个中文字符或等长英文字符',
                        `password` VARCHAR(255) NOT NULL COMMENT '数字+字母组合，总长度上限 16',
                        `email` VARCHAR(50) NOT NULL COMMENT '邮箱',
                        `role` SMALLINT NOT NULL default 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `address` (
                        `address_id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '使用自增主键',
                        `user_id` BIGINT NOT NULL COMMENT '用户 id',
                        `province` VARCHAR(30) NOT NULL COMMENT '省',
                        `city` VARCHAR(30) NOT NULL COMMENT '市',
                        `detail` VARCHAR(255) NOT NULL COMMENT '详细地址'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户地址表';
