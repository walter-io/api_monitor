package center

type (
    /**
     * 控制中心
     */
    Center struct {
        ParseResult chan Detail
        TailRows     chan string
    }

    /**
     * 数据详情
     */
    Detail struct {
        RemoteAddr  string
        RemoteUser  string
        Time        string
        Method      string
        RequestUrl  string
        Protocol    string
        Status      int
        Size        float64
        OriginUrl   string
        UserAgent   string
        RequestTime float64
    }

    /**
     * 抓包接口
     */
    TailInterface interface {
        TailQueue()
        TailQueueSubmit(string)
    }

    /**
     * 定义接口
     */
    ParserInterface interface {
        ParseQueue()
        ParseQueueSubmit(detail Detail)
    }
)

/**
 * 建立队列
 */
func (c *Center) ParseQueue() {
    c.ParseResult = make(chan Detail, 10)
}

/**
 * 把数据提交给队列
 */
func (c *Center) ParseQueueSubmit(detail Detail) {
    //fmt.Printf("%+v:", detail)
    c.ParseResult <- detail
}

/**
 * 抓包队列
 */
func (c *Center) TailQueue() {
    c.TailRows = make(chan string, 10)
}

/**
 * 提交行到抓包队列
 */
func (c *Center) TailSubmitQueue(row string) {
    c.TailRows <- row
}