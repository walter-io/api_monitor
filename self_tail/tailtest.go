package self_tail

import (
    "apiMonitor/config"
    "fmt"
    "github.com/hpcloud/tail"
)



func TestSelfTail()  {
    t, err := tail.TailFile(config.TailFile, tail.Config{Follow: true})
    if err != nil {
        panic(err)
    }
    var c = SelfTail{}
    for line := range t.Lines {
        res := c.Parse(line.Text)
        fmt.Printf("%+v\n", res)
    }
}
