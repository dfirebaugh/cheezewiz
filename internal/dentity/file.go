//go:build !js
// +build !js

package dentity

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

func pathToBytes(path string) []byte {
	f, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	return f
}

func parseJSON(path string) DynamicEntity {
	var d DynamicEntity
	json.Unmarshal(pathToBytes(path), &d)
	return d
}
