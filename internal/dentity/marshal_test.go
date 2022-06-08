package dentity

import (
	"testing"
)

func TestUnMarshalConfig(t *testing.T) {
	tests := []struct {
		message  string
		expected bool
		eval     func(EntityConfig) bool
		config   string
	}{
		{
			message:  "empty test should pass",
			expected: true,
			eval: func(EntityConfig) bool {
				return true
			},
			config: ``,
		},
		{
			message:  "should get list of keys as component labels",
			expected: true,
			eval: func(e EntityConfig) bool {
				for _, c := range e.Components {
					println(c)
				}
				return len(e.Components) == 5
			},
			config: `{
				"jellyBeanTag": "",
				"xp": 1,
				"spriteSheet": "./assets/jellybeanblue.png",
				"position": {
					"x": 0,
					"y": 0,
					"cx": 16,
					"cy": 16
				},
				"rigidBody": {
					"l": 16,
					"r": 16,
					"t": 16,
					"b": 16
				}
			}`,
		},
	}

	for _, test := range tests {
		e := EntityConfig{}

		// json.Unmarshal([]byte(test.config), &e)
		// e.Components = getComponentLabels([]byte(test.config))
		e.Unmarshal([]byte(test.config))

		if test.eval(e) != test.expected {
			t.Errorf(test.message)
		}
	}
}

func TestGetComponentLabels(t *testing.T) {
	labels := getComponentLabels([]byte(`{"xp": "", "spriteSheet": "", "position": ""}`))
	if len(labels) != 3 {
		t.Errorf("should get 3 labels")
	}
}
