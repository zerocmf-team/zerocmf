-- 规格明细表 - product_attr_key
CREATE TABLE IF NOT EXISTS product_attr_key (
    attr_key_id INT AUTO_INCREMENT, -- 属性ID，主键，自增
    attribute_key VARCHAR(10) NOT NULL default '', -- 属性名称，如"颜色"、"内存"等
    `created_at` bigint(20) NOT NULL DEFAULT 0,
    `updated_at` bigint(20) NOT NULL DEFAULT 0,
    `deleted_at` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (attr_key_id),
    UNIQUE idx_attribute_key (attribute_key)
);