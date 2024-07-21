package sylphy

type Config struct {
	Symbols
}

func DefaultRuntimeConfig() Config {
	return Config{
		Symbols: Symbols{
			ArrayOpen:          "[",
			ArrayClose:         "]",
			ArrayItemSeparator: ", ",
			MapOpen:            "[",
			MapClose:           "]",
			MapKVDivider:       ": ",
			MapItemSeparator:   ", ",
		},
	}
}

type Symbols struct {
	ArrayOpen          string
	ArrayClose         string
	ArrayItemSeparator string

	MapOpen          string
	MapClose         string
	MapKVDivider     string
	MapItemSeparator string
}
