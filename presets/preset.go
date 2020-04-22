package presets

type Preset = map[string][]byte

var Presets = map[string]Preset{}

func Register(name string, s Preset) {
	Presets[name] = s
}
