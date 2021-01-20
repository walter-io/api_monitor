package main

import (
	"fmt"
	"github.com/modood/table"
	"os"
	"time"
)

type House struct {
	ID    string `table:"ID"`
	Name  string `table:"接口"`
	Sigil string `table:"IP"`
	Motto int    `table:"请求次数"`
}

func main() {
	n := 0
	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	//
	//for i :=0;i!=10;i=i+1{
	//	fmt.Fprintf(os.Stdout,"result is %d\r",i)
	//	time.Sleep(time.Second*1)
	//}
	//fmt.Println("Over")


	for {
		s := []House{
			{"1", "api/login/login", "192.168.0.138", 3},
			{"2", "api/user/getUserInfo", "192.168.0.138", 2},
			{"3", "api/user/detail", "192.168.0.138", n},
		}
		t := table.Table(s)
		fmt.Fprintf(os.Stdout, "%s\r", t)
		n++
		time.Sleep(time.Second*1)
	}


}
