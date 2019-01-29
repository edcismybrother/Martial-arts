package main

import (
	"Martial-arts/global"
	"time"
)

func main() {
	var ch = make(chan bool)
	m := global.NewComplateManager()
	// m := global.NewManager()
	m.Run()
	go func() {
		t := time.NewTicker(2 * time.Second)
		generatePersonTime := time.NewTicker(1 * time.Second)
		for {
			//模拟通信
			select {
			case <-t.C:
				m.UniverseManager[0].SendMsg([]byte("Hello world"))
			case <-generatePersonTime.C:
				m.GeneratePerson()
			}
		}
	}()
	<-ch
}
