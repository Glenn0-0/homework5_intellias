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
	isEdible() error
}

//amount of food for one animal per month. Gotten my multiplying the consumption per kg by the weight of the animal.
func (a animal) getFoodAmount() int {
	return int(a.consumption * a.weight)
}

//returns info and possible errors for every animal in the farm.
func (a animal) getInfo() (string, error) {
	err := a.validateName() //validation starting point.
	if err != nil {
		err = fmt.Errorf("for %s %s: validation failed: %w\n", a.animalType, a.name, err)
	}
	return fmt.Sprintf("%s: weights %dkg and needs %vkg of food per month.\n", a.name, a.weight, a.getFoodAmount()), err

}


//ckeck if name corresponds to the animal type.
func (a animal) validateName() (err error) {

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

	if realType != a.animalType { //comparing given animal type with the remebered one, if differs - error occurs.
		err = fmt.Errorf("mismatched name and type: %s has to be a %s: %s is a %s", a.name, realType, a.name, a.animalType)
	}

	//validate the weight of an animal in case no prior errors occured.
	if err == nil {
		err = a.validateWeight()
	}

	return err
}

//check if this abimal has proper weight.
func (a animal) validateWeight() (err error) {

	switch a.animalType { //checking if weight is inbetween min and max values
	case dogType:

		if dogWeightMin > a.weight {
			err = fmt.Errorf("improper weight: minimal dog's weight is %v: %s's weight is %v", dogWeightMin, a.name, a.weight)
		} else if a.weight > dogWeightMax {
			err = fmt.Errorf("improper weight: maximal dog's weight is %v: %s's weight is %v", dogWeightMax, a.name, a.weight)
		}

	case catType:

		if catWeightMin > a.weight {
			err = fmt.Errorf("improper weight: minimal cat's weight is %v: %s's weight is %v", catWeightMin, a.name, a.weight)
		} else if a.weight > catWeightMax {
			err = fmt.Errorf("improper weight: maximal cat's weight is %v: %s's weight is %v", catWeightMax, a.name, a.weight)
		}

	case cowType:

		if cowWeightMin > a.weight {
			err = fmt.Errorf("improper weight: minimal cow's weight is %v: %s's weight is %v", cowWeightMin, a.name, a.weight)
		} else if a.weight > cowWeightMax {
			err = fmt.Errorf("improper weight: maximal cow's weight is %v: %s's weight is %v", cowWeightMax, a.name, a.weight)
		}

	}

	if err == nil { //if there are still no errors - check if an animal is edible.
		err = a.isEdible()
	}

	return err
}

//check if this animal is edible.
func (a animal) isEdible() (err error) {
	if a.animalType == dogType || a.animalType == catType {
		err = fmt.Errorf("%[1]s is not edible: %[1]s is a %s", a.name, a.animalType) //an animal isn't edible if it's a dog or a cat.
	}
	return err
}