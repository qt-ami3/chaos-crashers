package main

import (
	"math/rand"
	"log"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var ( //declvare variable for images, name *ebiten.Image.
	background *ebiten.Image
	player1 *ebiten.Image
	axeZombie *ebiten.Image
	screenWidth = 750
	screenHeight = 750
	player1InitX = float64((screenWidth / 2) + 50)
	player1InitY = float64((screenHeight / 2) + 50)
	axeZombieInitXTemp = float64 (randFloat(1,100))
	axeZombieInitYTemp = float64 (randFloat(1,100))
	player1hp = 9999
)


func abs(f float64) float64 {
  if f < 0 {
    return -f
  }
  return f
}

func randFloat(min, max float64) float64 {
  return min + rand.Float64()*(max-min)
}

func init() { //initialize images to variables here.
	var err error
	
	background, _, err = ebitenutil.NewImageFromFile("assets/images/go.png") //name, _, etc.
	if err != nil {
		log.Fatal(err)
	}

	player1, _, err = ebitenutil.NewImageFromFile("assets/images/Sprite-0001.png") //will not run if empty
	if err != nil {
		log.Fatal(err)
	}

	axeZombie, _, err = ebitenutil.NewImageFromFile("assets/sprites/enemies/axeZombie/Sprite-0002.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error { //game logic
	
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) == true { //player movement, inverted
		player1InitX = player1InitX + 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) == true {
		player1InitX = player1InitX - 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) == true {
		player1InitY = player1InitY + 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) == true {
		player1InitY = player1InitY - 2
	}

	if axeZombieInitXTemp < (player1InitX - 80){ //enemie movement
		axeZombieInitXTemp++
	}
	if axeZombieInitXTemp > (player1InitX + 80){
		axeZombieInitXTemp--
	}
	if axeZombieInitYTemp < (player1InitY - 80){
		axeZombieInitYTemp++
	}
	if axeZombieInitYTemp > (player1InitY + 80){
		axeZombieInitYTemp--
	}


	// enemy damage when close enough
	hitRange := 100.0 // adjust to taste

	if abs(axeZombieInitXTemp - player1InitX) < hitRange && abs(axeZombieInitYTemp - player1InitY) < hitRange {
  player1hp--
  fmt.Println("hp:", player1hp)
}
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {  //called every frame, graphics.
	screen.DrawImage(background, nil)

	op := &ebiten.DrawImageOptions{}	
	opAxeZombie := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(player1InitX,player1InitY)
	opAxeZombie.GeoM.Translate(axeZombieInitXTemp,axeZombieInitYTemp)
	
	screen.DrawImage(axeZombie, opAxeZombie)

	screen.DrawImage(player1, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(screenHeight, screenWidth)
	ebiten.SetWindowTitle("Render an image")
	
	if err := ebiten.RunGame(&Game{}); err != nil { 
		log.Fatal(err)
	}
	
}
