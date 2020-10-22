package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/alibaba/sentinel-golang/util"

	"github.com/alibaba/sentinel-golang/core/flow"

	sentinel "github.com/alibaba/sentinel-golang/api"
)

func main() {
	// 务必先初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.FlowRule{
		{
			Resource:        "pledge",
			MetricType:      flow.QPS,
			Count:           10,
			ControlBehavior: flow.Reject,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 埋点逻辑，埋点资源名为pledge
				e, b := sentinel.Entry("pledge")
				if b != nil {
					// 请求被拒绝，在此处进行处理
					fmt.Println("reject", b)
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					// 请求允许通过，此处编写业务逻辑
					fmt.Println(util.CurrentTimeMillis(), "Passed")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

					// 务必保证业务结束后调用Exit
					e.Exit()
				}
			}
		}()
	}

	<-ch
}
