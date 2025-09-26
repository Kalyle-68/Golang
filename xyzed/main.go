package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Kalyle-68/Golang/extras"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var WinDim = extras.Dim2{W: 1000, H: 750}
var FocalLength float64 = 860
var CamPos = extras.Vector3{X: 0, Y: 0, Z: 0}
var CamRot = extras.Vector3{X: 0, Y: 0, Z: 0}
var Env = []triangle{
	{[]extras.Vector3{{X: -1, Y: -1, Z: 1}, {X: -1, Y: 1, Z: 1},
		{X: 1, Y: 1, Z: 1}, {X: 1, Y: -1, Z: 1}}, color.RGBA{150, 75, 255, 255}},
	{[]extras.Vector3{{X: -1, Y: -1, Z: 3}, {X: -1, Y: 1, Z: 3},
		{X: 1, Y: 1, Z: 3}, {X: 1, Y: -1, Z: 3}}, color.RGBA{150, 75, 255, 255}},
	{[]extras.Vector3{{X: -1, Y: -1, Z: 1}, {X: 1, Y: -1, Z: 1},
		{X: 1, Y: -1, Z: 3}, {X: -1, Y: -1, Z: 3}}, color.RGBA{150, 75, 255, 255}},
	{[]extras.Vector3{{X: -1, Y: 1, Z: 1}, {X: 1, Y: 1, Z: 1},
		{X: 1, Y: 1, Z: 3}, {X: -1, Y: 1, Z: 3}}, color.RGBA{150, 75, 255, 255}},
	//{[]extras.Vector3{{X: 1, Y: 1, Z: 1}, {X: -1, Y: 1, Z: 1}, {X: 1, Y: -1, Z: 1}}, color.RGBA{255, 255, 255, 255}},
}

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		CamPos.X += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		CamPos.X -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		CamPos.Y += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		CamPos.Y -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		CamPos.Z += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		CamPos.Z -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		FocalLength++
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		FocalLength--
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		CamRot.Z++
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		CamRot.Z--
	}
	var x, y = ebiten.CursorPosition()
	CamRot = extras.Vector3{X: (float64(x) - WinDim.W/2) / (WinDim.W / 360),
		Y: -(float64(y) - WinDim.H/2) / (WinDim.H / 360), Z: CamRot.Z}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 40, 255})
	drawTriangles(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Relative Mouse Position: %f, %f", CamRot.X, CamRot.Y))
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(WinDim.W), int(WinDim.H)
}
func main() {
	ebiten.SetWindowSize(int(WinDim.W), int(WinDim.H))
	ebiten.SetWindowTitle("XYZed engine v1.0.0")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
