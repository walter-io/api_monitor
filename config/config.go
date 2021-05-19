package config

type Config struct {}

const (
    SaveResit     = true
    RedisHost     = "192.168.0.235"
    RedisPort     = "6379"
    RedisAuth     = ""
    RedisPassword = ""
    RedisProtocol = "tcp"

    SaveMysql     = true
    MysqlHost     = "192.168.0.235"
    MysqlPort     = "3306"
    MysqlDatabase = "monitor"
    MysqlUsername = "root"
    MysqlPassword = "123456"

    //TailFile = "www.lara.com.access_log.log"
    TailFile = "www.phpshjgame.com_access_log.log"
)



