# ZeroCMF完全开发手册

## 平台简介
ZeroCMF是一款通用快速开发框架，在保持极致的性能下依然可以帮您快速实现业务

* 基于MIT协议，免费开源，可商业使用
* go-zero + gorm + casbin，享受原滋原味的golang开发

## 环境要求
> mysql  
> redis  
> nginx  
> apisix // 微服务必须  
> etcd   // 微服务必须

## 框架特性
* MIT协议
* 基于go-zero，渐进式开发
* 前台采用react + umi 框架，支持ssr，对SEO友好
* 核心化：独立核心代码包
* 应用化：开发者以应用的形式增加项目模块
* 插件化：更强的插件机制，开发者以插件形式扩展功能
* 模板化：模板完全傻瓜式，用户无须改动任何代码即可在后台完成模板设计和配置 [x]
* 增加URL美化功能，支持别名设置，更简单
* 统一的资源管理，相同文件只保存一份 [x]
* 文件存储插件化，默认支持七牛文件存储插件 [x]

## 特色功能
* 菜单管理
* 用户管理
* 角色管理
* 权限管理
* 文件资源管理

## 官方服务（插件）
* 门户系统 - 配合大量模板和实现快速建站
* 评论系统 - 可快速实现反馈，论坛，社区，评论等
* 一键登录 - 集成常见的三方登录，如：微信，QQ，微博等

## 目录介绍
```
zerocmf 根目录
├─common 通用模块
└─ ...
├─service 内置服务
│  ├─admin 核心管理模块
│  ├─user  用户模块
│  ├─portal 门户模块
│  └─ ...
```

### 快速开始
推荐使用docker-compose一键运行脚本
[zerocmf docker-compose](https://github.com/zerocmf-team/docker-compose/tree/zeroCmf)

```
cd ~
git clone git@github.com:zerocmf-team/docker-compose.git
git checkout zerocmf
cd workspace
docker-compose up --build -d
```
