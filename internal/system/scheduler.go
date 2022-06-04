package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/event"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Scheduler struct {
	query  *query.Query
	events map[uint32][]func()
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
			return func() {
				event.JobTypes[se.EventName].Callback(world, args)
			}
		}())
	}

	return scheduler
}

func (s Scheduler) Update(w donburi.World) {
	s.query.EachEntity(w, func(entry *donburi.Entry) {
		countdown := component.GetCountdown(entry)

		s.RunEvents(w, countdown.CountDownInSec)
	})
}

func (s Scheduler) RunEvents(w donburi.World, now uint32) {
	if _, ok := s.events[now]; !ok {
		return
	}

	for _, event := range s.events[now] {
		event()
	}

	s.events[now] = nil
}
