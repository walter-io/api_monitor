package self_tail

import (
    "apiMonitor/center"
    "apiMonitor/config"
    "fmt"
    "github.com/hpcloud/tail"
    "regexp"
    "strconv"
    "strings"
)

/**
 * 中心结构体
 */
type SelfTail struct {
    Center center.Center
}

/**
 * 开始处理: 抓包只需要一个协程, 解析器需要多个协程, 存数据也要多个协程
 */
func (c *SelfTail) Run() {

    //row := "192.168.0.68 - - [19/Jan/2021:19:09:05 +0800] \"GET /admin/gameManage.ReleaseVersion/add.html HTTP/1.1\" 200 12557 \"http://www.phpshjgame.com/admin/gameManage.ReleaseVersion/index.html\" \"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36\""
    //row := "192.168.0.68 - - [29/Jan/2021:16:01:48 +0800] \"GET /admin/adminManage.Purview/edit/purviewId/225 HTTP/1.1\" 200 14971 \"http://www.phpshjgame.com/admin/adminManage.Purview/list.html?name=shuisheng\" \"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36\""
    //// 交给解析器解析
    //c.parse(row)
    //panic("结束")

    // 开协程抓包放队列
    go func() {
        c.Center.TailQueue()
        // 开始抓包
        t, err := tail.TailFile(config.TailFile, tail.Config{Follow: true}) // TODO 加配置
        if err != nil {
            panic(err)
        }
        for line := range t.Lines {
            c.Center.TailSubmitQueue(line.Text)
        }
    }()

    // 开协程从抓包队列中获取数据进行解析
    go func() {
        c.Center.ParseQueue()
        for {
            // 拿access_log行
            row := <- c.Center.TailRows
            // 交给解析器解析
            temp := c.Parse(row)
            // 把数据提交给通道
            c.Center.ParseQueueSubmit(temp)
        }
    }()

    // 开协程从通道拿数据存数据库和redis
    //clientRedis := *drivers.ClientRedis
    for {
        select {
        case t := <-c.Center.ParseResult:
            // 存redis
            //if config.SaveResit {
            //    clientRedis.Do("ZADD", "api_monitor", "INCR", 1, t.OriginUrl)
            //}
            // 存mysql TODO HERE 修改nginx配置，detail结构把响应时间加上，再存到数据库
            //if config.SaveMysql {
            //
            //}

            fmt.Printf("%+v\n", t)
        default:
            //fmt.Println("没有数据\n")
        }
    }
}

/**
 * 解析内容
 */
func (c *SelfTail) Parse(row string) center.Detail {
    // 匹配nginx日志的正则
    strRegexp := `(\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})\s-\s(.*)\[(.*)\]\s\"(.*)\s(.*)\s(.*)\"\s(\d+)\s(\d+)\s\"(.*)\"\s\"(.*)\"\s(.*)`
    resp       := regexp.MustCompile(strRegexp)
    body       := resp.FindAllStringSubmatch(row, -1)

    /**
     * 接收解析后的数据RechargeController
     * $remote_addr 客户端地址 211.28.65.253
     * $remote_user 客户端用户名称 --
     * $time_local 访问时间和时区 18/Jul/2012:17:00:01 +0800
     * $request 请求的URI和HTTP协议 "GET /article-10000.html HTTP/1.1"
     * $http_host 请求地址，即浏览器中你输入的地址（IP或域名） www.it300.com
     * 192.168.100.100
     * $status HTTP请求状态 200
     * $upstream_status upstream状态 200 、、、、
     * $body_bytes_sent 发送给客户端文件内容大小 1547
     * $http_referer url跳转来源 https://www.baidu.com/
     * $http_user_agent 用户终端浏览器等信息 "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; SV1; GTB7.0; .NET4.0C;
     * $ssl_protocol SSL协议版本 TLSv1
     * $ssl_cipher 交换数据中的算法 RC4-SHA
     * $upstream_addr 后台upstream的地址，即真正提供服务的主机地址 10.10.10.100:80
     * $request_time 整个请求的总时间 0.205
     * $upstream_response_time 请求过程中，upstream响应时间 0.002
     * 192.168.0.68 - [19/Jan/2021:14:16:06 +0800] GET /admin/gameManage.gameRegisterEmail/add.html HTTP/1.1 200 31595 http://www.phpshjgame.com/admin/gameManage.GameRegisterEmail/index.html Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 - 3.580
     */
    detail := center.Detail{}
    for _, match := range body {
        detail.RemoteAddr   = match[1]
        detail.RemoteUser   = match[2]
        detail.Time         = match[3]
        detail.Method       = match[4]
        detail.RequestUrl   = match[5]
        detail.Protocol     = match[6]
        detail.Status, _    = strconv.Atoi(match[7])
        detail.OriginUrl    = match[9]
        detail.UserAgent    = match[10]
        size, _ := strconv.ParseFloat(match[8], 64)
        latestSize := strconv.FormatFloat((size / 1024 / 1024), 'f', 2, 64)
        detail.Size, _ = strconv.ParseFloat(latestSize, 64)

        // TODO 拿到的数据为空
        requestTime, _ :=  strconv.ParseFloat(match[11], 64)
        detail.RequestTime  = requestTime

        // 链接去掉参数
        symbolIndex := strings.Index(detail.OriginUrl, "?")
        if symbolIndex > 0 {
            detail.OriginUrl = detail.OriginUrl[0:symbolIndex]
        }
        //fmt.Printf("%s\n", detail.OriginUrl)
    }

    return detail
}

