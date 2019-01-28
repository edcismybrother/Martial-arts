package universe

import (
	"Martial-arts/component/planet"
	"fmt"
)

// Universe a bigest area
type Universe struct {
	Planets []*planet.Planet
	Channel chan []byte
}

func (u *Universe) Run() {
	fmt.Println("启动宇宙")
	for data := range u.Channel {
		u.processMsg(data)
	}
}

func (u *Universe) processMsg(data []byte) {
	fmt.Println(string(data))
}

func (u *Universe) SendMsg(data []byte) {
	p := u.Planets[0]
	p.Channel <- data
}
