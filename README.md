# HR_Go [![Build Status](https://ci.scutbot.cn/api/badges/Web/HR_Go/status.svg)](https://ci.scutbot.cn/Web/HR_Go)

基于 GoLang 的面试系统后端

## 启动

启动 Docker 容器

``` shell
cd unittest && docker-compose up -d
```

启动 Service 层

``` shell
cd hr-service && go run .
cd hr-admin-service && go run .
```

启动 API 层

``` shell
cd hr-api && go run .
cd hr-admin-api && go run .
```

检测服务启动情况

```shell
curl --request GET 'http://localhost:8000/api/ping'
curl --request GET 'http://localhost:8000/api/admin/ping'
```

## 启动 （微服务容器组）

```shell
cd unittest
vi docker-compose-aio.yml
```

### 自定义容器组，如：

- 增加 service 微服务实例
- 修改端口映射
- 添加容器数据库
- 修改为从 registry 拉取(取消image: 注释)，而非本地构建

修改完成后启动容器

```shell
docker-compose -f docker-compose-aio.yml up -d
```

## 目录

``` text
.
├── MakeFile
├── README.md
├── build 项目构建目录
├── common 项目公共模块
├── dal 数据访问层
├── go.mod
├── go.sum
├── hr-admin-api 管理员 API 层
├── hr-admin-service 管理员 Service 层
├── hr-api 面试者 API 层
├── hr-service 面试者 Service 层
├── unittest 单元测试
└── util 工具模块

```

## 技术栈

### [go-zero](https://github.com/zeromicro/go-zero)

- API (Application Programming Interface)
- Service

### [gorm](https://github.com/go-gorm/gorm)

- DAL (Data Access Layer)

没有使用 go-zero 原生的数据库框架

### [Docker](https://docs.docker.com)

- Nginx
    - nginx:latest
- MySQL
    - mysql:8.0
    - arm64v8/mysql:8.0
