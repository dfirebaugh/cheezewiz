package system

import (
	"cheezewiz/internal/entity"
	"cheezewiz/internal/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

type RegisterPlayer struct {
	gamepadIDsBuf []ebiten.GamepadID
	gamepadIDs    map[ebiten.GamepadID]*donburi.Entry
}

func NewRegisterPlayer() *RegisterPlayer {
	return &RegisterPlayer{}
}

func (rp *RegisterPlayer) Update(world donburi.World) {
	if rp.gamepadIDs == nil {
		rp.gamepadIDs = make(map[ebiten.GamepadID]*donburi.Entry)
	}

	// Log the gamepad connection events.
	rp.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(rp.gamepadIDsBuf[:0])
	for _, id := range rp.gamepadIDsBuf {
		entry := entity.MakePlayer(world, input.GamePad{ID: id})
		rp.gamepadIDs[id] = entry
	}

	for id, entry := range rp.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(rp.gamepadIDs, id)
			world.Remove(entry.Entity())
		}
	}
}
