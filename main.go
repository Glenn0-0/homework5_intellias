package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

//customizable variables for every animal.
var (
	//dog
	dogConsumption = 2
	dogWeightMax   = 30
	dogWeightMin   = 5
	dogNames       = []string{"Bob", "Jack", "Bull", "Peter", "Rick"}
	dogType        = "Dog"

	//cat
	catConsumption = 7
	catWeightMax   = 12
	catWeightMin   = 2
	catNames       = []string{"Kylie", "Bun", "Murka", "Jenna", "Lois"}
	catType        = "Cat"

	//cow
	cowConsumption = 2
	cowWeightMax   = 720
	cowWeightMin   = 700
	cowNames       = []string{"Manya", "Elza", "Ella", "Nanno", "Rita"}
	cowType        = "Cow"
)

//gets random index to choose a name for an animal from slice.
func randNameIndex(numOfNames int) int {
	return rand.Intn(numOfNames)
}

//function creates and adds to a farm a certain number of random animals with random weight.
func createRandFarm(numOfAnimals int) (animals []foodAmountGetter) {

	var entity foodAmountGetter
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOfAnimals; i++ {

		randAnimal := rand.Intn(3) //0 = dog; 1 = cat; 2= cow.

		//based on the number gotten, adding to the farm an animal with random weight.
		switch randAnimal {

		case 0:
			randWeight := rand.Intn(dogWeightMax-dogWeightMin+1) + dogWeightMin //gets random weight for an animal in range of max and min values.

			entity = animal{ //adds an animal according to base animal structure.
				consumption: dogConsumption,
				weight:      randWeight,
				name:        dogNames[randNameIndex(len(dogNames))],
				animalType: dogType,
			}

		case 1:
			randWeight := rand.Intn(catWeightMax-catWeightMin+1) + catWeightMin

			entity = animal{
				consumption: catConsumption,
				weight:      randWeight,
				name:        catNames[randNameIndex(len(catNames))],
				animalType:  catType,
			}

		case 2:
			randWeight := rand.Intn(cowWeightMax-cowWeightMin+1) + cowWeightMin

			entity = animal{
				consumption: cowConsumption,
				weight:      randWeight,
				name:        cowNames[randNameIndex(len(cowNames))],
				animalType:  cowType,
			}
		}

		animals = append(animals, entity)
	}

	return animals
}

func main() {
	totalFoodAmount := 0

	numOfAnimals := 5                    //number of animals in the farm.
	farm := createRandFarm(numOfAnimals) //generates a random farm.

	fmt.Printf("\nFarm info:\n\n")

	for _, an := range farm {
		info, err := an.getInfo() //gets info about every animal in the farm and possiable errors.
		fmt.Print(info) //prints out info for every animal in the farm.
		if err != nil {
			fmt.Println(err) //prints out the error and stops program execution.
			os.Exit(0)
		}
		totalFoodAmount += an.getFoodAmount() //adds required amount of food for a specified animal to total sum.         
	}

	fmt.Printf("\nIn total, this farm needs %vkg of food\n\n", totalFoodAmount) //prints out total sum of food needed to feed this farm.

}
