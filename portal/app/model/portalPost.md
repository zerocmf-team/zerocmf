# 门户文章文档

## 点赞功能

用户点赞可以查看自己的点赞明细

### 表字段

#### portal_like_post

|  字段   | 长度  | 描述 | 备注 |
|  ----  | ----  | ----  | ----  |
| id  | bigint(20)  |   | 自增id |
| post_id  | bigint(20)  | 文章id  |  |
| user_id  | bigint(20)  | 所属用户id  |  |
| status  | tinyint(3)  | 点赞状态，默认值为：1  | 0：未点赞，1：点赞 |
| create_at  | bigint(20)  | 创建时间  |  |
| update_at  | bigint(20)  | 更新时间  |  |

### 操作逻辑

是否已经点赞，存在数据则更改点赞状态。否则新增一条点赞关联信息

### 定义方法

#### 查找单条：show()
查询当前文章是否已经被该用户操作过

#### 新增单条：create()
插入一条数据到表中

#### 修改单条：save()
修改一条数据到表中

### go结构体
```
type PostLikePost struct {
	Id         int    `json:"id"`
	PostId     int    `gorm:"type:bigint(20);comment:文章id;not null" json:"post_id"`
	UserId     int    `gorm:"type:bigint(20);comment:用户id;not null" json:"user_id"`
	Status     int    `gorm:"type:tinyint(3);comment:状态,1:点赞;0:未点赞;default:1;not null" json:"status"`
	CreateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"create_at"`
	UpdateAt   int64  `gorm:"type:bigint(20);NOT NULL" json:"update_at"`
	CreateTime string `gorm:"-" json:"create_time"`
	UpdateTime string `gorm:"-" json:"update_time"`
}
```


