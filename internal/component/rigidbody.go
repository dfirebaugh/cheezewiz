package component

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
)

type CollisionHandlerLabel string

const (
	PlayerCollisionLabel    CollisionHandlerLabel = "player"
	EnemyCollisionLabel     CollisionHandlerLabel = "enemy"
	BossCollisionLabel      CollisionHandlerLabel = "boss"
	RocketCollisionLabel    CollisionHandlerLabel = "rocket"
	JellyBeanCollisionLabel CollisionHandlerLabel = "jellybean"
)

type attackGroup interface {
	AddPlayerDamage(reciever *donburi.Entry, amount float64, origin *donburi.Entry)
	AddEnemyDamage(reciever *donburi.Entry, amount float64, Origin *donburi.Entry)
	ApplyDamageToEnemy(w donburi.World, enemy *donburi.Entry, amount float64)
}

type RigidBodyData struct {
	R                     float64 `json:"r"`
	L                     float64 `json:"l"`
	T                     float64 `json:"t"`
	B                     float64 `json:"b"`
	CollisionHandler      func(w donburi.World, e *donburi.Entry)
	CollisionHandlerLabel string `json:"collisionLabel"`
	am                    attackGroup
}

var RigidBody = donburi.NewComponentType(RigidBodyData{})

func GetRigidBody(entry *donburi.Entry) *RigidBodyData {
	return (*RigidBodyData)(entry.Component(RigidBody))
}

func (r RigidBodyData) GetHeight() float64 {
	return r.R + r.L
}
func (r RigidBodyData) GetWidth() float64 {
	return r.T + r.B
}

func (r *RigidBodyData) SetBorder(wx float64, wy float64) {
	r.R = wx / 2
	r.L = wx / 2
	r.T = wy / 2
	r.B = wy / 2
}

func (r *RigidBodyData) SetCollisionHandler(label interface{}) {
	var l CollisionHandlerLabel
	var ok bool

	logrus.Info(string(label.(CollisionHandlerLabel)))
	if l, ok = label.(CollisionHandlerLabel); !ok || l == "" {
		logrus.Error("not a defined collision handler")
		return
	}

	handler, err := r.GetCollisionHandler(l)
	if err != nil {
		logrus.Errorf("not able to set collision handler %s -- %s", label, err)
		r.CollisionHandler = func(w donburi.World, e *donburi.Entry) {}
		return
	}

	r.CollisionHandler = handler
}

func (r RigidBodyData) GetCollisionHandler(label CollisionHandlerLabel) (func(w donburi.World, e *donburi.Entry), error) {
	c := map[CollisionHandlerLabel]func(w donburi.World, e *donburi.Entry){
		EnemyCollisionLabel: func(w donburi.World, e *donburi.Entry) {
			if !w.Valid(e.Entity()) {
				return
			}
			if e.Archetype().Layout().HasComponent(PlayerTag) {
				state := GetActorState(e)
				state.Set(HurtState)

				// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
				// r.am.AddPlayerDamage(e, 10, nil)
			}
			if e.Archetype().Layout().HasComponent(ProjectileTag) {
				logrus.Info("enemy collided with projectile")
				// w.Remove(e.Entity())
			}
		},
		RocketCollisionLabel: func(w donburi.World, e *donburi.Entry) {
			if !w.Valid(e.Entity()) {
				return
			}

			if e.Archetype().Layout().HasComponent(EnemyTag) {
				logrus.Info("missile collided with enemy")
				// w.Remove(e.Entity())
				// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
				// r.am.AddEnemyDamage(e, 10, nil)
			}
		},
		BossCollisionLabel: func(w donburi.World, e *donburi.Entry) {
			if !w.Valid(e.Entity()) {
				return
			}
			if e.Archetype().Layout().HasComponent(PlayerTag) {
				logrus.Info("collision with boss")
				// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
				// r.am.AddEnemyDamage(e, 10, nil)
			}

		},
		PlayerCollisionLabel: func(w donburi.World, e *donburi.Entry) {
			if !w.Valid(e.Entity()) {
				return
			}
			if e.Archetype().Layout().HasComponent(EnemyTag) {
				logrus.Info("player collided with enemy")
				// MakeDamageLabel(w, position.X, position.Y, strconv.Itoa(10))
				// r.am.AddEnemyDamage(e, 10, nil)
			}
			if e.Archetype().Layout().HasComponent(JellyBeanTag) {
				logrus.Info("player collided with enemy")
				w.Remove(e.Entity())
			}
		},
	}

	if _, ok := c[label]; !ok {
		return nil, fmt.Errorf("could not find collision handler: %s", label)
	}

	return c[label], nil
}
