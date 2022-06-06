//go:build js
// +build js

package dentity

import (
	"encoding/json"
)

func pathToBytes(path string) []byte {
	return embededLookup[pathLabel(path)]
}

func parseJSON(path string) DynamicEntity {
	var d DynamicEntity
	json.Unmarshal(pathToBytes(path), &d)
	return d
}
