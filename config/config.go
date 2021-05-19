package config

type Config struct {}

const (
    // redis 连接配置
    SaveResit     = true
    RedisHost     = "192.168.0.235"
    RedisPort     = "6379"
    RedisAuth     = ""
    RedisPassword = ""
    RedisProtocol = "tcp"

    // mysql 连接配置
    SaveMysql     = true
    MysqlHost     = "192.168.0.235"
    MysqlPort     = "3306"
    MysqlDatabase = "monitor"
    MysqlUsername = "root"
    MysqlPassword = "123456"

    // 抓取的日志
    TailFile = "www.phpshjgame.com_access_log.log"
)



