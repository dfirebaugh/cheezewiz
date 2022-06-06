package assets

import (
	"bytes"
	"cheezewiz/assets/scenes/garden"
	"cheezewiz/assets/scenes/kitchen"
	"encoding/xml"

	"github.com/lafriks/go-tiled"
)

func GetGardenMap() (*tiled.Map, error) {
	return tiled.LoadReader("./assets/scenes/garden", bytes.NewReader(garden.GardenRaw))
}

type fakeFS struct {
}

func (fs fakeFS) Open() {

}

func GetKitchenMap() (*tiled.Map, error) {
	m := &tiled.Map{}

	m.UnmarshalXML(xml.NewDecoder(bytes.NewReader(kitchen.Kitchen1Raw)), xml.StartElement{
		Name: xml.Name{Local: "Tile Layer 1"},
	})

	println(m.Height)

	return m, nil
	// return tiled.LoadReader("./assets/scenes/kitchen", bytes.NewReader(kitchen.Kitchen1Raw))
	// return tiled.LoadReader("./assets/scenes/kitchen", func(){

	// })
	// return tiled.LoadFile("./assets/scenes/kitchen/kitchen1.tmx")
}
