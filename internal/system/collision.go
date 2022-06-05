package system

import (
	"cheezewiz/internal/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type attackMediator interface {
	AddPlayerDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry)
	AddEnemyDamage(destination *donburi.Entry, amount float64, origin *donburi.Entry)
}

type Collision struct {
	query *query.Query
}

func NewCollision(attack_handler attackMediator) *Collision {
	return &Collision{
		query: query.NewQuery(filter.Contains(component.RigidBody)),
	}
}

func (c *Collision) Update(w donburi.World) {
	c.query.EachEntity(w, func(entry *donburi.Entry) {
		id := entry.Id()
		rb := component.GetRigidBody(entry)
		p := component.GetPosition(entry)

		ax := p.X - rb.L
		ay := p.Y - rb.T
		aw := rb.GetWidth()
		ah := rb.GetHeight()
		c.query.EachEntity(w, func(e *donburi.Entry) {
			if id == e.Id() {
				return
			}
			p2 := component.GetPosition(e)
			erb := component.GetRigidBody(e)
			bx := p2.X - erb.L
			by := p2.Y - erb.T
			bw := erb.GetWidth()
			bh := erb.GetHeight()

			if c.IsCollide([]float64{ax, ay, aw, ah}, []float64{bx, by, bw, bh}) {
				if rb.CollisionHandler == nil {
					return
				}
				rb.CollisionHandler(w, e)
			}
		})
	})
}

func (c *Collision) Draw(w donburi.World, screen *ebiten.Image) {
}

func (c *Collision) IsCollide(a []float64, b []float64) bool {
	ax := a[0]
	ay := a[1]
	aw := a[2]
	ah := a[3]

	bx := b[0]
	by := b[1]
	bw := b[2]
	bh := b[3]

	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ah+ay > by
}
