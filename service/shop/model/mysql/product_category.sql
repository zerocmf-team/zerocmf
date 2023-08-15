CREATE TABLE IF NOT EXISTS product_category
(
    `product_category_id` bigint(20) AUTO_INCREMENT,
    `parent_id`           bigint(20) COMMENT '父级id',
    `path`                varchar(255) COMMENT '层级path' NOT NULL DEFAULT '',
    `name`                varchar(20) COMMENT '分类名称'  NOT NULL DEFAULT '',
    `icon`                varchar(255) COMMENT '分类图标',
    `desc`                varchar(255) COMMENT '分类描述',
    `list_order`          double                          NOT NULL DEFAULT '10000' COMMENT '排序',
    `status`              tinyint(3)                      NOT NULL DEFAULT '1' COMMENT '0 => 停用;1 => 启用',
    `created_at`          bigint(20)                      NOT NULL DEFAULT 0,
    `updated_at`          bigint(20)                      NOT NULL DEFAULT 0,
    `deleted_at`          bigint(20)                      NOT NULL DEFAULT 0,
    PRIMARY KEY (product_category_id),
    INDEX idx_parent_id (parent_id)
) COMMENT ='商品分类表';