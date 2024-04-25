package main

import (
	"fmt"//for Printing Stuff
	"math/rand"//for Character Generation
)

type Character struct {
	STR, DEX, CON, toHit, AC, init, maxHP, HP int
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
	c.init = (c.DEX - 10)/2
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
func Attack(atk Character, def Character) int{
	d := 0
	hitRoll := RollToHit()
	if hitRoll == 20 {
		fmt.Println("20!! Rolling to Confirm Crit...")
		hitRoll = RollToHit()
		if hitRoll + atk.toHit >= def.AC {
			fmt.Println(hitRoll, "! Crit Confirmed!! Rolling Damage!!")
			d = RollDamage() + RollDamage()
			fmt.Println(d, " Damage!")
			return d
		} else {
			fmt.Println(hitRoll, "... Crit not Confirmed. Rolling Damage...")
			d = RollDamage()
			fmt.Println(d, " Damage")
			return d
		}
	}else if hitRoll + atk.toHit >= def.AC {
			fmt.Println(hitRoll, ", Hit! Rolling Damage...")
			d = RollDamage()
			fmt.Println(d, " Damage")
			return d
		} else {
			fmt.Println(hitRoll, ", Miss")
			return d
		}
}
func main(){
	a := NewCharacter()
	b := NewCharacter()
	aInit := 0
	bInit := 0
	winner := false
	aFirst := true
	order := [2]Character{}
	fmt.Println("Character A", a)
	fmt.Println("Character B", b)
	for aInit == bInit {
		aInit = RollToHit() + a.init
		bInit = RollToHit() + b.init
	}
	if aInit > bInit{
		fmt.Println("A Goes First")
		order[0] = a
		order[1] = b
		aFirst = true
	} else {
		fmt.Println("B Goes First")
		order[0] = a
		order[1] = b
		aFirst = false
	}
	fmt.Println("Time to Fight!")
	for !winner {
		if aFirst{
			fmt.Println("A's Punch, Roll Needed: ", b.AC - a.toHit)
		}else{
			fmt.Println("B's Punch, Roll Needed: ", a.AC - b.toHit)
		}
		order[1].HP -= Attack(order[0], order[1])
		fmt.Println("HP Left:", order[1].HP)
		if order[1].HP <= 0 {
			fmt.Println("Knockout!")
			if aFirst == true{
				fmt.Println("The Winner is A!")
			} else {
				fmt.Println("The Winner is B!")
			}
			winner = true
			break;
		}
		if aFirst == true{
			fmt.Println("B's Punch, Roll Needed: ", a.AC - b.toHit)
		} else {
			fmt.Println("A's Punch, Roll Needed: ", b.AC - a.toHit)
		}
		order[0].HP -= Attack(order[1], order[0])
		fmt.Println("HP Left:", order[0].HP)
		if order[0].HP <= 0 {
			fmt.Println("Knockout!")
			if aFirst{
				fmt.Println("The Winner is B!")
			} else {
				fmt.Println("The Winner is A!")
			}
			winner = true
			break;
		}
	}
	
	return 
}