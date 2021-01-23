package main

import (
	"apiMonitor/tail"
	"fmt"
)

func main()  {
	//top.Top()
	tail.Tail()

	fmt.Println("test123")
	fmt.Scan()
}