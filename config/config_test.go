package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var numTypeConfig = `
    types:
      - name: sample
        type: number
        display:
          fg: red
    `

var invalidConfig = `
    addafgs
    `

func TestLoadConfig(t *testing.T) {
	a := assert.New(t)
	validInput := strings.NewReader(numTypeConfig)

	cfg, err := LoadConfig(validInput)
	a.NoError(err)
	a.IsType(cfg, &Config{})

	invalidInput := strings.NewReader(invalidConfig)
	cfg, err = LoadConfig(invalidInput)
	a.Error(err)
	a.Nil(cfg)
}

func TestTypesParsing(t *testing.T) {
	a := assert.New(t)
	validInput := strings.NewReader(numTypeConfig)

	cfg, err := LoadConfig(validInput)
	a.NoError(err)

	a.Equal(1, len(cfg.Types), "len(Config.Types)")

	typ := cfg.Types[0]
	a.Equal("sample", typ.Name, "Type.Name")
	// it assigns DisplayProps
	a.Equal("red", typ.Fg, "Type.Fg")
}

var eventsConfig = `
  events:
    - name: ReqLog
      matcher:
        foo: bar
      scheme:
        - name: count
          type: number
  `

func TestEventsParsing(t *testing.T) {
	a := assert.New(t)
	validInput := strings.NewReader(eventsConfig)
	cfg, err := LoadConfig(validInput)
	a.NoError(err)

	a.Equal(1, len(cfg.Events), "len(Config.Events)")

	evt := cfg.Events[0]
	a.Equal("ReqLog", evt.Name, "Event.Name")

	expectedMatcher := map[string]string{
		"foo": "bar",
	}
	a.Equal(expectedMatcher, evt.Matcher, "Event.Matcher")

	scheme := evt.Scheme
	a.Equal(1, len(scheme), "len(Event.Scheme)")

	prop := scheme[0]
	a.Equal("count", prop.Name, "Property.Name")
	a.Equal(NumberType, prop.Type, "Property.Type")
}

var invalidPropertyTypeCfg = `
  events:
    - name: Foo
      scheme:
        - name: count
          type: 234
    `

func TestPropertyTypeParsing(t *testing.T) {
	a := assert.New(t)
	validInput := strings.NewReader(invalidPropertyTypeCfg)
	_, err := LoadConfig(validInput)
	if a.Error(err) {
		a.ErrorAs(&ErrUnknownPropertyType{}, &err)
	}
}
