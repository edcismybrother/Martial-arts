package main

import (
	"Martial-arts/global"
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)
	m := global.NewComplateManager()
	m.Run()
	t := time.NewTicker(1 * time.Second)
	generatePersonTime := time.NewTicker(1 * time.Second)
	go func() {
		defer func() {
			fmt.Println("end sendMessage")
		}()
		for {
			select {
			case <-t.C:
				{
					//模拟通信
					fmt.Println("SendMessage")
					m.UniverseManager[0].SendMsg([]byte("Hello world"))
				}
			}
		}
	}()

	go func() {
		for {
			//模拟通信
			select {
			case <-generatePersonTime.C:
				{
					fmt.Println("GeneratePerson")
					m.GeneratePerson()
				}
			}
		}
	}()
	<-ch
}

func main2() {
	var ch chan bool
	t1 := time.NewTicker(1 * time.Second)
	c1 := make(chan int, 1)
	t2 := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-t1.C:
				c1 <- 1
				fmt.Println("t1")
			}
		}
	}()
	func() {
		for {
			select {
			case <-t2.C:
				fmt.Println("t2")
			}
		}
	}()
	<-ch
}

func main1() {
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		fmt.Println(i)
	}
}
