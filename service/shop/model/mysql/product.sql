-- 商品列表表 - product
CREATE TABLE IF NOT EXISTS product
(
    product_id            BIGINT AUTO_INCREMENT COMMENT '商品ID，主键，自增长整数',
    product_name          VARCHAR(100)  NOT NULL DEFAULT '默认产品' COMMENT '商品名称，不可为空',
    userId                BIGINT(20)    NOT NULL DEFAULT 0 COMMENT '创建人',
    attributes            json          NOT NULL COMMENT '规格属性选项',
    product_barcode       VARCHAR(32)   NOT NULL DEFAULT '' COMMENT '商品条码，不可为空',
    product_category      BIGINT(20)    NOT NULL DEFAULT 0 COMMENT '商品分类',
    product_thumbnail     json COMMENT '商品缩略图',
    main_video            VARCHAR(255) COMMENT '主图视频，存储视频的URL或文件路径',
    explanation_video     VARCHAR(255) COMMENT '讲解视频，存储视频的URL或文件路径',
    price                 DECIMAL(8, 2) NOT NULL DEFAULT 0 COMMENT '商品售价',
    price_negotiable      tinyint(3) COMMENT '价格面议',
    stock_unit            varchar(10) COMMENT '库存单位',
    stock                 INT(11) COMMENT '库存',
    share_description     VARCHAR(100) COMMENT '分享描述，用于在分享时显示的商品描述',
    product_selling_point VARCHAR(200) COMMENT '商品卖点，突出商品的特点',
    original_price        DECIMAL(8, 2) COMMENT '划线价，商品的原价或划线价',
    cost_price            DECIMAL(8, 2) COMMENT '成本价，商品的成本价',
    hide_remaining_stock  tinyint(3) COMMENT '商品详情不显示剩余件数',
    delivery_method       tinyint(3) COMMENT '配送方式',
    product_content       TEXT COMMENT '图文信息，存储商品的图文信息',
    status                TINYINT(3)    NOT NULL DEFAULT 1 COMMENT '状态（0 => 停用;1 => 启用）',
    `created_at`          bigint(20)    NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at`          bigint(20)    NOT NULL DEFAULT 0 COMMENT '更新时间',
    `deleted_at`          bigint(20)    NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (product_id)
);