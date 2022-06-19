package system

import (
	"cheezewiz/internal/event"
	"cheezewiz/internal/world"
)

type Scheduler struct {
	events    map[uint32][]func()
	countdown struct {
		seconds uint32
	}
}

func NewScheduler(sceneEvents []event.SceneEvent, w world.World) *Scheduler {
	scheduler := &Scheduler{
		events: map[uint32][]func(){},
		countdown: struct {
			seconds uint32
		}{
			seconds: 120,
		},
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

var tpsCount = 0

func (s *Scheduler) Update() {
	if s.countdown.seconds == 0 {
		return
	}
	s.RunEvents()

	tpsCount++
	if tpsCount%60 == 0 && s.countdown.seconds > 0 {
		s.countdown.seconds--
	}
}

func (s *Scheduler) RunEvents() {
	now := s.countdown.seconds
	if _, ok := s.events[now]; !ok {
		return
	}

	for _, event := range s.events[now] {
		event()
	}

	s.events[now] = nil
}
