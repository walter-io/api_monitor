package main

import (
    "apiMonitor/controller"
)

func main()  {
	// todo 把redis和mysql配置放好，再把整个包做成api（配置做成传参），在用gin框架做数据展示，访问超量报警
    // todo 动态检测，拉起在后台运行
    // 获取日志中的接口地址
    tailPackage := controller.SelfTail{}
    tailPackage.Run()

    // 测试
    //self_tail.TestSelfTail()
}


