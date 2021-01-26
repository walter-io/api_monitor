package center

/**
 * 控制中心
 */
type Center struct {
	ParseResult chan Detail
}

/**
 * 定义接口
 */
type CentralInterface interface {
	Queue()
	Submit(detail Detail)
}

/**
 * 数据详情
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
 * 建立队列
 */
func (c *Center) Queue()  {
	c.ParseResult = make(chan Detail, 10)
}

/**
 * 把数据提交给队列
 */
func (c *Center) Submit(detail Detail)  {
	//fmt.Printf("%+v:", detail)
	c.ParseResult <- detail
}