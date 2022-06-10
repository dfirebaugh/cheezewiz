package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Expbar struct {
}

func NewExpbar() *Expbar {
	return &Expbar{}
}

func (ne *Expbar) Update() {
	// ne.query.EachEntity(world, func(entry *donburi.Entry) {
	// 	exp := component.GetExp(entry)
	// 	nextLevel := component.GetNextLevel(entry)

	// 	if exp.CurrentExp >= exp.DesiredExp {

	// 		ne.jellyBeanQuery.EachEntity(world, func(entry *donburi.Entry) {
	// 			xp := (*component.XPData)(component.GetXP(entry))
	// 			xp.Value = xp.Value / 2
	// 		})

	// 		exp.CurrentExp = 0
	// 		nextLevel.CurrentLevel += 1
	// 	}
	// })
}

func (ne *Expbar) Draw(screen *ebiten.Image) {
	// ne.query.EachEntity(world, func(entry *donburi.Entry) {
	// 	exp := component.GetExp(entry)
	// 	position := component.GetPosition(entry)
	// 	nextLevel := component.GetNextLevel(entry)

	// 	ebitenutil.DrawRect(screen, position.X, position.Y, exp.DesiredExp*100, 16, colornames.Black)
	// 	ebitenutil.DrawRect(screen, position.X, position.Y, exp.CurrentExp*100, 16, colornames.Blue700)
	// 	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Lvl %d", nextLevel.CurrentLevel), (int(position.X)+int(exp.DesiredExp)*100)-50, int(position.Y))
	// })
}
