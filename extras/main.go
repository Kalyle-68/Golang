package extras

import(
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}


type Dim2 struct {
	W float64
	H float64
}

type Dim3 struct {
	W float64
	H float64
	D float64
}

func MouseRot(x float64, y float64, ancX float64, ancY float64) float64 {
	var pos = Vector2{float64(x), float64(y)}
	var rot float64 = float64(math.Atan2(float64(pos.Y-300), float64(pos.X-400)) + math.Pi/4)
	return rot
}
func IsoVec(x float64, y float64, dir float64, xSqu float64, ySqu float64, ancX float64, ancY float64) Vector2 {
	return Vector2{
		(x*xSqu*math.Cos(dir*math.Pi/180) - y*xSqu*math.Sin(dir*math.Pi/180)) + ancX,
		(x*ySqu*math.Sin(dir*math.Pi/180) + y*ySqu*math.Cos(dir*math.Pi/180)) + ancY,
	}
}