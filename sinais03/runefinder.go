package main

import (
	"strconv"
	"strings"
)

func AnalisarLinha(ucdLine string) (rune, string) {
	campos := strings.Split(ucdLine, ";")
	codigo, _ := strconv.ParseInt(campos[0], 16, 32)
	return rune(codigo), campos[1]
}
