package client

import (
	"MengGoods/pkg/merror"
	"context"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// 初始化Redis客户端
func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	if err := redisotel.InstrumentTracing(client); err != nil {
		return nil, fmt.Errorf("failed to instrument redis tracing: %w", err)
	}

	// 3. 【新增】可选：注入 Metrics 监控（上报到 Prometheus）
	if err := redisotel.InstrumentMetrics(client); err != nil {
		return nil, fmt.Errorf("failed to instrument redis metrics: %w", err)
	}
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, merror.NewMerror(
			merror.InternalDatabaseErrorCode,
			fmt.Sprintf("Redis连接失败: %v", err),
		)
	}
	return client, nil
}