
//	Chaos Crashers copyright (c) 2026 River Knuuttila, common alias: Annie Valentine or aval. All Rights Reserved.
//	Do not redistribute or reuse code without accrediting and explicit permission from author.
//	Contact: 
//	+1 (808) 223 4780
//	riverknuuttila2@outlook.com

package main

import "fmt"

func zombieLogic() {

	zombieWalkCycleUpdate(axeZombieAnimationSpeed)
	zombieHitAnimationUpdate(axeZombieHitAnimationSpeed)	
	zombieDeathAnimationUpdate(3)

	for i := range zombies { //keeps track of how long zombies should be "hit" for
		if zombies[i].hitTimer > 0 {
			zombies[i].hitTimer--
			zombies[i].hit = true
		} else if zombies[i].hitTimer == 0 {
			zombies[i].hit = false
		}
	}

	for i := range zombies { //zombie ai / logic

		if zombies[i].hp <= 0 {
			continue
		}
		
		// movement (once per zombie)
		zombies[i].x, zombies[i].y = enemyMovement(
			p.x,
			p.y,
			zombies[i].x,
			zombies[i].y,
			zombies[i].speed,
			zombies[i].knockbackSpeed,
			p.swordLocation,
			zombies,
			i,
		)
		
		//zombie attack using displayed sprite sizes
		if axeZombieSprites != nil && len(axeZombieSprites) > 0 && axeZombieSprites[0] != nil && player1 != nil {
			zombieW := float64(axeZombieSprites[0].Bounds().Dx()) * axeZombieSpriteScale
			zombieH := float64(axeZombieSprites[0].Bounds().Dy()) * axeZombieSpriteScale
			playerW := float64(player1.Bounds().Dx())
			playerH := float64(player1.Bounds().Dy())

			zombieCX := zombies[i].x + zombieW/2
			zombieCY := zombies[i].y + zombieH/2
			playerCX := p.x + playerW/2
			playerCY := p.y + playerH/2

			if abs(zombieCX-playerCX) < (zombieW+playerW)/2 &&
				abs(zombieCY-playerCY) < (zombieH+playerH)/2 &&
				tickCount%150 == 0 {
				p.hp--
				fmt.Println("hp:", p.hp)
			}
		}

		//	Player attack using displayed sword and zombie sprite sizes.
		if p.attackActive && swordSprites != nil && len(swordSprites) > 0 && swordSprites[0] != nil &&
			axeZombieSprites != nil && len(axeZombieSprites) > 0 && axeZombieSprites[0] != nil {

			const swordScale = 2.0
			frame := p.attackFramesTimer
			if frame >= len(swordSprites) {
				frame = len(swordSprites) - 1
			}
			sw := float64(swordSprites[frame].Bounds().Dx())
			sh := float64(swordSprites[frame].Bounds().Dy())

			//	Visual center of sword matches the draw transform: translate (-cx,-cy) -> scale → rotate -> translate (swordX+cx, swordY+cy).
			swordCX := p.swordX + sw/2
			swordCY := p.swordY + sh/2

			//	Half-dimensions of displayed sword, swap axes for vertical swings 90 degree rotation.
			var swordHalfW, swordHalfH float64
			switch p.swordLocation {
			case 'd', 'a':
				swordHalfW = sw * swordScale / 2
				swordHalfH = sh * swordScale / 2
			case 'w', 's':
				swordHalfW = sh * swordScale / 2
				swordHalfH = sw * swordScale / 2
			}

			zombieW := float64(axeZombieSprites[0].Bounds().Dx()) * axeZombieSpriteScale
			zombieH := float64(axeZombieSprites[0].Bounds().Dy()) * axeZombieSpriteScale
			zombieCX := zombies[i].x + zombieW/2
			zombieCY := zombies[i].y + zombieH/2

			if abs(swordCX-zombieCX) < swordHalfW+zombieW/2 && !zombies[i].invulnerable &&
				abs(swordCY-zombieCY) < swordHalfH+zombieH/2 &&
				zombies[i].hitTimer <= 0 {
				zombies[i].hp--
				zombies[i].hit = true
				zombies[i].inHitAnimation = true
				zombies[i].hitTimer = p.hitFrameDuration
				zombies[i].hitFrame = 0
				zombies[i].hitAnimTimer = 0
				zombies[i].invulnerable = true
				fmt.Println("Zombie", i, "hp:", zombies[i].hp)
			}
		}

		if zombies[i].hit {
			zombies[i].invulnerable = true
		}
	}
		
		if p.hitFrameDuration > 0 {
			p.attackActive = true
			p.hitFrameDuration--
	}
}
