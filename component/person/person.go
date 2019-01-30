package person

import (
	"Martial-arts/code"
	"fmt"
)

// Person component of world
type Person struct {
	Identify uint64
	Channel  chan []byte
	Status   int
}

const (
	STATUSRUNNING = iota
	STATUSSTOP
)

// Run maintain a session of person
func (p *Person) Run() {
	defer func() {
		p.Status = STATUSSTOP
		fmt.Printf("person:%v stop work\n", p.Identify)
	}()
	for {
		fmt.Printf("person:%v process data\n", p.Identify)
		data := <-p.Channel
		info, err := code.Unmarshal(data)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("person:%v will exec the event witch eventid is :%v and from :%v", p.Identify, info.EventID, info.From)
	}
}
