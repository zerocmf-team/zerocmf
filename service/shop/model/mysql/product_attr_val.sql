-- 规格明细表 - product_attr_val
CREATE TABLE IF NOT EXISTS product_attr_val (
    attr_val_id INT AUTO_INCREMENT, -- 属性选项ID，主键，自增
    attribute_value VARCHAR(10) NOT NULL default '', -- 属性选项名称，如"黑色"、"白色"等
    `created_at` bigint(20) NOT NULL DEFAULT 0,
    `updated_at` bigint(20) NOT NULL DEFAULT 0,
    `deleted_at` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (attr_val_id),
    UNIQUE idx_attribute_value (attribute_value)
);