package main

import "fmt"

type ToyFactory interface {
	CreateToy() Toy
}

type Toy interface {
	Play()
}

type StuffedAnimalFactory struct{}

func (saf StuffedAnimalFactory) CreateToy() Toy {
	return &StuffedAnimal{}
}

type StuffedAnimal struct{}

func (sa StuffedAnimal) Play() {
	fmt.Println("Hug me! I am stuffed animal.")
}

type ActionFigureFactory struct{}

func (aff ActionFigureFactory) CreateToy() Toy {
	return &ActionFigure{}
}

type ActionFigure struct{}

func (af ActionFigure) Play() {
	fmt.Println("I amm an action figure. Lets go on an adventure!")
}

func main() {
	stuffedAnimalFactory := StuffedAnimalFactory{}
	stuffedAnimalToy := stuffedAnimalFactory.CreateToy()
	stuffedAnimalToy.Play()

	actionFigureFactory := ActionFigureFactory{}
	actionFigureToy := actionFigureFactory.CreateToy()
	actionFigureToy.Play()
}
