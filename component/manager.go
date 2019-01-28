package component

import (
	"Martial-arts/component/planet"
	"Martial-arts/component/universe"
	"Martial-arts/component/world"
	"sync"
)

// Manager Global manager
type Manager struct {
	Wg sync.WaitGroup
	// PlanetManager   []*planet.Planet
	UniverseManager []*universe.Universe
	// WorldManager    []*world.World
}

// RegestPlanet add a new planet
func (m *Manager) RegestPlanet(p *planet.Planet) {
	m.UniverseManager[0].Planets = append(m.UniverseManager[0].Planets, p)
}

// RegestUniverse add a new universe
func (m *Manager) RegestUniverse(u *universe.Universe) {
	m.UniverseManager = append(m.UniverseManager, u)
}

// RegestWorld add a new world
func (m *Manager) RegestWorld(w *world.World) {
	m.UniverseManager[0].Planets[0].Worlds = append(m.UniverseManager[0].Planets[0].Worlds, w)
}

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
	// for _, pl := range m.PlanetManager {
	// 	go pl.Run()
	// }
	// for _, w := range m.WorldManager {
	// 	go w.Run()
	// }

}
