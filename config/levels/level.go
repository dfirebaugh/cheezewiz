package levels

import _ "embed"

//go:embed level1.json
var Level1Raw []byte

//go:embed stresstest.json
var StressTest []byte
