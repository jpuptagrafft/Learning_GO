package char_stats

import (
	"math/rand"//for Character Generation
)

type Character struct {
	STR, DEX, CON, toHit, AC, maxHP, HP int
	//STR = Character Strength, Determins toHit
	//DEX = Character Dexterity, Determins AC
	//CON = Character Constitution, Determins HP
	//toHit = (STR - 10)//2
	//AC = 10 + (DEX - 10)//2
	//HP = (CON - 10)//2
	
}
func NewCharacter() Character {
	c := Character{}
	c.STR = statRoll()
	c.DEX = statRoll()
	c.CON = statRoll()
	c.toHit = (c.STR - 10)/2
	if c.STR < 10 {
		c.toHit -= 1
	}
	c.AC = 10 + (c.DEX - 10)/2
	if c.DEX < 10 {
		c.AC -= 1
	}
	c.maxHP = ((c.CON - 10)/2) + 4
	if c.CON < 10 {
		c.maxHP -= 1
	}
	c.HP = c.maxHP 
	return c
}

func statRoll() int{
	x := rand.Intn(6) + 1
	y := rand.Intn(6) + 1
	z := rand.Intn(6) + 1
	return x+y+z
}

func RollToHit() int{
	return rand.Intn(20) + 1
}

func RollDamage() int{
	return rand.Intn(4) + 1
}