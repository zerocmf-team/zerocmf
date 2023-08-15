# 模板管理

## 应用表(theme)

| 字段        | 类型           | 备注    |
|-----------|--------------|-------|
| id        | int(11)      | 自增id  |
| name      | varchar(20)  | 应用名称  |
| version   | varchar(10)  | 应用版本  |
| desc      | varchar(255) | 应用描述  |
| thumbnail | varchar(255) | 应用缩略图 |
| userId    | bigint((20)  | 创建人   |
| createAt  | bigint(20)   | 创建时间  |
| updateAt  | bigint(20)   | 更新时间  |
| listOrder | double(10)   | 排序    |

## 页面表(page)

| 字段       | 类型           | 备注      |
|----------|--------------|---------|
| id       | int(11)      | 自增id    |
| themeId  | int(11)      | 所属主题    |
| isPublic | tinyint(3)   | 是否是公共模块 |
| name     | varchar(20)  | 页面名称    |
| desc     | varchar(255) | 页面描述    |
| createAt | bigint(20)   | 创建时间    |
| updateAt | bigint(20)   | 更新时间    |
| order    | double(10)   | 排序      |
