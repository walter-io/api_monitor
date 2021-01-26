package main

import (
	"apiMonitor/self_tail"
)

func main()  {
	tailPackage := self_tail.SelfTail{}
	tailPackage.Run()

	//fmt.Scan()

	//top.Top()
}