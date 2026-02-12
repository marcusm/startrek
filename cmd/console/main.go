package main

import (
	"fmt"

	"github.com/marcusm/startrek-go/internal"
)

func main() {
	k := internal.Klingon{}
	k.Energy = 200
	k.X = 1
	k.Y = 1

	fmt.Println(k)
}
