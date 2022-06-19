package entity

import (
	"cheezewiz/internal/component"
	"cheezewiz/internal/input"
	"cheezewiz/internal/tag"
	"cheezewiz/pkg/animation"
	"cheezewiz/pkg/gamemath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Entity interface {
	GetHealth() *component.Health
	GetState() *component.State
	GetDirection() *component.Direction
	GetPosition() *component.Position
	GetRigidBody() *component.RigidBody
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions)
	DebugDraw(screen *ebiten.Image)
	IterFrame()
	AddTag(t tag.Tag)
	HasTag(t tag.Tag) bool
	GetInputDevice() input.PlayerInput
}

type entity struct {
	*component.Animation
	*component.State
	*component.Position
	*component.Health
	*component.RigidBody
	*component.InputDevice
	*component.Direction
	*component.TagSet
}

func (e entity) GetRigidBody() *component.RigidBody {
	return e.RigidBody
}
func (e entity) GetFrame() *ebiten.Image {
	current := e.State.GetCurrent()

	return e.Animation.Animation[current].GetFrame()
}
func (e entity) GetPosition() *component.Position {
	return e.Position
}
func (e entity) GetCurrentState() component.StateType {
	return e.State.GetCurrent()
}
func (e entity) GetCurrent() *animation.Animation {
	return e.Animation.Animation[e.GetCurrentState()]
}
func (e entity) IterFrame() {
	e.GetCurrent().IterFrame()
}
func (e entity) GetHealth() *component.Health {
	return e.Health
}
func (e entity) GetState() *component.State {
	return e.State
}
func (e entity) GetDirection() *component.Direction {
	return e.Direction
}
func (e entity) AddTag(t tag.Tag) {
	e.TagSet.Add(t)
}
func (e entity) HasTag(t tag.Tag) bool {
	return e.TagSet.Contains(t)
}
func (e entity) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	position := e.GetPosition()
	offset := gamemath.MakeVector(position.X, position.Y).Offset(gamemath.MakeVector(position.CX, position.CY))
	// worldCoord := offset.Offset()
	// wX, wY := r.getWorldCoord(position)

	op.GeoM.Translate(offset[0], offset[1])
	screen.DrawImage(e.GetFrame(), op)

	if !e.HasTag(tag.Player) {
		return
	}
	e.drawHealthBar(screen)
}
func (e entity) drawHealthBar(screen *ebiten.Image) {
	position := e.GetPosition()
	health := e.GetHealth()
	// x, y := r.getWorldCoord(position)
	offset := gamemath.MakeVector(position.X, position.Y).Offset(gamemath.MakeVector(position.CX, position.CY))

	x := offset[0]
	y := offset[1]

	var marginBottom float64 = 35

	ebitenutil.DrawRect(screen, float64(x), marginBottom+float64(y), health.Max/3, 3, colornames.Grey)
	ebitenutil.DrawRect(screen, float64(x), marginBottom+float64(y), health.Current/3, 3, colornames.Red)
}
func (e entity) DebugDraw(screen *ebiten.Image) {
	p := e.GetPosition()
	rb := e.GetRigidBody()
	ebitenutil.DrawRect(screen, p.X, p.Y, rb.GetWidth(), rb.GetHeight(), colornames.Aliceblue)
}
func (e entity) GetInputDevice() input.PlayerInput {
	return e.InputDevice.Device
}
