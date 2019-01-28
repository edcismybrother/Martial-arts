package main

import (
	"Martial-arts/global"
	"time"
)

func main() {
	var ch = make(chan bool)
	m := global.NewComplateManager()
	m.Run()
	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			//模拟通信
			<-t.C
			m.UniverseManager[0].SendMsg([]byte("Hello world"))
		}
	}()
	<-ch
}
