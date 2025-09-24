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
var CamPos = extras.Vector3{X: 0, Y: 0, Z: -2}
var CamRot = extras.Vector2{X: 90, Y: 60}
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
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		CamPos.Y += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		CamPos.Y -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		CamPos.Z += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		CamPos.Z -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		FocalLength++
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		FocalLength--
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 40, 255})
	drawTriangles(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Focal Length: %f", FocalLength))
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
