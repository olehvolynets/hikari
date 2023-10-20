package main

import (
// "encoding/json"
// "fmt"
)

type a struct {
	A int
	B float64
	C string
	D []bool
	E b
	F uint
}

type b struct {
	EA string
	EB int
}

var sample a = a{
	A: -123,
	B: 456.789,
	C: "lskjdlaksj",
	D: []bool{true, false, false, true},
	E: b{EA: "EAdeeznuts", EB: 0},
	F: 42,
}

func init() {
	// bs, _ := json.Marshal(sample)
	// fmt.Println(string(bs))
}
