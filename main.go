package main

import (
	"math/rand"
	"log"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

var ( //declvare variable for images, name *ebiten.Image.
	background *ebiten.Image
	player1 *ebiten.Image

	axeZombieSprites []*ebiten.Image

	lightSaber *ebiten.Image

	screenHeight = 1080
	screenWidth = 1920

	player1InitX = float64(560)
	player1InitY = float64(240)
	axeZombieInitXTemp = float64 (randFloat(1,100))
	axeZombieInitYTemp = float64 (randFloat(1,100))
	lightSaberX float64
	lightSaberY float64

	player1hp = 20

	tickCount = 0

	zombies []axeZombie
)

type axeZombie struct{
	level 	int
	hp 			int
	x, y 		float64
	speed		float64
	}



func spawnZombies() {
  count := randInt(3, 6)

  for i := 0; i < count; i++ {
    z := axeZombie{
    x:     randFloat(0, float64(screenWidth + 100)),
    y:     randFloat(0, float64(screenHeight + 100)),
    hp:    randInt(3, 10),
    level: randInt(1, 3),
    speed: randFloat(0.3, 1.0),
    }
    
		zombies = append(zombies, z)
  }
}

func enemyMovement(targetX, targetY, enemyX, enemyY, speed float64) (float64, float64) {
	if enemyX < (targetX - 80){ //enemie movement
		enemyX += speed
	}
	if enemyX > (targetX + 80){
		enemyX -= speed
	}
	if enemyY < (targetY - 80){
		enemyY += speed
	}
	if enemyY > (targetY + 80){
		enemyY -= speed
	}

	return enemyX, enemyY
}

func loadAxeZombieSprites() {
  axeZombieSprites = make([]*ebiten.Image, 8)

  for i := 1; i <= 8; i++ {
    filename := fmt.Sprintf("assets/sprites/enemies/axeZombie/axeZombieSprite%02d.png", i)

    img, _, err := ebitenutil.NewImageFromFile(filename)
    if err != nil {
    log.Fatal(err)
    }
    axeZombieSprites[i-1] = img
	}
}

func abs(f float64) float64 {
  if f < 0 {
    return -f
  }
  return f
}

func randFloat(min, max float64) float64 {
  return min + rand.Float64()*(max-min)
}


func isBlocked(px, py float64, dx, dy float64, blockRange float64, zombies []axeZombie) bool {
  for _, z := range zombies {
  	// Project the check range in the direction the player wants to move
    checkX := px + dx*blockRange
    checkY := py + dy*blockRange

    // If an enemy is near that projected point → blocked
    if abs(z.x-checkX) < 50 && abs(z.y-checkY) < 50 {
      return true
    }
  }
 
	return false
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

	lightSaber, _, err = ebitenutil.NewImageFromFile("assets/images/lightSaber.png") //will not run if empty
	if err != nil {
		log.Fatal(err)
	}

	loadAxeZombieSprites()
	spawnZombies()
}

type Game struct{}



func (g *Game) Update() error { //game logic

	tickCount++
	
	lightSaberX = float64 (player1InitX + 100)
	lightSaberY = float64 (0)


moveSpeed := 3.0
blockRange := 40.0

// MOVE RIGHT (D)
if ebiten.IsKeyPressed(ebiten.KeyD) &&
  !isBlocked(player1InitX, player1InitY, 1, 0, blockRange, zombies) {
  player1InitX += moveSpeed
}

// MOVE LEFT (A)
if ebiten.IsKeyPressed(ebiten.KeyA) &&
  !isBlocked(player1InitX, player1InitY, -1, 0, blockRange, zombies) {
  player1InitX -= moveSpeed
}

// DOWN (S)
if ebiten.IsKeyPressed(ebiten.KeyS) &&
  !isBlocked(player1InitX, player1InitY, 0, 1, blockRange, zombies) {
  player1InitY += moveSpeed
}

// UP (W)
if ebiten.IsKeyPressed(ebiten.KeyW) &&
  !isBlocked(player1InitX, player1InitY, 0, -1, blockRange, zombies) {
	player1InitY -= moveSpeed
}


for i := range zombies {
  
	zombies[i].x, zombies[i].y = enemyMovement(
    player1InitX,
  	player1InitY,
    zombies[i].x,
    zombies[i].y,
    zombies[i].speed,
  )

  hitRange := 80.0 // damage player if close
  if abs(zombies[i].x-player1InitX) < hitRange &&
  abs(zombies[i].y-player1InitY) < hitRange {
    if tickCount%150 == 0 {
      player1hp--
    }
  }
}

  fmt.Println("hp:", player1hp)

	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {  //called every frame, graphics.
	
	screen.DrawImage(background, nil)

	op := &ebiten.DrawImageOptions{}	
	opAxeZombie := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(player1InitX,player1InitY)
	opAxeZombie.GeoM.Translate(axeZombieInitXTemp,axeZombieInitYTemp)		

	screen.DrawImage(player1, op)	

	opLightSaber := &ebiten.DrawImageOptions{} //todo: fix
	opLightSaber.GeoM.Translate(lightSaberX, lightSaberY)

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) == true {
		screen.DrawImage(lightSaber, opLightSaber)
	}

frame := (tickCount / 8) % len(axeZombieSprites)
sprite := axeZombieSprites[frame]

	for _, z := range zombies {
		op := &ebiten.DrawImageOptions{}
 		op.GeoM.Translate(z.x, z.y)
  	screen.DrawImage(sprite, op)
	}
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Render an image")
	
	if err := ebiten.RunGame(&Game{}); err != nil { 
		log.Fatal(err)
	}	
}
