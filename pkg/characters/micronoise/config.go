package micronoise

// FlickerConfig defines random color flicker behavior for micro avatars
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

// ThemePalette holds the 7 theme colors for flicker effect (Cozy theme)
// Each character in the avatar randomly picks from this palette each frame
var ThemePalette = []string{
	"#E18B8B", // Sa - Rose quartz
	"#E5A679", // Ri - Amber glow
	"#E6CC94", // Ga - Honey
	"#99C794", // Ma - Balanced green
	"#78AED4", // Pa - Clear sky
	"#B592D4", // Da - Soft violet
	"#DE99B8", // Ni - Blush
}

// GetConfig returns the flicker config for a state, or nil if no flicker effect.
func GetConfig(stateName string) *FlickerConfig {
	if cfg, ok := StateConfigs[stateName]; ok {
		return &cfg
	}
	return nil
}
