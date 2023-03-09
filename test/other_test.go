package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
)

var eventMap = make(map[string]interface{})
var eventChan = make(chan interface{})

func YUU() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case evt := <-eventChan:
			eventMap[uuid.New().String()] = evt
			if len(eventMap) >= 100 {
				// 触发发送事件
				log.Println("普通")
				sendEvents(eventMap)
				eventMap = make(map[string]interface{})
			}
		case <-ticker.C:
			log.Println("触发定时器")
			if len(eventMap) > 0 {
				log.Println("定时器")
				sendEvents(eventMap)
				eventMap = make(map[string]interface{})
			}
		}
	}
}

func sendEvents(events map[string]interface{}) {
	// 发送事件的代码
	fmt.Println("Send events:", events)
}

func TestYUU(t *testing.T) {
	go YUU()

	for i := 0; i < 120; i++ {
		eventChan <- i
	}

	time.Sleep(time.Second * 15)

	eventChan <- "1231"
	for {
		select {}
	}
}
