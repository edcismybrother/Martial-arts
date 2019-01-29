package planet

import (
	"Martial-arts/component/world"
	"fmt"
)

// Planet a big area
type Planet struct {
	ID      int // Uniquely identifies
	Worlds  []*world.World
	Channel chan []byte
}

func (p *Planet) Run() {
	fmt.Println("启动星球")
	for data := range p.Channel {
		fmt.Println("接收到消息")
		fmt.Println(string(data))
	}
}
