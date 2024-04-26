package main

import (
	"fmt"//for Printing Stuff
	"math/rand"//for Character Generation & Dice rolling
	"encoding/json"//for cleaner printing of the Character Struct
)

type Character struct {
	STR, DEX, CON, MaxHP int
	//The variables below me won't show up in the print
	toHit int //`json:"-"`
	AC int //`json:"-"`
	init int //`json:"-"` 
	HP int `json:"-"`
	//STR = Character Strength, Determins toHit
	//DEX = Character Dexterity, Determins AC and Inititive
	//CON = Character Constitution, Determins HP
	//toHit = (STR - 10)//2
	//AC = 10 + (DEX - 10)//2
	//init = (DEX - 10)//2
	//HP = (CON - 10)//2
	
}
func NewCharacter() Character {
	c := Character{}
	c.STR = statRoll()
	c.DEX = statRoll()
	c.CON = statRoll()
	//The Base Stats are determined via a roll of 3 D6 Dice
	c.toHit = (c.STR - 10)/2
	if c.STR < 10 && c.STR%2 == 1{
		c.toHit -= 1
	}
	c.AC = 10 + ((c.DEX - 10)/2)
	if c.DEX < 10 && c.DEX%2 == 1{
		c.AC -= 1
	}
	c.init = (c.DEX - 10)/2
	if c.DEX < 10 && c.DEX%2 == 1{
		c.init -= 1
	}
	c.MaxHP = ((c.CON - 10)/2) + 4
	if c.CON < 10 && c.CON%2 == 1{
		c.MaxHP -= 1
	}
	//For our Characters, we are assuming only a level in commoner; therefore we are only giving them their modifiers in toHit, Armor Class, Inititive, and HP
	//To determine Modifiers, Use (MOD - 10)/2, Rounded down. So, a score of 9 would result in a -1 modifier, for example
	//This is why we have that odd "if c.CON < 10 && c.CON%2 == 1" clause
	c.HP = c.MaxHP 
	return c
}

func statRoll() int{
	x := rand.Intn(6) + 1
	y := rand.Intn(6) + 1
	z := rand.Intn(6) + 1
	//Rolling 3 D6
	return x+y+z
}

func RollToHit() int{
	return rand.Intn(20) + 1
	//Rolling a D20
}

func RollDamage() int{
	return rand.Intn(4) + 1
	//Rolling a D4
}
func Attack(atk Character, def Character) int{
	d := 0
	hitRoll := RollToHit()
	if hitRoll == 20 { //A natural roll of 20 will threaten a crit if it would hit itself. As in this system the highest you would need to roll to hit is a 19, a 20 always hits.
		fmt.Println("20!! Rolling to Confirm Crit...")
		hitRoll = RollToHit()
		if hitRoll + atk.toHit >= def.AC { //If the hit roll after threatening a crit plus the to hit modifier equals or exceeds the targets Armor class, the crit is confirmed, and deals double damage. Otherwise, it simply does normal damage 
			fmt.Println(hitRoll, "! Crit Confirmed!! Rolling Damage!!")
			d = RollDamage() + RollDamage() + atk.toHit * 2
			if d < 1 { //Minimum Damage is always 1
				d = 1
			}
			fmt.Println(d, " Damage!")
			return d
		} else {
			fmt.Println(hitRoll, "... Crit not Confirmed. Rolling Damage...")
			d = RollDamage() + atk.toHit
			if d < 1 { //Minimum Damage is always 1
				d = 1
			}
			fmt.Println(d, " Damage")
			return d
		}
	}else if hitRoll + atk.toHit >= def.AC { //if the hit roll plus the to hit modifier equals or exceeds the targets Armor class, the attacker hits and deals damage. Otherwise, the attack misses
			fmt.Println(hitRoll, ", Hit! Rolling Damage...")
			d = RollDamage() + atk.toHit
			if d < 1 { //Minimum Damage is always 1
				d = 1
			}
			fmt.Println(d, " Damage")
			return d
	} else {
		fmt.Println(hitRoll, ", Miss")
		return d
	}
}
func main(){
	a := NewCharacter() //Create Characters
	b := NewCharacter()
	ap, ape := json.MarshalIndent(a, "", " ") //Create cleaner printouts for Characters
	bp, bpe := json.MarshalIndent(b, "", " ")
	if ape != nil || bpe != nil{ //Error checking; This shouldn't run.
		fmt.Println("an Error!? Prepostorous!")
		fmt.Println(ape)
		fmt.Println(bpe)
		return
	}
	aInit := 0 //Initial Variables
	bInit := 0
	winner := false
	aFirst := true
	order := [2]Character{} //Used to dictate Turn Order
	fmt.Println("Character A", string(ap))
	fmt.Println("Character B", string(bp))
	for aInit == bInit { //Inititive Roll off, with Ties rerolling
		aInit = RollToHit() + a.init
		bInit = RollToHit() + b.init
	}
	fmt.Println("Inititive Roll")
	fmt.Println("A: ", aInit)
	fmt.Println("B: ", bInit)
	if aInit > bInit{
		fmt.Println("A Goes First")
		order[0] = a
		order[1] = b
		aFirst = true
	} else {
		fmt.Println("B Goes First")
		order[1] = a
		order[0] = b
		aFirst = false
	}
	fmt.Println("Time to Fight!")
	for !winner { //Check to see if Fight is over
		if aFirst{
			fmt.Println("A's Punch, Roll Needed: ", b.AC - a.toHit)
		}else{
			fmt.Println("B's Punch, Roll Needed: ", a.AC - b.toHit)
		}
		order[1].HP -= Attack(order[0], order[1])
		fmt.Println("HP Left:", order[1].HP)
		if order[1].HP <= 0 { //Check to see if Attack Knocked out Defender
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
		if order[0].HP <= 0 { //Check to see if Attack Knocked out Defender
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