package main

import "fmt"

//base animal structure.
//whether its a cat, cow or a dog, they will still have same properties such as consumption of food per kg of its weight, the weight, its name and type.
type animal struct {
	consumption int
	weight      int
	name        string
	animalType  string
}


type foodAmountGetter interface {
	getFoodAmount() int
	getInfo() (string, error)

	validateName() error
	validateWeight() error
	isEdible() (string, string, error)
}

//amount of food for one animal per month. Gotten my multiplying the consumption per kg by the weight of the animal.
func (a animal) getFoodAmount() int {
	return int(a.consumption * a.weight)
}

//returns info and possible errors for every animal in the farm.
func (a animal) getInfo() (string, error) {
	if err := a.validateName(); err != nil { //validation starting point.
		return "", fmt.Errorf("for %s %s: validation failed: %w\n", a.animalType, a.name, err)
	}
	return fmt.Sprintf("%s: weights %dkg and needs %vkg of food per month.\n", a.name, a.weight, a.getFoodAmount()), nil

}


//check if name corresponds to the animal type.
func (a animal) validateName() error {

	//validate name accordance to type. Iterating through list of names and if matched, remembering the type of animal this name list was for.
	var realType string

	for _, name := range dogNames {
		if a.name == name {
			realType = dogType
		}
	}

	for _, name := range catNames {
		if a.name == name {
			realType = catType
		}
	}

	for _, name := range cowNames {
		if a.name == name {
			realType = cowType
		}
	}

	if realType != a.animalType { //comparing given animal type with the remembered one, if differs - error occurs.
		return fmt.Errorf("mismatched name and type: %s has to be a %s: %s is a %s", a.name, realType, a.name, a.animalType)
	}

	//validate the weight of an animal in case no prior errors occurred.
	if err := a.validateWeight(); err != nil {
		return err
	}

	return nil
}

//check if this animal has proper weight.
func (a animal) validateWeight() error {

	switch a.animalType { //checking if weight is in between min and max values
	case dogType:

		if dogWeightMin > a.weight {
			return fmt.Errorf("improper weight: minimal dog's weight is %v kg: %s's weight is %v kg", dogWeightMin, a.name, a.weight)
		} else if a.weight > dogWeightMax {
			return fmt.Errorf("improper weight: maximal dog's weight is %v kg: %s's weight is %v kg", dogWeightMax, a.name, a.weight)
		}

	case catType:

		if catWeightMin > a.weight {
			return fmt.Errorf("improper weight: minimal cat's weight is %v kg: %s's weight is %v kg", catWeightMin, a.name, a.weight)
		} else if a.weight > catWeightMax {
			return fmt.Errorf("improper weight: maximal cat's weight is %v kg: %s's weight is %v kg", catWeightMax, a.name, a.weight)
		}

	case cowType:

		if cowWeightMin > a.weight {
			return fmt.Errorf("improper weight: minimal cow's weight is %v kg: %s's weight is %v kg", cowWeightMin, a.name, a.weight)
		} else if a.weight > cowWeightMax {
			return fmt.Errorf("improper weight: maximal cow's weight is %v kg: %s's weight is %v kg", cowWeightMax, a.name, a.weight)
		}

	}

	return nil
}

//check if this animal is edible. If not edible, returns animal's type, name and error itself.
func (a animal) isEdible() (string, string, error) {
	if a.animalType == dogType || a.animalType == catType {
		return a.animalType, a.name, fmt.Errorf("%[1]s is not edible: %[1]s is a %s", a.name, a.animalType) //an animal isn't edible if it's a dog or a cat.
	}
	return "", "", nil
}