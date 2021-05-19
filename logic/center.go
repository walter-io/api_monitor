package logic

import "apiMonitor/model"

type (
    /**
     * 控制中心
     */
    Center struct {
        ParseResult chan model.Detail
        TailRows     chan string
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
        ParseQueueSubmit(detail model.Detail)
    }
)

/**
 * 建立队列
 */
func (c *Center) ParseQueue() {
    c.ParseResult = make(chan model.Detail, 10)
}

/**
 * 把数据提交给队列
 */
func (c *Center) ParseQueueSubmit(detail model.Detail) {
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