package config

type Event struct {
	Name    string         `yaml:"name"`
	Matcher map[string]any `yaml:"matcher"`
	Scheme  []SchemeItem   `yaml:"scheme"`
}

func (evt *Event) Match(entry map[string]any) bool {
	for key, val := range evt.Matcher {
		param, ok := entry[key]
		if !ok {
			return false
		}

		if param != val {
			return false
		}
	}

	return true
}
