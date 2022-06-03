package system

import (
	"cheezewiz/examples/choppa/internal/entity"

	"github.com/yohamta/donburi"
)

type Spawner struct {
	tick uint
}

func NewSpawner() *Spawner {
	return &Spawner{
		tick: 1,
	}
}

func (s *Spawner) Update(w donburi.World) {
	if s.tick%100 == 0 {
		entity.MakeFighter(w)
	}
	if s.tick%150 == 0 {
		entity.MakeChippa(w)
	}

	s.tick++
}
