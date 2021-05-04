package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dragon struct{
	Name string
	HitPoint int
	Damage int
}
type Hero struct{
	Name string
	HitPoint int
	Damage int
}

func main(){

	dragon1 := Dragon{
		Name:     DragonNameAdder() ,
		HitPoint: 20,
		Damage:   2,
	}
	hero1 := Hero{
		Name: humanNameAdder(),
		HitPoint: 10,
		Damage: 4,
	}
	StoryTelling(&dragon1, &hero1)

}

func DragonNameAdder() string{
	rand.Seed(time.Now().UnixNano())
	var name string

	PotentialNames := []string{"Saphira", "Smaug", "Draco", "Puff", "Concar"}
	name = PotentialNames[rand.Intn(len(PotentialNames))]

	return name
}
func humanNameAdder() string{
	rand.Seed(time.Now().UnixNano())
	var name string

	PotentialNames := []string{"Gerald", "Richard", "Steward", "Dooku"}
	name = PotentialNames[rand.Intn(len(PotentialNames))]

	return name
}
func StoryTelling(dragon *Dragon, hero *Hero){
	fmt.Printf(`	A long, long time ago there was a village settled just near a river full of delicious fishes. 
People on this town were happy and friendly. These villagers were living in peace until a monstrous drgon by the 
name of %s appeared from the mountains on the north of the village. It was destroying villagers' homes and crops
leaving them without shelter or food. The old leader of the village sent a wonderer by the name of Dandelion to 
seek a Hero whom would rescue them from this treacherous beast! Dandelion finally found a mercenary named %s who 
would be able to fight with the Dragon %s. %s, traveled to the village and climbed the mountain which the Dragon 
was living in. The Dragon %s, saw %s and they started fighting. %s`,dragon.Name, hero.Name, dragon.Name, hero.Name, dragon.Name, hero.Name, "\n\n")
	time.Sleep(time.Second * 10)
	//time.Sleep(time.Second * 2)

	GamePlay(dragon, hero)


}
func GamePlay(dragon *Dragon, hero *Hero){
	rand.Seed(time.Now().UnixNano())

	var(
		playerDice int
		dragonDice int
	)

	for hero.HitPoint > 0 && dragon.HitPoint > 0{
		fmt.Println("To throw a dice please enter a key...")
		_, err := fmt.Scanln()
		if err != nil{
			fmt.Println("Wrong input")
		}

		playerDice = rand.Intn(10)
		fmt.Printf("You threw %v\n", playerDice)
		time.Sleep(time.Second)

		dragonDice = rand.Intn(10)
		fmt.Printf("Dragon threw %v\n\n",dragonDice)
		time.Sleep(time.Second)

		switch {
		case playerDice > dragonDice:
			DragonTakingDamage(dragon, hero)
		case playerDice < dragonDice:
			HeroTakingDamage(dragon, hero)
		case playerDice == dragonDice:
			fmt.Printf("You and the Dragon both threw the same number\n\n")
		}
	}

	switch  {
	case hero.HitPoint <= 0:
		fmt.Println("YOU LOST!")
	case dragon.HitPoint <= 0:
		fmt.Println("YOU WON!")
	}

}

func HeroTakingDamage(dragon *Dragon, hero *Hero){
	fmt.Printf("Dragon hit you %v damage\n", dragon.Damage)
	hero.HitPoint = hero.HitPoint - dragon.Damage
	fmt.Printf("Your reamaining health is %v\n\n", hero.HitPoint)
}

func DragonTakingDamage(dragon *Dragon, hero *Hero){
	fmt.Printf("You hit %v damage\n", hero.Damage)
	dragon.HitPoint = dragon.HitPoint - hero.Damage
	fmt.Printf("Reamaining Dragon health is %v\n\n", dragon.HitPoint)
}
