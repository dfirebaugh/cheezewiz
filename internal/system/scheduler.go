package system

import (
	"cheezewiz/internal/event"
	"cheezewiz/pkg/ecs"
)

type Scheduler struct {
	events     map[uint32][]func()
	weaponFire []func()
	countdown  struct {
		seconds uint32
	}
}

func NewScheduler(sceneEvents []event.SceneEvent, w ecs.World) *Scheduler {
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

func (s *Scheduler) Update() {
	if s.countdown.seconds == 0 {
		return
	}
	// s.world.FilterBy()
	// s.RunWeaponFire()
	s.RunEvents()
	println("second:", s.countdown.seconds)

	tpsCount++
	if tpsCount%60 == 0 && s.countdown.seconds > 0 {
		s.countdown.seconds--
	}
	// s.query.EachEntity(s.world, func(entry *donburi.Entry) {
	// 	countdown := component.GetCountdown(entry)

	// 	s.RunEvents(countdown.CountDownInSec)
	// })
}

func (s *Scheduler) RunWeaponFire() {
	for _, event := range s.weaponFire {
		event()
	}

	s.weaponFire = nil
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
