package main

import (
	"fmt"
	"math/rand"
	"time"
)

//base animal structure.
//whether its a cat, cow or a dog, they will still have same properties such as consumption of food per kg of its weight, the weight and its name.
type Animal struct {
	consumption int
	weight int
	name string
}

type foodAmountGetter interface {
	getFoodAmount() int
	printInfo()
}

//amount of food for one animal per month. Gotten my multiplying the consumption per kg by the weight of the animal.
func (a Animal) getFoodAmount() int {
	return int(a.consumption * a.weight)
}

//prints out info for every animal in the farm.
func (a Animal) printInfo() {
	fmt.Printf("%s: weights %dkg and needs %vkg of food per month.\n", a.name, a.weight, a.getFoodAmount())
}

//customizable constants for every animal.
const (
	//dog
	dogConsumption = 2
	dogWeightMax = 30
	dogWeightMin = 5
	dogName = "Dog"

	//cat
	catConsumption = 7
	catWeightMax = 12
	catWeightMin = 2
	catName = "Cat"

	//cow
	cowConsumption = 2
	cowWeightMax = 720
	cowWeightMin = 700
	cowName = "Cow"
)

//breaks program if animal's weight <= 0
func checkWeight(weight int) {
	if weight <= 0 {
		panic("Weight can't be lower or equal to 0!")
	}
}

//function creates and adds to a farm a certain number of random animals with random weight.
func createRandFarm(numOfAnimals int) (animals []foodAmountGetter) {

	var entity foodAmountGetter
	min := 1 //indexes to get a random animal - see code below.
	max := 3
	rand.Seed(time.Now().UnixNano())

	for i:= 0; i < numOfAnimals; i++ {

		randAnimal := rand.Intn(max - min + 1) + min //1 = dog; 2 = cat; 3 = cow.

		//based on the number gotten, adding to the farm an animal with random weight.
		if randAnimal == 1 {
			randWeight := rand.Intn(dogWeightMax - dogWeightMin + 1) + dogWeightMin //gets random weight for an animal in range of max and min values.
			checkWeight(randWeight) //checks if weight > 0; else ==> panic.

			entity = Animal{ //adds an animal according to base animal structure.
				consumption: dogConsumption,
				weight: randWeight,
				name: dogName,
			}

		} else if randAnimal == 2 {
			randWeight := rand.Intn(catWeightMax - catWeightMin + 1) + catWeightMin
			checkWeight(randWeight)

			entity = Animal{
				consumption: catConsumption,
				weight: randWeight,
				name: catName,
			}

		} else if randAnimal == 3 {
			randWeight := rand.Intn(cowWeightMax - cowWeightMin + 1) + cowWeightMin
			checkWeight(randWeight)

			entity = Animal{
				consumption: cowConsumption,
				weight: randWeight,
				name: cowName,
			}
		}

		animals = append(animals, entity)
	}
	
	return animals
}

func main() {
	totalFoodAmount := 0

	numOfAnimals := 5 //number of animals in the farm.
	farm := createRandFarm(numOfAnimals) //generates a random farm.

	fmt.Println("\nFarm info:\n")

	for _, animal := range farm {
		totalFoodAmount += animal.getFoodAmount() //adds required amount of food for a specified animal to total sum.
		animal.printInfo() //prints out info for every animal in the farm.
	}

	fmt.Printf("\nIn total, this farm needs %vkg of food\n\n", totalFoodAmount) //prints out total sum of food needed to feed this farm.
	
}