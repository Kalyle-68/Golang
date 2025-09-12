package extras

import(
	"math"
)

type Vector2 struct {
	x float64
	y float64
}

type Vector3 struct {
	x float64
	y float64
	z float64
}

type Dimens2 struct {
	w float64
	h float64
}

type Dimens3 struct {
	w float64
	h float64
	d float32
}

func mouseRot(mX int, mY int) float64 {
	var mouse_pos = Vector2{float64(mX), float64(mY)}
	var dirOfMouse float64 = float64(math.Atan2(float64(mouse_pos.y-300), float64(mouse_pos.x-400)) + math.Pi/4)
	return dirOfMouse
}
func isoVec(x float64, y float64, dir float64, xSqu float64, ySqu float64, ancX float64, ancY float64) Vector2 {
	return Vector2{
		(x*xSqu*math.Cos(dir*math.Pi/180) - y*xSqu*math.Sin(dir*math.Pi/180)) + ancX,
		(x*ySqu*math.Sin(dir*math.Pi/180) + y*ySqu*math.Cos(dir*math.Pi/180)) + ancY,
	}
}