# Redis Docker 项目

这个项目演示了如何使用Docker启动Redis服务，并通过Go程序进行连接和操作。

## 项目结构

```
.
├── docker-compose.yaml    # Docker Compose配置文件
├── redis.conf            # Redis配置文件
├── redis/
│   └── main.go          # Go Redis客户端程序
└── README.md            # 项目说明文档
```

## 快速开始

### 1. 启动Redis服务

使用Docker Compose启动Redis服务：

```bash
docker-compose up -d
```

这将启动一个Redis容器，数据会持久化到Docker卷中。

### 2. 检查服务状态

```bash
docker-compose ps
```

### 3. 查看Redis日志

```bash
docker-compose logs redis
```

### 4. 运行Go程序

确保你已经安装了Go和Redis客户端库：

```bash
# 安装Redis客户端库
go mod init redis-demo
go get github.com/redis/go-redis/v9

# 运行程序
cd redis
go run main.go
```

### 5. 停止服务

```bash
docker-compose down
```

## 数据持久化

这个配置使用了两种持久化方式：

1. **RDB持久化**：定期将内存数据快照保存到磁盘
2. **AOF持久化**：记录每个写操作，提供更好的数据安全性

数据存储在Docker卷 `redis_data` 中，即使容器重启或删除，数据也不会丢失。

## 配置说明

### Docker Compose配置

- **端口映射**：`6379:6379` - 将容器的6379端口映射到主机的6379端口
- **数据卷**：`redis_data:/data` - 数据持久化
- **配置文件**：`./redis.conf:/usr/local/etc/redis/redis.conf` - 自定义Redis配置
- **重启策略**：`unless-stopped` - 除非手动停止，否则自动重启

### Redis配置

- **RDB持久化**：每900秒（15分钟）如果有1个键变化就保存
- **AOF持久化**：启用，每秒同步一次
- **内存限制**：256MB
- **内存策略**：LRU淘汰策略

## 常用命令

### Docker Compose命令

```bash
# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 重启服务
docker-compose restart

# 查看日志
docker-compose logs -f redis

# 进入容器
docker-compose exec redis redis-cli
```

### Redis CLI命令

```bash
# 连接到Redis
docker-compose exec redis redis-cli

# 在Redis CLI中的常用命令
PING                    # 测试连接
SET key value          # 设置键值对
GET key                # 获取值
DEL key                # 删除键
KEYS *                 # 查看所有键
FLUSHALL              # 清空所有数据
INFO                  # 查看服务器信息
```

## 注意事项

1. 确保Docker和Docker Compose已安装
2. 确保6379端口未被其他服务占用
3. 首次启动可能需要下载Redis镜像，请耐心等待
4. 数据会持久化在Docker卷中，删除容器不会丢失数据

## 故障排除

### 端口冲突

如果6379端口被占用，可以修改 `docker-compose.yaml` 中的端口映射：

```yaml
ports:
  - "6380:6379"  # 改为6380端口
```

然后更新Go程序中的连接地址：

```go
Addr: "localhost:6380"
```

### 连接失败

1. 检查Redis容器是否正常运行：`docker-compose ps`
2. 检查端口是否正确：`netstat -an | grep 6379`
3. 查看Redis日志：`docker-compose logs redis` 