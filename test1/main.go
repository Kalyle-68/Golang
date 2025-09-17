package main

import (
	"fmt"
	"github.com/Kalyle-68/Golang/extras"
)

func main() {
	var vector extras.Vector2 = extras.Vector2{8, 6}
	println(fmt.Sprint(vector.X) + ", " + fmt.Sprint(vector.Y))
}
