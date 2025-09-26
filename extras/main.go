package extras

import (
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

func DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func RadToDeg(rad float64) float64 {
	return rad / (math.Pi / 180)
}

func PersProj(oriPoint Vector3, CamPos Vector3, CamRot Vector3, WinDim Dim2, nearPlane float64, farPlane float64, clip bool) Vector2 {
	point := Vector3{
		X: oriPoint.X - CamPos.X,
		Y: oriPoint.Y - CamPos.Y,
		Z: oriPoint.Z - CamPos.Z,
	}
	camRad := Vector3{
		X: DegToRad(-CamRot.X),
		Y: DegToRad(CamRot.Y),
		Z: DegToRad(CamRot.Z),
	}
	tempX := point.X*math.Cos(camRad.X) + point.Z*math.Sin(camRad.X)
	tempZ := -point.X*math.Sin(camRad.X) + point.Z*math.Cos(camRad.X)
	point.X = tempX
	point.Z = tempZ
	tempY := point.Y*math.Cos(camRad.Y) - point.Z*math.Sin(camRad.Y)
	tempZ = point.Y*math.Sin(camRad.Y) + point.Z*math.Cos(camRad.Y)
	point.Y = tempY
	point.Z = tempZ
	tempX = point.X*math.Cos(camRad.Z) - point.Y*math.Sin(camRad.Z)
	tempY = point.X*math.Sin(camRad.Z) + point.Y*math.Cos(camRad.Z)
	point.X = tempX
	point.Y = tempY
	if clip && (point.Z < nearPlane || point.Z > farPlane) {
		point.Z = 0
	}
	scale := Dim2{
		W: math.Min(WinDim.W/2, WinDim.H/2),
		H: math.Min(WinDim.W/2, WinDim.H/2),
	}
	screenX := (point.X/point.Z)*scale.W + WinDim.W/2
	screenY := -(point.Y/point.Z)*scale.H + WinDim.H/2
	return Vector2{X: screenX, Y: screenY}
}
