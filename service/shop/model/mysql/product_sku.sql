-- 商品规格表 - goods_sku
CREATE TABLE IF NOT EXISTS product_sku
(
    sku_id         BIGINT(20) AUTO_INCREMENT COMMENT 'SKU的唯一标识符，主键，自增',
    product_id     BIGINT(20)    NOT NULL DEFAULT 0 COMMENT '所属SPU的标识符，外键关联SPU表',
    sku_code       VARCHAR(32) COMMENT '规格编码',
    sku_barcode    VARCHAR(32) COMMENT '规格条码',
    attrs_val      VARCHAR(255)  NOT NULL DEFAULT '' COMMENT '属性值组合，例如"颜色:红色,存储容量:256G,网络类型:全网通"',
    retail_price   DECIMAL(8, 2) NOT NULL DEFAULT 0 COMMENT '零售价',
    stock          INT           NOT NULL DEFAULT 0 COMMENT '商品库存数量',
    original_price DECIMAL(8, 2) COMMENT '标准价',
    cost_price     DECIMAL(8, 2) COMMENT '成本价',
    weight         FLOAT(0) COMMENT '重量（g）',
    status         TINYINT(3)    NOT NULL DEFAULT 1 COMMENT '状态（0 => 停用;1 => 启用）',
    `created_at`   bigint(20)    NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at`   bigint(20)    NOT NULL DEFAULT 0 COMMENT '更新时间',
    `deleted_at`   bigint(20)    NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (sku_id)
);