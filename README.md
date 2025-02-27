# qqlx

## 简介

GO 编写的 WEB 框架，此框架基于流行的`Gin`框架构建。

## 核心特性

1. 用户管理：用户注册、登录、资料更新等功能。
2. 角色管理：支持多种角色定义，便于组织结构化权限分配。
3. 权限控制：利用`Casbin`实现精确到资源级别的访问控制。

## 技术栈

1. `Gin`：轻量级的Go web框架，提供路由、中间件支持等功能。
2. `GORM`：强大的对象关系映射（ORM）工具，简化数据库操作。
3. `Casbin`：功能丰富的访问控制库，实现`RBAC`访问模型。
4. `MySQL`：可靠的关系型数据库管理系统，适用于大规模数据存储需求。
5. `Redis`：高性能键值存储系统，用于加速角色信息检索过程。
6. `Wire`：依赖注入

## 接口文档

<https://apifox.com/apidoc/shared-8db2216b-8451-4ead-b0fd-019ce8676f1e>

## 初始化数据

```bash
go build -o .

# env
export CONFIG_PATH=configPath
export CASBIN_MODE_PATH=modelPath
./qqlx init

# 选项
./qqlx init -C configPath -M modelPath
```

## 启动服务

### 二进制启动

```bash
go build -o .

# env
export CONFIG_PATH=configPath
export CASBIN_MODE_PATH=modelPath
./qqlx run

# 选项
./qqlx run -C configPath -M modelPath
```

### Docker 启动

```bash
IMAGE_VERSION=1.0 docker compose up -d
```
