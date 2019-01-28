package global

import (
	"Martial-arts/component"
	"Martial-arts/component/planet"
	"Martial-arts/component/universe"
	"Martial-arts/component/world"
)

// Manager Global manager factory
var Manager *component.Manager

func init() {
	Manager = &component.Manager{
		PlanetManager:   make([]*planet.Planet, 0),
		UniverseManager: make([]*universe.Universe, 0),
		WorldManager:    make([]*world.World, 0),
	}
}
