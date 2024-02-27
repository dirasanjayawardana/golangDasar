package basic

// interface --> sebagai kerangka method-method
type IAnimal interface{
	Sound()
	NumberOfLegs()
}

func Animal(animal IAnimal){
	animal.Sound()
	animal.NumberOfLegs()
}