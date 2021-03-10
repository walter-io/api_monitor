package self_top

import (
    "C"
    "apiMonitor/drivers"
    "apiMonitor/extends"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "os"
    "time"
)

type House struct {
    ID    string `table:"ID"`
    Name  string `table:"接口"`
    //Sigil string `table:"IP"`
    Motto int    `table:"请求次数"`
}

func Top() {
    str := tempTop()

    // TODO 还是做不了覆盖刷新，只能累计显示
    for i :=0;i!=10;i=i+1{
       fmt.Fprintf(os.Stdout,"%s\r", fmt.Sprint(str))
       time.Sleep(time.Second*1)
    }
    fmt.Println("Over")
}

func tempTop() string {
    // 构建table
    headers := []string{"interface", "count"}
    tb1, err := extends.CreateTable(headers)
    if err != nil {
        fmt.Println("Create table failed: ", err.Error())
        return ""
    }

    // 获取redis成员：前10条  TODO 每次拿出来的数据顺序不一样
    clientRedis := *drivers.ClientRedis
    reply, err := redis.StringMap(clientRedis.Do("ZREVRANGE", "api_monitor", 0, 10, "WITHSCORES"))
    if err != nil {
        panic(err)
    }
    for k, v := range reply {
        value := extends.CreateEmptyValueMap()
        value["interface"] = extends.CreateValue(k)
        value["count"] = extends.CreateValue(v)
        // 往table加元素
        tb1.AddValue(value)
    }

    // 显示表格
    return tb1.PrintTable()
}