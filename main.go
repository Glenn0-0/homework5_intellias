package main

import (
	"fmt"
	"math/rand"
	"time"
)

//base animal structure.
//whether its a cat, cow or a dog, they will still have same properties such as consumption of food per kg of its weight, the weight and its name.
type animal struct {
	consumption int
	weight int
	name string
}

type foodAmountGetter interface {
	getFoodAmount() int
	printInfo()
}

//amount of food for one animal per month. Gotten my multiplying the consumption per kg by the weight of the animal.
func (a animal) getFoodAmount() int {
	return int(a.consumption * a.weight)
}

//prints out info for every animal in the farm.
func (a animal) printInfo() {
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

//returns false to ignore an animal if animal's weight <= 0
func checkWeight(weight int) bool {
	if weight <= 0 {
		return false
	}
	return true
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
		switch randAnimal {

		case 1:
			randWeight := rand.Intn(dogWeightMax - dogWeightMin + 1) + dogWeightMin //gets random weight for an animal in range of max and min values.
			realWeight := checkWeight(randWeight) //checks if weight > 0; else ==> ignores an animal and doesn't add it to the farm.

			if realWeight {
				entity = animal{ //adds an animal according to base animal structure.
					consumption: dogConsumption,
					weight: randWeight,
					name: dogName,
				}
			}

		case 2:
			randWeight := rand.Intn(catWeightMax - catWeightMin + 1) + catWeightMin
			realWeight := checkWeight(randWeight) 

			if realWeight {
				entity = animal{
					consumption: catConsumption,
					weight: randWeight,
					name: catName,
				}
			}

		case 3:
			randWeight := rand.Intn(cowWeightMax - cowWeightMin + 1) + cowWeightMin 
			realWeight := checkWeight(randWeight)

			if realWeight {
				entity = animal{
					consumption: cowConsumption,
					weight: randWeight,
					name: cowName,
				}
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

	for _, an := range farm {
		totalFoodAmount += an.getFoodAmount() //adds required amount of food for a specified animal to total sum.
		an.printInfo() //prints out info for every animal in the farm.
	}

	fmt.Printf("\nIn total, this farm needs %vkg of food\n\n", totalFoodAmount) //prints out total sum of food needed to feed this farm.
	
}