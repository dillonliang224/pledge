package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/util"
)

func main() {
	// 初始化配置
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置规则
	_, err = flow.LoadRules([]*flow.FlowRule{
		{
			Resource:        "some-test",
			MetricType:      flow.QPS,
			Count:           10,
			ControlBehavior: flow.Throttling,
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
				e, b := sentinel.Entry("some-test")
				if b != nil {
					// 请求被拒绝
					fmt.Println(util.CurrentTimeMillis(), b)
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					// 请求允许通过
					fmt.Println(util.CurrentTimeMillis(), "Passed")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

					// 业务结束后调用exit
					e.Exit()
				}
			}
		}()
	}

	<-ch
}
