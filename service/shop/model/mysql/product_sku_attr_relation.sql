-- 商品规格表 - goods_sku_attr_relation
CREATE TABLE IF NOT EXISTS product_sku_attr_relation
(
    attribute_id BIGINT(20) AUTO_INCREMENT,     -- 属性ID，主键，自增
    sku_id       BIGINT(20) NOT NULL DEFAULT 0, -- SKU的唯一标识，关联goods_sku表
    attr_key_id  INT        NOT NULL DEFAULT 0, -- 属性key的唯一ID，关联goods_attr_key表
    attr_val_id  INT        NOT NULL DEFAULT 0, -- 属性value的唯一ID，关联goods_attr_val表
    `created_at` bigint(20) NOT NULL DEFAULT 0,
    `updated_at` bigint(20) NOT NULL DEFAULT 0,
    `deleted_at` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (attribute_id)
);