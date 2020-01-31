package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jwlarocque/engine"
)

var diamCollider engine.Collider
var diamColliderOpts *ebiten.DrawImageOptions
var diamColliderImg *ebiten.Image

var squareCollider engine.Collider
var squareColliderOpts *ebiten.DrawImageOptions
var squareColliderImg *ebiten.Image

func drawCollider(image *ebiten.Image, collider engine.Collider) {
	ebitenutil.DrawLine(image, collider.Vertices[len(collider.Vertices)-1].X, collider.Vertices[len(collider.Vertices)-1].Y, collider.Vertices[0].X, collider.Vertices[0].Y, color.RGBA{255, 0, 0, 255})
	for i := 0; i < len(diamCollider.Vertices)-1; i++ {
		ebitenutil.DrawLine(image, collider.Vertices[i].X, collider.Vertices[i].Y, collider.Vertices[i+1].X, collider.Vertices[i+1].Y, color.RGBA{255, 0, 0, 255})
	}
}

func collidersInteractiveUpdate(screen *ebiten.Image) error {
	drawCollider(diamColliderImg, diamCollider)
	drawCollider(squareColliderImg, squareCollider)
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		diamCollider.Position.Y += 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		diamCollider.Position.Y -= 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		diamCollider.Position.X += 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		diamCollider.Position.X -= 1.0
	}
	diamColliderOpts = &ebiten.DrawImageOptions{}
	diamColliderOpts.GeoM.Translate(diamCollider.Position.X, diamCollider.Position.Y)
	screen.DrawImage(diamColliderImg, diamColliderOpts)
	screen.DrawImage(squareColliderImg, squareColliderOpts)
	if diamCollider.Collides(&squareCollider) {
		ebitenutil.DebugPrint(screen, "Colliding!")
	} else {
		ebitenutil.DebugPrint(screen, "Not colliding!")
	}
	return nil
}

func main() {
	diamColl, err := engine.NewCollider([]*engine.Vector2{{10.0, 0.0}, {20.0, 10.0}, {10.0, 20.0}, {0.0, 10.0}})
	if err != nil {
		log.Fatal((err))
	}
	diamCollider = *diamColl
	diamColliderImg, err = ebiten.NewImage(200.0, 200.0, ebiten.FilterDefault)
	if err != nil {
		log.Fatal((err))
	}
	diamColliderOpts = &ebiten.DrawImageOptions{}
	diamColliderOpts.GeoM.Translate(100.0, 50.0)

	squareColl, err := engine.NewCollider([]*engine.Vector2{{0.0, 0.0}, {20.0, 0.0}, {20.0, 20.0}, {0.0, 20.0}})
	if err != nil {
		log.Fatal((err))
	}
	squareCollider = *squareColl
	squareColliderImg, err = ebiten.NewImage(200.0, 200.0, ebiten.FilterDefault)
	if err != nil {
		log.Fatal((err))
	}
	squareColliderOpts = &ebiten.DrawImageOptions{}
	squareColliderOpts.GeoM.Translate(50.0, 100.0)
	squareCollider.Position = engine.Vector2{50.0, 100.0}

	log.Println(squareCollider.String())

	if err := ebiten.Run(collidersInteractiveUpdate, 400, 240, 2, "Interactive Colliders Test"); err != nil {
		log.Fatal(err)
	}
}
