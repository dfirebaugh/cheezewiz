package system

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/event"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type eventQueue []func()

func (ej eventQueue) Consume() eventQueue {
	ej[0]()
	ej = append(ej[:0], ej[0+1:]...)

	return ej
}

type Scheduler struct {
	query  *query.Query
	events map[uint32]eventQueue
}

func NewScheduler(sceneEvents []event.SceneEvent, w donburi.World) *Scheduler {
	scheduler := &Scheduler{
		query:  query.NewQuery(filter.Contains(component.Countdown)),
		events: map[uint32]eventQueue{},
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

		if events, ok := s.events[countdown.CountDownInSec]; ok {
			s.RunEvents(w, events, countdown.CountDownInSec)
		}
	})
}

func (s Scheduler) RunEvents(w donburi.World, events eventQueue, now uint32) {
	if len(events) == 0 {
		return
	}

	s.events[now] = events.Consume()
	s.RunEvents(w, s.events[now], now)
}
