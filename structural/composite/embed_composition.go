package composite

import "fmt"

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
