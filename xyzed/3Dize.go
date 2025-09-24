package main

import (
	"image/color"
	"math"

	"github.com/Kalyle-68/Golang/extras"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type triangle struct {
	vertices []extras.Vector3
	color    color.Color
}

var nearPlane float64 = 1

func persProj(point extras.Vector3) extras.Vector2 {
	var newX float64 = point.X - CamPos.X
	var newY float64 = -point.Y + CamPos.Y
	var newZ float64 = point.Z - CamPos.Z
	if newZ < nearPlane {
		newZ = 1
	}
	var scale = extras.Dim2{W: FocalLength * 1/math.Tan()), H: 0}
	return extras.Vector2{X: (newX/newZ)* + WinDim.W/2,
		Y: (newY/newZ)*FocalLength + WinDim.H/2}
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
