package sylphy

type Config struct {
	symbols
}

func DefaultRuntimeConfig() Config {
	return Config{
		symbols: symbols{
			arrayOpen:          "[",
			arrayClose:         "]",
			arrayItemSeparator: ", ",
			mapOpen:            "[",
			mapClose:           "]",
			mapKVDivider:       ": ",
			mapItemSeparator:   ", ",
		},
	}
}

type symbols struct {
	arrayOpen          string
	arrayClose         string
	arrayItemSeparator string

	mapOpen          string
	mapClose         string
	mapKVDivider     string
	mapItemSeparator string
}
