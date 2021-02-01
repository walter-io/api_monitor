package top

import (
	"apiMonitor/drivers"
	"fmt"
)

type House struct {
	ID    string `table:"ID"`
	Name  string `table:"接口"`
	//Sigil string `table:"IP"`
	Motto int    `table:"请求次数"`
}

func Top() {

	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("some input")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String())

	//for {
	//	fmt.Println("Over")
	//	time.Sleep(time.Second * 1)
	//	cmd := exec.Command("cmd", "/c", "cls")
	//	cmd.Stdout = os.Stdout
	//	cmd.Run()
	//}

	// 无刷新加次数
	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	//for i :=0;i!=10;i=i+1{
	//	fmt.Fprintf(os.Stdout,"result is %d\r",i)
	//	time.Sleep(time.Second*1)
	//}
	//fmt.Println("Over")

	// 打印表格
	clientRedis := *drivers.ClientRedis
	// 获取成员：前10条
	reply, _ := clientRedis.Do("ZREVRANGE", "api_monitor", 1, 10)
	// 通过成员获取分数:ZSCORE TODO HERE
	fmt.Printf("%+s\n", reply)
	//n := 0
	//for {
	//	// 从redis拿数据
	//
	//
	//	s := []House{
	//		//{"1", "api/login/login", "192.168.0.138", 3},
	//		//{"2", "api/user/getUserInfo", "192.168.0.138", 2},
	//		//{"3", "api/user/detail", "192.168.0.138", n},
	//		{"1", "api/login/login", 3},
	//		{"2", "api/user/getUserInfo", 2},
	//		{"3", "api/user/detail", n},
	//	}
	//	t := table.Table(s)
	//	fmt.Fprintf(os.Stdout, "%s\r", t)
	//	n++
	//	time.Sleep(time.Second*1)
	//	break
	//
	//	//app := "clear"
	//	//cmd := exec.Command(app)
	//	//stdout, err := cmd.Output()
	//	//
	//	//if err != nil {
	//	//	println(err.Error())
	//	//	return
	//	//}
	//	//print(string(stdout))
	//}


}
