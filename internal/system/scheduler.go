package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/event"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Scheduler struct {
	query      *query.Query
	events     map[uint32][]func()
	weaponFire []func()
}

func NewScheduler(sceneEvents []event.SceneEvent, w donburi.World) *Scheduler {
	scheduler := &Scheduler{
		query:  query.NewQuery(filter.Contains(component.Countdown)),
		events: map[uint32][]func(){},
	}

	for _, se := range sceneEvents {
		scheduler.events[se.Time] = append(scheduler.events[se.Time], func() func() {
			world := w
			args := se.EventArgs
			name := se.EventName
			return func() {
				event.JobTypes[name].Callback(world, args)
			}
		}())
	}

	return scheduler
}

func (s Scheduler) Update(w donburi.World) {
	s.RunWeaponFire()

	s.query.EachEntity(w, func(entry *donburi.Entry) {
		countdown := component.GetCountdown(entry)

		s.RunEvents(countdown.CountDownInSec)
	})
}

func (s *Scheduler) RunWeaponFire() {
	for _, event := range s.weaponFire {
		event()
	}

	s.weaponFire = nil
}

func (s Scheduler) RunEvents(now uint32) {
	if _, ok := s.events[now]; !ok {
		return
	}

	for _, event := range s.events[now] {
		event()
	}

	s.events[now] = nil
}
