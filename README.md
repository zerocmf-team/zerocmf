# Gincmf1.0.0 让开发变得简单而快乐

## Gin1.0.0主要特性
* 强劲内核，基于Gin框架打造
* mvc框架支持
* 路由支持
* 多主题支持
* REST API支持
* 独立核心代码包cmf

## 环境推荐
> Go1.11+  

> nginx-1.17.10+ 

> mysql5.7+
## 使用手册 
### 安装教程
第一步：git clone https://github.com/gimcmf/gincmf  
第二步：下载全部依赖 go get -u  
第三步：运行 默认端口为8000
### 路径解释
gincmf  
|-app //应用目录  
| |-controller //控制器目录  
| | |-api //api接口目录  
| | |-web //前台应用目录  
| |-middleware //中间件目录  
| |_model //模型层目录  
|-conf //配置文件目录  
| | |-config.json //配置文件  
|-public //公共目录  
| |\_themes //主题目录  
| | |-default //默认主题目录  
| | |\_...others //其他主题目录  
|-router //路由目录  
| | |-api.go //接口路由文件  
| | |-web.go //前台路由文件  
|-go.mod  //模块依赖  
|-go.sum  //模块依赖  
|-main.go  //入口文件  
|_readme.md
### 下载gincmf代码
推荐使用git下载  
> git clone https://github.com/gimcmf/gincmf  