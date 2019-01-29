package component

import (
	"Martial-arts/component/person"
	"Martial-arts/component/planet"
	"Martial-arts/component/universe"
	"Martial-arts/component/world"
	"fmt"
	"sync"
	"sync/atomic"
)

// Manager Global manager
type Manager struct {
	Wg sync.WaitGroup
	// PlanetManager   []*planet.Planet
	UniverseManager []*universe.Universe
	PersonID        *uint64
	// WorldManager    []*world.World
}

// RegisterPlanet add a new planet
func (m *Manager) RegisterPlanet(universeID int, p *planet.Planet) {
	for _, u := range m.UniverseManager {
		if u.ID == universeID {
			fmt.Printf("planet:%v is registed to universe:%v\n", p.ID, universeID)
			u.Planets = append(u.Planets, p)
		}
	}
}

// RegisterUniverse add a new universe
func (m *Manager) RegisterUniverse(u *universe.Universe) {
	m.UniverseManager = append(m.UniverseManager, u)
}

// RegisterWorld add a new world
func (m *Manager) RegisterWorld(universeID, planetID int, w *world.World) {
	for _, u := range m.UniverseManager {
		if u.ID == universeID {
			for _, p := range u.Planets {
				if p.ID == planetID {
					fmt.Printf("world:%v is regested to planet:%v-universe:%v\n", w.ID, planetID, universeID)
					p.Worlds = append(p.Worlds, w)
				}
			}
		}
	}
}

// Run start up all components
func (m *Manager) Run() {
	for _, u := range m.UniverseManager {
		go u.Run()
		for _, pl := range u.Planets {
			go pl.Run()
			for _, w := range pl.Worlds {
				go w.Run()
			}
		}
	}
}

// GeneratePerson created a new person into a exsited area(world)
func (m *Manager) GeneratePerson() {
	if len(m.UniverseManager) == 0 {
		fmt.Println("there didn't had any exist universes")
		return
	}
	if len(m.UniverseManager[0].Planets) == 0 {
		fmt.Println("there didn't had any exist planets")
		return
	}
	if len(m.UniverseManager[0].Planets[0].Worlds) == 0 {
		fmt.Println("there didn't had any exist worlds")
		return
	}
	w := m.UniverseManager[0].Planets[0].Worlds[0]

	p := &person.Person{
		Identify: atomic.AddUint64(m.PersonID, 1),
		Channel:  make(chan []byte, 0),
	}
	w.Persons = append(w.Persons, p)
	fmt.Printf("person-identify:%v is created into universe-%v:planet-%v:world-%v\n", p.Identify, m.UniverseManager[0].ID, m.UniverseManager[0].Planets[0].ID, m.UniverseManager[0].Planets[0].Worlds[0].ID)
}
