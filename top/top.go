package top

import (
	"apiMonitor/drivers"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/liushuochen/gotable"
)

type House struct {
	ID    string `table:"ID"`
	Name  string `table:"接口"`
	//Sigil string `table:"IP"`
	Motto int    `table:"请求次数"`
}

func Top() {

	// 无刷新加次数
	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	//for i :=0;i!=10;i=i+1{
	//	fmt.Fprintf(os.Stdout,"result is %d\r",i)
	//	time.Sleep(time.Second*1)
	//}
	//fmt.Println("Over")



	// TODO HERE 显示在表格中，并定时刷新
	headers := []string{"interface", "count"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	// 获取redis数据 TODO
	clientRedis := *drivers.ClientRedis
	// 获取成员：前10条 TODO 每次拿出来的数据顺序不一样
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

	tb.PrintTable()



	////
	//for i := 0; i < 5; i++ {
	//	temp := make([]interface{}, 2)
	//	temp[0] = s.Index(i)
	//	temp[1] = s.Index(i + 1)
	//	fmt.Printf("%+v\n", temp)
	//	data = append(data, temp)
	//}

	//for i := 0; i < s.Len(); i++ {
	//	fmt.Printf("%s\n", s.Index(i))
	//}

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
