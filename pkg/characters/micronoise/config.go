package micronoise

// FlickerConfig defines color flicker behavior for micro avatars
type FlickerConfig struct {
	Enabled bool // Whether flicker is active for this state
}

// StateConfigs maps state names to their flicker configuration.
// States not in this map will have no flicker effect.
var StateConfigs = map[string]FlickerConfig{
	"think":     {Enabled: true},
	"read":      {Enabled: true},
	"write":     {Enabled: true},
	"search":    {Enabled: true},
	"websearch": {Enabled: true},
	"webfetch":  {Enabled: true},
	"arise":     {Enabled: true},
	"approval":  {Enabled: true},
}

// BrightnessLevels defines the 8 brightness multipliers for the gradient.
// Goes from dark (left) to bright (right), creating a "marquee" effect.
// Values < 1.0 = darker, values > 1.0 = brighter
var BrightnessLevels = []float64{
	0.25, // Column 0 - darkest
	0.40, // Column 1
	0.55, // Column 2
	0.70, // Column 3
	0.85, // Column 4
	1.00, // Column 5
	1.15, // Column 6
	1.30, // Column 7 - brightest
}

// GetConfig returns the flicker config for a state, or nil if no flicker effect.
func GetConfig(stateName string) *FlickerConfig {
	if cfg, ok := StateConfigs[stateName]; ok {
		return &cfg
	}
	return nil
}
