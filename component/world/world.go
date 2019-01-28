package world

import (
	"Martial-arts/component/person"
	"fmt"
)

// World a based area
type World struct {
	Persons []*person.Person
	Channel chan []byte
}

func (w *World) Run() {
	fmt.Println("启动世界")
	for data := range w.Channel {
		fmt.Println(string(data))
	}
}
