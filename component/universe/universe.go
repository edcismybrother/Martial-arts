package universe

import (
	"Martial-arts/component/person"
	"Martial-arts/component/planet"
	"fmt"
)

// Universe a bigest area
type Universe struct {
	ID      int // Uniquely identifies
	Planets []*planet.Planet
	Channel chan []byte
}

func (u *Universe) Run() {
	fmt.Println("start up universe:", u.ID)
	for data := range u.Channel {
		u.processMsg(data)
	}
}

func (u *Universe) processMsg(data []byte) {
	fmt.Println(string(data))
}

func (u *Universe) SendMsg(data []byte) {
	fmt.Printf("universe:%v begin to broadcast message\n", u.ID)
	for _, planet := range u.Planets {
		for _, world := range planet.Worlds {
			for _, ps := range world.Persons {
				if ps.Status == person.STATUSRUNNING {
					fmt.Printf("person:%v received message from universe:%v-planet:%v-world:%v\n", ps.Identify, u.ID, planet.ID, world.ID)
					ps.Channel <- data
				}
			}
		}
	}
}
