package component

import (
	"time"

	"github.com/yohamta/donburi"
)

type LabelData struct {
	Label        string
	CreationTime time.Time
}

var ScreenLabel = donburi.NewComponentType(LabelData{})

func GetScreenLabel(e *donburi.Entry) *LabelData {
	return (*LabelData)(e.Component(ScreenLabel))
}
