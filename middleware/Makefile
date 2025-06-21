# ==============================================================================
# Go Application Management
# ==============================================================================
.PHONY: run build clean tidy test

# 设置变量
APP_NAME := my-redis-app
CMD_PATH := cmd/app/main.go
OUTPUT_DIR := bin

# 运行 Go 应用程序
run:
	@echo "🏃 Running the application..."
	go run $(CMD_PATH)

# 编译 Go 应用程序
build:
	@echo "🔨 Building the application..."
	go build -o $(OUTPUT_DIR)/$(APP_NAME) $(CMD_PATH)
	@echo "✅ Build complete: $(OUTPUT_DIR)/$(APP_NAME)"

# 清理编译产物
clean:
	@echo "🧹 Cleaning up built files..."
	rm -rf $(OUTPUT_DIR)
	@echo "✅ Cleanup complete."

# 整理 Go 模块依赖
tidy:
	@echo "📦 Tidying Go modules..."
	go mod tidy
	@echo "✅ Modules tidied."

# 运行单元测试
test:
	@echo "🧪 Running tests..."
	go test ./... -v

# ==============================================================================
# Docker Redis Management
# ==============================================================================
# Redis Docker 管理 Makefile

.PHONY: start-redis stop-redis restart-redis logs-redis

# 启动 Redis 服务
start-redis:
	@echo "🚀 启动 Redis 服务..."
	cd deploy/ && docker-compose up -d
	@echo "✅ Redis 服务已启动"

# 停止 Redis 服务
stop-redis:
	@echo "🛑 停止 Redis 服务..."
	cd deploy/ && docker-compose down
	@echo "✅ Redis 服务已停止"

# 重启 Redis 服务
restart-redis:
	@echo "🔄 重启 Redis 服务..."
	cd deploy/ && docker-compose restart
	@echo "✅ Redis 服务已重启"

# 查看 Redis 日志
logs-redis:
	@echo "📝 查看 Redis 日志..."
	cd deploy/ && docker-compose logs -f redis
