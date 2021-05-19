package main

import "apiMonitor/controller/self_tail"

func main()  {
	// todo 把redis和mysql配置放好，再把整个包做成api（配置做成传参），在用gin框架做数据展示，访问超量报警
    // 获取日志中的接口地址
    tailPackage := self_tail.SelfTail{}
    tailPackage.Run()
    // 测试
    //self_tail.TestSelfTail()

    // 显示接口访问排行
    //self_top.Top()
}


