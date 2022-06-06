//go:build js
// +build js

package dentity

func pathToBytes(path string) []byte {
	return embededLookup[pathLabel(path)]
}
