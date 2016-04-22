package main

import (
	"fmt"
	"time"

	"github.com/colinyl/stomp"
)

func main() {
	mq, err := stomp.NewStomp("192.168.101.161:61613")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for i := 1; i < 100; i++ {
			err = mq.Send("go:active_mq", fmt.Sprintf("%d", i))
			time.Sleep(time.Second)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Printf("send:%d\r\n", i)
			}
		}
	}()

	mq.Consume("go:active_mq", func(m stomp.MsgHandler) {
		fmt.Printf("recv:%s\r\n", m.GetMessage())
		m.Ack()
	})

}
