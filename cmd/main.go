package main

import (
	"cli-chess/internal/assignSymbol"
)

type Player struct {
	name string
	color
}

func main() {
	assignSymbol.AssignSymbol()
}
