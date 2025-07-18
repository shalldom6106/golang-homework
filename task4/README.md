# blog_test api

这是基于 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

## 项目运行环境

- Go 1.24.0
- MySQL8.0.36 数据库
- Windows11 操作系统

## 项目结构

```
task4/
├── config/         # 配置及初始化
├── controllers/    # 控制器,逻辑处理
├── middleware/     # 中间件
├── models/         # 数据模型定义
├── routers/        # 路由配置
├── utils/          # 工具函数
├── .env            # 配置参数
└── main.go         # 应用入口
```

## 功能特性

- 项目初始化
- 数据库设计与模型定义
- 用户认证与授权
- 文章管理功能
- 评论功能
- 错误处理与日志记录

## 依赖安装步骤

   # Web框架
   go get -u github.com/gin-gonic/gin@v1.10.1
   
   # GORM框架
   go get -u gorm.io/gorm@v1.30.0
   go get -u gorm.io/driver/mysql@v1.6.0
   
   # JWT认证
   go get -u github.com/golang-jwt/jwt/v5@v5.2.2
   
   # 环境变量加载
   go get -u github.com/joho/godotenv@v1.5.1
   
   # 日志
   go get -u go.uber.org/zap@v1.27.0
   
   # 密码加密
   go get -u golang.org/x/crypto@v0.39.0


## 启动方式

1. 环境变量和数据库确保已配置完成

2. 在项目根目录执行：go run main.go

3. 服务将在`http://localhost:8080`启动
（注意：后端请求需要在"Authorization"中选择Bearer Token，再输入登录成功后的token进行身份验证，所有POST和PUT请求的数据格式为JSON）

## 后端路径

### 用户认证

- `POST /auth/register` - 用户注册
- `POST /auth/login` - 用户登录

### 文章管理

- `POST /posts/createpost` - 创建文章（需要认证）
- `GET /posts/findposts` - 获取所有文章
- `GET /posts/findpostbyid/:id` - 获取单篇文章
- `PUT /posts/updatepost/:id` - 更新文章（需要认证）
- `DELETE /posts/deletepost/:id` - 删除文章（需要认证）

### 评论管理

- `POST /comments/createcomment` - 发表评论（需要认证）
- `GET /comments/findcomment/:postID` - 获取文章评论
