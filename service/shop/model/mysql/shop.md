# 商城字段说明

本文档包含商城的字段说明及其备注。

## 商品列表 - product

| 字段名                   | 类型            | 备注                  |
|-----------------------|---------------|---------------------|
| product_id            | BIGINT        | 商品ID，主键，自增长整数       |
| product_name          | VARCHAR(100)  | 商品名称，不可为空           |
| userId                | bigint(20)    | 创建人                 |
| product_barcode       | VARCHAR(32)   | 商品条码，不可为空           |
| product_category      | BIGINT(20)    | 商品分类                |
| main_video            | VARCHAR(255)  | 主图视频，存储视频的URL或文件路径  |
| explanation_video     | VARCHAR(255)  | 讲解视频，存储视频的URL或文件路径  |
| stock_unit            | INT(11)       | 库存单位                |
| share_description     | VARCHAR(100)  | 分享描述，用于在分享时显示的商品描述  |
| product_selling_point | VARCHAR(200)  | 商品卖点，突出商品的特点        |
| strikethrough_price   | DECIMAL(8, 2) | 划线价，商品的原价或划线价       |
| hide_remaining_stock  | tinyint(3)    | 商品详情不显示剩余件数         |
| delivery_method       | tinyint(3)    | 配送方式                |
| product_content       | TEXT          | 图文信息，存储商品的图文信息      |
| status                | tinyint(3)    | 状态（0 => 停用;1 => 启用） |
| created_at            | bigint(20)    | 创建时间                |
| updated_at            | bigint(20)    | 修改时间                |
| deleted_at            | bigint(20)    | 删除时间                |

## 商品规格表 - product_sku

| 字段名            | 数据类型    | 描述                 |
|----------------|---------|--------------------|
| sku_id         | int     | SKU的唯一标识符，主键，自增    |
| spu_id         | int     | 所属SPU的标识符，外键关联SPU表 |
| sku_code       | varchar | 规格编码               |
| sku_barcode    | varchar | 规格条码               |
| retail_price   | decimal | 零售价                |
| standard_price | decimal | 标准价                |
| cost_price     | decimal | 成本价                |
| weight         | decimal | 重量（kg）             |
| stock          | int     | 商品库存数量             |

## 商品规格表 - product_sku_attr_relation

| 字段名          | 数据类型    | 描述           |
|--------------|---------|--------------|
| attribute_id | INT(11) | 属性ID，主键，自增   |
| sku_id       | INT(11) | SKU的唯一标识     |
| attr_key_id  | INT(11) | 属性key的唯一ID   |
| attr_val_id  | INT(11) | 属性value的唯一ID |

## 规格明细表 - product_attr_key

| 字段名           | 类型      | 备注               |
|---------------|---------|------------------|
| attr_key_id   | INT(11) | 属性ID，主键，自增       |
| attribute_key | varchar | 属性名称，如"颜色"、"内存"等 |

## 规格明细表 - product_attr_val

| 字段名             | 类型      | 备注                 |
|-----------------|---------|--------------------|
| attr_val_id	    | INT(11) | 属性选项ID，主键，自增       |
| attribute_value | VARCHAR | 属性选项名称，如"黑色"、"白色"等 |

## 商品分类 - product_category