package redis

import (
	"fmt"
	"testing"
)

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/3 11:25 上午
 * @Desc:
 */

var (
	maxPoolSize = 1000
	maxIdle     = 50
	idleTimeout = 180
)

func NewRedis(host string, port int64, db int64, password string) (*ToolRedis, error) {
	// 这里的参数可以从配置文件中获取
	t := &ToolRedis{}
	address := fmt.Sprintf("%s:%d", host, port)
	if err := t.Connect(
		address,
		password,
		int(db),
		maxIdle,
		maxPoolSize,
		idleTimeout,
	); err != nil {
		return nil, err
	}
	return t, nil
}

func TestToolRedisConnect(t *testing.T) {
	tRedis, err := NewRedis(
		"127.0.0.1",
		6379,
		0,
		"",
	)
	if err != nil {
		t.Fatalf("err=[%s]", err.Error())
	}
	pool := tRedis.GetPool()
	cli := tRedis.GetCli()
	t.Logf("tRedis=[%+v], pool=[%+v], cli=[%+v]", *tRedis, *pool, cli)
}
