package top

import (
    "C"
    "apiMonitor/drivers"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/liushuochen/gotable"
    "os"
    "os/exec"
    "time"
)

type House struct {
    ID    string `table:"ID"`
    Name  string `table:"接口"`
    //Sigil string `table:"IP"`
    Motto int    `table:"请求次数"`
}

func Top() {
    // TODO HERE 显示在表格中，并定时刷新
    for {
        headers := []string{"interface", "count"}
        tb, err := gotable.CreateTable(headers)
        if err != nil {
            fmt.Println("Create table failed: ", err.Error())
            return
        }

        // 获取redis成员：前10条  TODO 每次拿出来的数据顺序不一样
        clientRedis := *drivers.ClientRedis
        reply, err := redis.StringMap(clientRedis.Do("ZREVRANGE", "api_monitor", 0, 10, "WITHSCORES"))
        if err != nil {
            panic(err)
        }
        for k, v := range reply {
            value := gotable.CreateEmptyValueMap()
            value["interface"] = gotable.CreateValue(k)
            value["count"] = gotable.CreateValue(v)
            tb.AddValue(value)
        }
        tb.PrintTable() // 显示表格

        time.Sleep(time.Second * 3) // 停留3秒

        clearConsole() // 清屏
    }
}

func clearConsole() {
    // 清屏
    fmt.Println("hello world")
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
    //fmt.Print("\033c")
    //fmt.Print("\x1bc")
}