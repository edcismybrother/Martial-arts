package component

import (
	"Martial-arts/component/planet"
	"Martial-arts/component/universe"
	"Martial-arts/component/world"
)

// Manager Global manager
type Manager struct {
	PlanetManager   []*planet.Planet
	UniverseManager []*universe.Universe
	WorldManager    []*world.World
}

// RegestPlanet add a new planet
func (m *Manager) RegestPlanet(p *planet.Planet) {
	m.PlanetManager = append(m.PlanetManager, p)
}

// RegestUniverse add a new universe
func (m *Manager) RegestUniverse(u *universe.Universe) {
	m.UniverseManager = append(m.UniverseManager, u)
}

// RegestWorld add a new world
func (m *Manager) RegestWorld(w *world.World) {
	m.WorldManager = append(m.WorldManager, w)
}
