package drivers

import (
    "apiMonitor/config"
    "github.com/gomodule/redigo/redis"
)

var ClientRedis *redis.Conn

/**
 * 初始化redis
 */
func init() {
    conn, err := redis.Dial(config.RedisProtocol, config.RedisHost+ ":" + config.RedisPort)
    if err != nil {
        panic(err)
    } else {
        ClientRedis = &conn
    }
}
