package tail

import (
	"fmt"
	"regexp"
	"strconv"
)

/**
 * 详情
 */
type Detail struct {
	RemoteAddr string
	RemoteUser string
	Time       string
	Method     string
	RequestUrl string
	Protocol   string
	Status     int
	Size       float64
	OriginUrl  string
	UserAgent  string
}

/**
 * 定义接口
 */
type CenterInterface interface {
	Queue()
	Submit(detail Detail)
	Run()
}

/**
 * 中心结构体
 */
type Center struct {
	CenterInterface CenterInterface
	ParseResult chan Detail
}

/**
 * 新建队列
 */
func (c *Center) Queue()  {
	c.ParseResult = make(chan Detail)
}

/**
 * 提交到队列
 */
func (c *Center) Submit(detail Detail) {

}

/**
 * 开始处理
 */
func (c *Center) Run()  {
	url := "192.168.0.68 - - [19/Jan/2021:19:09:05 +0800] \"GET /admin/gameManage.ReleaseVersion/add.html HTTP/1.1\" 200 12557 \"http://www.phpshjgame.com/admin/gameManage.ReleaseVersion/index.html\" \"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36\""
	temp := fetch(url)
	fmt.Printf("%+v\n", temp)

	// 做个并发保存数据：放对列 -> 队列存数据库和放redis


	//go func() {
	//	url := "192.168.0.68 - - [19/Jan/2021:19:09:05 +0800] \"GET /admin/gameManage.ReleaseVersion/add.html HTTP/1.1\" 200 12557 \"http://www.phpshjgame.com/admin/gameManage.ReleaseVersion/index.html\" \"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36\""
	//	temp := fetch(url)
	//	c.ParseResult <- temp
	//}()
	//

	//
	//for {
	//	t := <- c.ParseResult
	//	fmt.Printf("%+v\n", t)
	//}
}

func Tail() {
	c := Center{}
	c.Run()
	/*
	$remote_addr 客户端地址 211.28.65.253
	$remote_user 客户端用户名称 --
	$time_local 访问时间和时区 18/Jul/2012:17:00:01 +0800
	$request 请求的URI和HTTP协议 "GET /article-10000.html HTTP/1.1"
	$http_host 请求地址，即浏览器中你输入的地址（IP或域名） www.it300.com
	192.168.100.100
	$status HTTP请求状态 200
	$upstream_status upstream状态 200 、、、、
	$body_bytes_sent 发送给客户端文件内容大小 1547
	$http_referer url跳转来源 https://www.baidu.com/
	$http_user_agent 用户终端浏览器等信息 "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; GTB7.0; .NET4.0C;
	$ssl_protocol SSL协议版本 TLSv1
	$ssl_cipher 交换数据中的算法 RC4-SHA
	$upstream_addr 后台upstream的地址，即真正提供服务的主机地址 10.10.10.100:80
	$request_time 整个请求的总时间 0.205
	$upstream_response_time 请求过程中，upstream响应时间 0.002
	192.168.0.68 - [19/Jan/2021:14:16:06 +0800] GET /admin/gameManage.gameRegisterEmail/add.html HTTP/1.1 200 31595 http://www.phpshjgame.com/admin/gameManage.GameRegisterEmail/index.html Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 - 3.580
	*/

	//t, err := tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true})
	//if err != nil {
	//	panic(err)
	//}

	//var c &Center{}
	//url := "192.168.0.68 - - [19/Jan/2021:19:09:05 +0800] \"GET /admin/gameManage.ReleaseVersion/add.html HTTP/1.1\" 200 12557 \"http://www.phpshjgame.com/admin/gameManage.ReleaseVersion/index.html\" \"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36\""
	//temp := fetch(url)
	//c.ParseResult <- temp

	//data := []Detail{}
	//for line := range t.Lines {
	//	temp := fetch(line.Text)
	//	c.ParseResult <- temp
	//	fmt.Printf("%v", temp)
	//
	//}

}



/**
 * 解析内容
 */
func fetch(url string) Detail {
	strRegexp := `(\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})\s-\s(.*)\[(.*)\]\s\"(.*)\s(.*)\s(.*)\"\s(\d+)\s(\d+)\s\"(.*)\"\s\"(.*)\"`
	resp 	  := regexp.MustCompile(strRegexp)
	body 	  := resp.FindAllStringSubmatch(url, -1)

	detail := Detail{}
	for _, match := range body {
		detail.RemoteAddr = match[1]
		detail.RemoteUser = match[2]
		detail.Time 	  = match[3]
		detail.Method 	  = match[4]
		detail.RequestUrl = match[5]
		detail.Protocol   = match[6]
		detail.Status, _  = strconv.Atoi(match[7])
		detail.OriginUrl  = match[9]
		detail.UserAgent  = match[10]
		size, _ := strconv.ParseFloat(match[8], 64)
		latestSize := strconv.FormatFloat((size / 1024 / 1024), 'f', 2, 64)
		detail.Size, _ = strconv.ParseFloat(latestSize, 64)
	}
	return detail
}
