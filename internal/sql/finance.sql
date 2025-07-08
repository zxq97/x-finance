CREATE TABLE t_main_order (
    `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `main_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '主订单id',
    `target_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '订单总金额',
    `real_pay_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '用户实付金额',
    `coupon_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '优惠券金额',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_mainid` (`main_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '总订单表';

CREATE TABLE t_sub_order (
    `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `main_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '主订单id',
    `order_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '子订单id',
    `order_type` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '订单类型',
    `order_status` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '订单状态 0:未支付 1:支付成功 2:已取消',
    `balance` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '用户余额',
    `discount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '优惠金额',
    `settle` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '分账金额',
    `self_settle` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '自分账金额',
    `other_settle` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '他分账金额',
    `profit` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '计佣金额',
    `self_profit` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '自计佣金额',
    `other_profit` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '他计佣金额',
    `divide_status` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '分账状态 0:未分账 1:已分账 2:分账失败',
    `notify_status` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '通知状态 0:未通知 1:已通知 2:通知失败',
    `paid_time` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '支付时间',
    `canceled_time` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '取消时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_mainid_orderid` (`main_id`, `order_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '支付子单表';

CREATE TABLE t_order_snapshot (
    `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `main_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '主订单id',
    `order_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '子订单id',
    `pay_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '支付id',
    `trade_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '交易id',
    `target` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '结算价',
    `real_pay` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '用户实付',
    `coupon` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '优惠券金额',
    `coupon_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '优惠券id',
    `pay_channel_id` INT unsigned NOT NULL DEFAULT 0 COMMENT '支付渠道id',
    `rate` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '抽佣率',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_mainid_orderid` (`main_id`, `order_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '支付子单快照表';

CREATE TABLE t_main_refund (
    `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `main_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '主订单id',
    `refund_no` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '退款单号',
    `deduct_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '扣款金额',
    `refund_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '退款金额',
    `refund_succeeded_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '成功退款金额',
    `refund_failed_amount` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '退款失败金额',
    `refund_reason` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '退款原因',
    `refund_type` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '退款类型',
    `refund_status` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '退款状态 80:退款中 81:退款成功 82:退款失败',
    `notify_status` TINYINT unsigned NOT NULL DEFAULT 0 COMMENT '通知状态 0:未通知 1:已通知 2:通知失败',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_mainid_refundno` (`main_id`, `refund_no`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '退款表';

CREATE TABLE t_sub_refund (
    `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `main_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '主订单id',
    `refund_no` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '退款单号',
    `order_id` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '子订单id',
    `refund` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '退款金额',
    `refund_real` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '实退金额',
    `refund_coupon` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '退券金额',
    `self_refund` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '自退金额',
    `other_refund` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '他退金额',
    `profit_refund` BIGINT unsigned NOT NULL DEFAULT 0 COMMENT '计佣金额退款'
    `refund_status` TINYINT NOT NULL DEFAULT 0 COMMENT '退款状态 80:退款中 81:退款成功 82:退款失败',
    `complete_time` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '退款完成时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '退款流水表';
