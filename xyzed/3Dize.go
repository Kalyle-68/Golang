package main

import (
	"image/color"

	"github.com/Kalyle-68/Golang/extras"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type triangle struct {
	vertices []extras.Vector3
	color    color.Color
}

var nearPlane float64 = 0.25
var farPlane float64 = 32

/*func persProj(oriPoint extras.Vector3, CamPos extras.Vector3, CamRot extras.Vector3, nearPlane = 0.25, farPlane = 32) extras.Vector2 {
	point := extras.Vector3{
		X: oriPoint.X - CamPos.X,
		Y: oriPoint.Y - CamPos.Y,
		Z: oriPoint.Z - CamPos.Z,
	}
	camRad := extras.Vector3{
		X: extras.DegToRad(-CamRot.X),
		Y: extras.DegToRad(CamRot.Y),
		Z: extras.DegToRad(CamRot.Z),
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
	if point.Z < nearPlane || point.Z > farPlane {
		point.Z = 0
	}
	scale := extras.Dim2{
		W: math.Min(WinDim.W/2, WinDim.H/2),
		H: math.Min(WinDim.W/2, WinDim.H/2),
	}
	screenX := (point.X/point.Z)*scale.W + WinDim.W/2
	screenY := -(point.Y/point.Z)*scale.H + WinDim.H/2
	return extras.Vector2{X: screenX, Y: screenY}
}*/

func persProj(oriPoint extras.Vector3) extras.Vector2 {
	return extras.PersProj(oriPoint, CamPos, CamRot, WinDim, 0.25, 32)
}

func drawTriangles(screen *ebiten.Image) {
	for triangle := 0; triangle < len(Env); triangle++ {
		//var Env2D []ebiten.Vertex
		//var indices []uint16
		for vertex := 0; vertex < len(Env[triangle].vertices); vertex++ {
			var curPoint extras.Vector2 = persProj(Env[triangle].vertices[vertex])
			var nextPoint extras.Vector2 = persProj(Env[triangle].vertices[(vertex+1)%len(Env[triangle].vertices)])
			vector.StrokeLine(screen, float32(curPoint.X), float32(curPoint.Y),
				float32(nextPoint.X), float32(nextPoint.Y), 8, Env[triangle].color, true)
			/*
				var R, G, B, A = Env[triangle].color.RGBA()
				var srcX float32 = 0
				var srcY float32 = 0
				Env2D = append(Env2D, ebiten.Vertex{DstX: float32(point.X), DstY: float32(point.Y),
					SrcX: srcX, SrcY: srcY, ColorR: float32(R / 255), ColorG: float32(G / 255), ColorB: float32(B / 255), ColorA: float32(A / 255)})
				indices = append(indices, uint16(vertex))*/
		} /*
			texture := ebiten.NewImage(1,1)
			texture.Fill(color.Transparent)
			options := &ebiten.DrawTrianglesOptions{}
			polygon := ebiten.NewImage(800, 600)
			screen.DrawTriangles(Env2D, indices, texture, options)
			screen.DrawImage(polygon, nil)*/
	}
}
