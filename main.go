package main

import (
    "apiMonitor/self_tail"
)

func main()  {
    // 获取日志中的接口地址
    tailPackage := self_tail.SelfTail{}
    tailPackage.Run()

    // 显示接口访问排行
    //self_top.Top()
}


