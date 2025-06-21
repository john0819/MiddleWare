package redis

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

// TestRedisClient_SetAndGet_Success 测试 Set 和 Get 方法成功时的场景
func TestRedisClient_SetAndGet_Success(t *testing.T) {
	// 使用 redismock.NewClientMock() 来创建模拟客户端
	db, mock := redismock.NewClientMock()
	client := &redisClient{client: db}

	ctx := context.Background()
	key := "test_key"
	value := "hello_world"

	// 3. 设置期望 (Expectations)
	// 告诉 mock：我期望接下来会有一个 Set 命令被调用。
	// 它应该接收 'test_key', 'hello_world' 和 0 作为参数。
	// 如果这个期望达成了，它应该返回一个表示成功的 StatusCmd。
	mock.ExpectSet(key, value, 0*time.Second).SetVal("OK")

	// 告诉 mock：我还期望接下来会有一个 Get 命令被调用。
	// 它应该接收 'test_key' 作为参数。
	// 如果这个期望达成了，它应该返回 'hello_world'。
	mock.ExpectGet(key).SetVal(value)

	// 4. 执行我们真正要测试的方法
	setResult := client.Set(ctx, key, value, 0)
	getResult := client.Get(ctx, key)

	// 5. 断言结果是否符合预期
	assert.NoError(t, setResult.Err())
	assert.Equal(t, "OK", setResult.Val())

	assert.NoError(t, getResult.Err())
	assert.Equal(t, value, getResult.Val())

	// 6. (非常重要) 检查所有的期望是否都已经被满足了
	// 如果我们设置了期望，但代码没有调用对应的命令，这里会报错。
	err := mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// TestRedisClient_Get_Fail_RedisError 测试当 Redis 返回错误时的场景
func TestRedisClient_Get_Fail_RedisError(t *testing.T) {
	// 使用 redismock.NewClientMock() 来创建模拟客户端
	db, mock := redismock.NewClientMock()
	client := &redisClient{client: db}

	ctx := context.Background()
	key := "non_existent_key"
	expectedError := errors.New("a simulated redis error")

	// 设置期望：当 Get 被调用时，模拟一个错误发生
	mock.ExpectGet(key).SetErr(expectedError)

	// 执行被测试的方法
	result := client.Get(ctx, key)

	// 断言结果
	assert.Error(t, result.Err())
	assert.Equal(t, expectedError, result.Err())

	// 检查期望是否满足
	err := mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

// 注意：我们没有为 NewRedisClient 写单元测试，因为它是一个构造函数，
// 其核心逻辑是创建真实的连接。对它的测试属于"集成测试"的范畴。
