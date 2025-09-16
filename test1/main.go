package main

import (
	"fmt"
	"github.com/Kalyle-68/Golang/extras"
)

func main() {
	var vector extras.main.Vector2 = extras.main.Vector2{8, 6}
	println(fmt.Sprint(vector.x) + ", " + fmt.Sprint(vector.y))
}
