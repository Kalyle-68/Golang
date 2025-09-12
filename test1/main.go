package main

import (
	"fmt"
	"github.com/Kalyle-68/Golang/function/latest"
)

func main() {
	var vector functions.Vector2 = functions.Vector2{8, 6}
	println(fmt.Sprint(vector.x) + ", " + fmt.Sprint(vector.y))
}
