package self_tail

import (
    "apiMonitor/logic"
    "apiMonitor/config"
    "apiMonitor/drivers"
    "apiMonitor/model"
    "github.com/hpcloud/tail"
    "regexp"
    "strconv"
    "strings"
)

/**
 * 中心结构体
 */
type SelfTail struct {
    Center logic.Center
}

/**
 * 开始处理: 抓包只需要一个协程, 解析器需要多个协程, 存数据也要多个协程
 */
func (s *SelfTail) Run() {
    // 开协程抓包放队列
    go func() {
        s.Center.TailQueue()
        // 开始抓包
        t, err := tail.TailFile(config.TailFile, tail.Config{Follow: true})
        if err != nil {
            panic(err)
        }
        for line := range t.Lines {
            s.Center.TailSubmitQueue(line.Text)
        }
    }()

    // 开协程从抓包队列中获取数据进行解析
    go func() {
        s.Center.ParseQueue()
        for {
            // 拿access_log行
            row := <- s.Center.TailRows
            // 交给解析器解析
            temp := s.Parse(row)
            // 把数据提交给通道
            s.Center.ParseQueueSubmit(temp)
        }
    }()

    // 开协程从通道拿数据存数据库和redis
    clientRedis := *drivers.ClientRedis
    for {
        select {
        case t := <-s.Center.ParseResult:
            // 存redis
            if config.SaveResit {
               clientRedis.Do("ZADD", "api_monitor", "INCR", 1, t.OriginUrl)
            }
            // 存mysql
            if config.SaveMysql {
                model.InsertDetail(t)
            }
        default:
            //fmt.Println("没有数据\n")
        }
    }
}

/**
 * 解析内容 TODO 拆分成独立的解析器
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
func (s *SelfTail) Parse(row string) model.Detail {
    // 匹配nginx日志的正则
    strRegexp := `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s(.*)\[(.*)\]\s\"(.*)\s(.*)\s(.*)\"\s(\d+)\s(\d+)\s\"(.*)\"\s\"(.*)\"\s"(.*)\"\s\"(.*)\"`
    resp       := regexp.MustCompile(strRegexp)
    body       := resp.FindAllStringSubmatch(row, -1)

    detail := model.Detail{}
    for _, match := range body {
        size, _ := strconv.ParseFloat(match[8], 64)
        latestSize := strconv.FormatFloat((size / 1024 / 1024), 'f', 2, 64)
        detail.Size, _ = strconv.ParseFloat(latestSize, 64)
        requestTime, _ :=  strconv.ParseFloat(match[12], 64)     // 发送给客户端文件内容大小

        detail.RemoteAddr   = match[1]                  // 客户端ip
        detail.RemoteUser   = match[2]                  // 客户端名称
        detail.Time         = match[3]                  // 访问时间和失去
        detail.Method       = match[4]                  // http 请求方法
        detail.RequestUrl   = match[5]                  // 请求url
        detail.Protocol     = match[6]                  // ssl协议版本
        detail.Status, _    = strconv.Atoi(match[7])    // http请求状态
        detail.OriginUrl    = match[9]                  // 跳转来源
        detail.UserAgent    = match[10]                 // 用户终端浏览器信息
        detail.RequestTime  = requestTime               // 请求总时间

        // 链接去掉参数
        symbolIndex := strings.Index(detail.OriginUrl, "?")
        if symbolIndex > 0 {
            detail.OriginUrl = detail.OriginUrl[0:symbolIndex]
        }
    }

    return detail
}

