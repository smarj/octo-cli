package linter

const (
	PresetFormatting  = "format"
	PresetComplexity  = "complexity"
	PresetStyle       = "style"
	PresetBugs        = "bugs"
	PresetUnused      = "unused"
	PresetPerformance = "performance"
)

type Config struct {
	Linter           Linter
	EnabledByDefault bool

	NeedsTypeInfo bool
	NeedsSSARepr  bool

	InPresets        []string
	Speed            int // more value means faster execution of linter
	AlternativeNames []string

	OriginalURL string // URL of original (not forked) repo, needed for autogenerated README
}

func (lc Config) WithTypeInfo() Config {
	lc.NeedsTypeInfo = true
	return lc
}

func (lc Config) WithSSA() Config {
	lc.NeedsTypeInfo = true
	lc.NeedsSSARepr = true
	return lc
}

func (lc Config) WithPresets(presets ...string) Config {
	lc.InPresets = presets
	return lc
}

func (lc Config) WithSpeed(speed int) Config {
	lc.Speed = speed
	return lc
}

func (lc Config) WithURL(url string) Config {
	lc.OriginalURL = url
	return lc
}

func (lc Config) WithAlternativeNames(names ...string) Config {
	lc.AlternativeNames = names
	return lc
}

func (lc Config) GetSpeed() int {
	return lc.Speed
}

func (lc Config) AllNames() []string {
	return append([]string{lc.Name()}, lc.AlternativeNames...)
}

func (lc Config) Name() string {
	return lc.Linter.Name()
}

func NewConfig(linter Linter) *Config {
	return &Config{
		Linter: linter,
	}
}
