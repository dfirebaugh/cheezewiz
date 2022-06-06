package dentity

import (
	"bytes"
	"cheezewiz/internal/component"
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type componentLabel string

// the elements in the components array should match these labels
const (
	XP           componentLabel = "XP"
	SpriteSheet  componentLabel = "SpriteSheet"
	Position     componentLabel = "Position"
	RigidBody    componentLabel = "RigidBody"
	JellyBeanTag componentLabel = "JellyBeanTag"
)

var componentTable = map[componentLabel]*donburi.ComponentType{
	XP:           component.XP,
	SpriteSheet:  component.SpriteSheet,
	Position:     component.Position,
	RigidBody:    component.RigidBody,
	JellyBeanTag: component.JellyBeanTag,
}

// DynamicEntity is an entity that can be configured at runtime by parsing a json file
//  the structue of the json file will have to marshal out correctly
type DynamicEntity struct {
	// elements in the Components array should match strings in the `componentLabel` enum
	Components   []componentLabel           `json:"components"`
	XP           *component.XPData          `json:"xp"`
	SpriteSheet  *component.SpriteSheetData `json:"spriteSheet"`
	Position     *component.PositionData    `json:"position"`
	RigidBody    *component.RigidBodyData   `json:"rigidBody"`
	JellyBeanTag *donburi.ComponentType     `json:"jellyBeanTag"`
}

func MakeRandDynamicEntity(w donburi.World, path []string, x float64, y float64) {
	d := parseJSON(path[rand.Intn(len(path))])
	entry := w.Entry(w.Create(getComponents(d)...))
	d.initializeValues(entry, x, y)
}

func MakeDynamicEntity(w donburi.World, path string, x float64, y float64) {
	d := parseJSON(path)
	entry := w.Entry(w.Create(getComponents(d)...))

	d.initializeValues(entry, x, y)
}

func getComponents(d DynamicEntity) []*donburi.ComponentType {
	var ct []*donburi.ComponentType
	for _, key := range d.Components {
		if componentTable[key] == nil {
			println("component was not added: ", key, componentTable[key], componentTable[key] == nil)
			continue
		}
		ct = append(ct, componentTable[key])
	}
	return ct
}

func (d DynamicEntity) initializeValues(entry *donburi.Entry, x, y float64) {
	d.setPosition(entry, x, y)
	d.setXP(entry)
	d.setSpriteSheet(entry)
	d.setRigidBody(entry)
}
func (d DynamicEntity) setXP(entry *donburi.Entry) {
	if d.XP != nil {
		xp := component.GetXP(entry)
		xp.Value = d.XP.Value
	}
}
func (d DynamicEntity) setSpriteSheet(entry *donburi.Entry) {
	if d.SpriteSheet != nil {
		spriteSheet := component.GetSpriteSheet(entry)
		spriteSheet.Path = d.SpriteSheet.Path

		img, _, _ := image.Decode(bytes.NewReader(pathToBytes(d.SpriteSheet.Path)))
		spriteSheet.IMG = ebiten.NewImageFromImage(img)
	}
}
func (d DynamicEntity) setPosition(entry *donburi.Entry, x, y float64) {
	if d.Position != nil {
		position := component.GetPosition(entry)
		position.Set(x, y, d.Position.CX, d.Position.CY)
	}
}
func (d DynamicEntity) setRigidBody(entry *donburi.Entry) {
	if d.RigidBody != nil {
		rb := component.GetRigidBody(entry)
		rb.SetBorder(d.RigidBody.R, d.RigidBody.B)
	}
}
