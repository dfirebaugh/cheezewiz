package main

import (
	"cheezewiz/config"
	"cheezewiz/internal/component"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/tag"
	"cheezewiz/internal/world"
	"cheezewiz/internal/world/query"
	"cheezewiz/internal/world/query/filter"
	"cheezewiz/pkg/gamemath"
	"cheezewiz/pkg/throttle"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var rectConfig = entity.EntityConfig{
	Tags: []string{"animatable", "collidable"},
	RigidBody: component.RigidBody{
		R: 8,
		L: 8,
		T: 8,
		B: 8,
	},
}

type mover struct{}

func (m mover) Update() {
	world.Instance.EachEntity(func(handle world.EntityHandle) {
		entity := world.Instance.GetEntity(handle)

		if !entity.HasTag(tag.Collidable) {
			return
		}
		if entity.HasTag(tag.Bound) {
			return
		}
		if entity.HasTag(tag.Player) {
			return
		}

		position := entity.GetPosition()
		dir := entity.GetDirection()
		position.X += math.Cos(dir.Angle)
		position.Y += math.Sin(dir.Angle)
	})
}

type collision struct{}

func (c collision) getCollidedEntities(origin world.EntityHandle) {
	e := world.Instance.GetEntity(origin)
	erb := e.GetRigidBody()
	oPosition := e.GetPosition()
	oRect := gamemath.Rect([]float64{oPosition.X, oPosition.Y, erb.L + erb.R, erb.T + erb.B})

	world.Instance.EachEntity(func(handle world.EntityHandle) {
		if origin == handle {
			return
		}
		entity := world.Instance.GetEntity(handle)
		position := entity.GetPosition()
		rb := entity.GetRigidBody()

		if !oRect.IsAxisAlignedCollision(gamemath.Rect([]float64{position.X, position.Y, rb.L + rb.R, rb.T + rb.B})) {
			return
		}
		c.resolveCollision(origin, handle)
	})
}

func (c collision) isOutOfBounds(position *component.Position) bool {
	return position.X > float64(config.Get().Window.Width/config.Get().ScaleFactor) || position.X < 0 || position.Y > float64(config.Get().Window.Height/config.Get().ScaleFactor) || position.Y < 0
}
func (c collision) isWithinBoundLimit(position *component.Position) bool {
	limit := 15.0
	w := float64(config.Get().Window.Width / config.Get().ScaleFactor)
	h := float64(config.Get().Window.Height / config.Get().ScaleFactor)
	return position.X+limit > w || position.X-limit < 0 || position.Y+limit > h || position.Y-limit < 0
}

// check if the entity has collided with the wall
func (c collision) checkBounds(handle world.EntityHandle) bool {
	e := world.Instance.GetEntity(handle)
	position := e.GetPosition()
	dir := e.GetDirection()
	if c.isOutOfBounds(position) {
		world.Instance.Remove(handle)
		return true
	}
	if c.isWithinBoundLimit(position) {
		center := gamemath.MakeVector(float64(config.Get().Window.Width/config.Get().ScaleFactor)/2, float64(config.Get().Window.Height/config.Get().ScaleFactor)/2)
		pVec := gamemath.MakeVector(position.X, position.Y)

		// move toward center
		dir.Angle = center.GetHeading(pVec)
		return true
	}

	return false
}

func (c collision) Update() {
	if throttle.ShouldThrottle("collision", 3) {
		return
	}
	world.Instance.EachEntity(func(handle world.EntityHandle) {
		e := world.Instance.GetEntity(handle)
		if !e.HasTag(tag.Collidable) {
			return
		}
		if c.checkBounds(handle) {
			return
		}
		c.getCollidedEntities(handle)
	})
}

// handler for when two entities collide
func (c collision) resolveCollision(h world.EntityHandle, t world.EntityHandle) {
	a := world.Instance.GetEntity(h)
	b := world.Instance.GetEntity(t)

	aPosition := a.GetPosition()
	bPosition := b.GetPosition()

	av := gamemath.MakeVector(aPosition.X, aPosition.Y)
	bv := gamemath.MakeVector(bPosition.X, bPosition.Y)

	a.GetDirection().Angle = av.GetHeading(bv)
	b.GetDirection().Angle = bv.GetHeading(av)
}

type renderer struct{}

func (r renderer) Draw(screen *ebiten.Image) {
	query.Each(world.Instance, filter.GetAnimatables, func(handle world.EntityHandle) {
		e := world.Instance.GetEntity(handle)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		e.Draw(screen, op)
	})
}

type clickSpawner struct{}

func (s clickSpawner) generate(n int, clickX int, clickY int) {

	radians_spread := (2.0 * math.Pi) / float64(6)
	distance := 44

	for i := 0; i < n; i++ {
		x := math.Cos(radians_spread * float64(i))
		y := math.Sin(radians_spread * float64(i))
		clickLocation := gamemath.MakeVector(float64(clickX), float64(clickY))
		spawnLocation := gamemath.MakeVector(x, y)

		spawnLocation.Scale(float64(distance))
		spawnLocation = spawnLocation.Add(clickLocation)

		e := entity.BuildEntity(
			rectConfig,
			spawnLocation[0],
			spawnLocation[1],
		)

		// move away from click location
		e.GetDirection().Angle = spawnLocation.GetHeading(clickLocation)
		world.Instance.Add(e)
	}
}
func (s clickSpawner) remove(x int, y int) {
	query.Each(world.Instance, filter.GetCollidables, func(handle world.EntityHandle) {
		e := world.Instance.GetEntity(handle)
		if e.HasTag(tag.Player) {
			return
		}

		position := e.GetPosition()
		clickVector := gamemath.MakeVector(float64(x), float64(y))
		v := gamemath.MakeVector(position.X, position.Y)
		if v.GetDistance(clickVector) > 55 {
			return
		}
		world.Instance.Remove(handle)
	})
}
func (s clickSpawner) Update() {
	x, y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.generate(6, x, y)
		return
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.remove(x, y)
		return
	}
}

type hud struct{}

func (h hud) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "click the screen", 0, config.Get().Window.Height/config.Get().ScaleFactor-20)
	ebitenutil.DebugPrintAt(screen, "right click to remove", config.Get().Window.Width/config.Get().ScaleFactor-130, config.Get().Window.Height/config.Get().ScaleFactor-20)
}
