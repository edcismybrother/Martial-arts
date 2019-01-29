package global

import (
	"Martial-arts/component"
	"Martial-arts/component/person"
	"Martial-arts/component/planet"
	"Martial-arts/component/universe"
	"Martial-arts/component/world"
)

// Manager Global manager factory
// var Manager *component.Manager

// func init() {
// 	Manager = &component.Manager{
// 		PlanetManager:   make([]*planet.Planet, 0),
// 		UniverseManager: make([]*universe.Universe, 0),
// 		WorldManager:    make([]*world.World, 0),
// 	}
// }

// NewManager 返回新的管理
func NewManager() *component.Manager {
	return &component.Manager{
		// PlanetManager:   make([]*planet.Planet, 0),
		UniverseManager: make([]*universe.Universe, 0),
		// WorldManager:    make([]*world.World, 0),
		PersonID: new(uint64),
	}
}

func NewComplateManager() *component.Manager {
	m := &component.Manager{
		// PlanetManager:   make([]*planet.Planet, 0),
		UniverseManager: make([]*universe.Universe, 0),
		// WorldManager:    make([]*world.World, 0),
		PersonID: new(uint64),
	}

	// m.RegisterUniverse(&universe.Universe{
	// 	Planets: []*planet.Planet{&planet.Planet{
	// 		Worlds: []*world.World{&world.World{
	// 			Persons: make([]*person.Person, 0),
	// 			Channel: make(chan []byte, 100),
	// 		}},
	// 		Channel: make(chan []byte, 100),
	// 	}},
	// 	Channel: make(chan []byte, 100)})
	u := &universe.Universe{
		ID:      1,
		Planets: make([]*planet.Planet, 0),
		Channel: make(chan []byte, 100),
	}
	m.RegisterUniverse(u)
	p := &planet.Planet{
		ID:      1,
		Worlds:  make([]*world.World, 0),
		Channel: make(chan []byte, 100),
	}
	m.RegisterPlanet(u.ID, p)
	w := &world.World{
		ID:      1,
		Persons: make([]*person.Person, 0),
		Channel: make(chan []byte, 100),
	}
	m.RegisterWorld(u.ID, p.ID, w)

	// m.RegestPlanet(&planet.Planet{
	// 	Worlds: []*world.World{&world.World{
	// 		Persons: make([]*person.Person, 0),
	// 		Channel: make(chan []byte, 100),
	// 	}},
	// 	Channel: make(chan []byte, 100),
	// })
	// m.RegestWorld(
	// 	&world.World{
	// 		Persons: make([]*person.Person, 0),
	// 		Channel: make(chan []byte, 100),
	// 	},
	// )
	return m
}
