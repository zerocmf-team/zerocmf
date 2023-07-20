CREATE TABLE IF NOT EXISTS goods_category
(
    `id`         bigint(20) AUTO_INCREMENT,
    `parent_id`  bigint(20) COMMENT '父级id',
    `path`       varchar(255) COMMENT '层级path',
    `name`       varchar(20) COMMENT '分类名称',
    `icon`       varchar(255) COMMENT '分类图标',
    `desc`       varchar(255) COMMENT '分类描述',
    `list_order` double     DEFAULT '10000' COMMENT '排序',
    `status`     tinyint(3) DEFAULT '1' COMMENT '0 => 停用;1 => 启用',
    `created_at` bigint(20) NULL,
    `updated_at` bigint(20) NULL,
    `deleted_at` bigint(20) NULL,
    PRIMARY KEY (id),
    INDEX idx_parent_id (parent_id)
) COMMENT ='商品分类表';