package component

import (
	"cheezewiz/internal/input"

	"github.com/yohamta/donburi"
)

type InputDeviceData struct {
	Device input.PlayerInput `json:"device"`
}

var InputDevice = donburi.NewComponentType(InputDeviceData{})

func GetInputDevice(entry *donburi.Entry) *InputDeviceData {
	return (*InputDeviceData)(entry.Component(InputDevice))
}
