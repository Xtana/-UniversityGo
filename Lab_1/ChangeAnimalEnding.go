package main

import (
	"fmt"
	"math/rand"
)

func main() {
	place := "лугу"
	action := "пасется"
	animal := "коров"
	randNum := rand.Intn(100)
	changeAnimalEnding(randNum, place, action, animal)
}

func changeAnimalEnding(numberOfCow int, place string, action string, animal string) {
	if numberOfCow%10 == 1 && numberOfCow != 11 {
		fmt.Println("На "+place+" "+action, numberOfCow, animal+"а")
	} else if numberOfCow%10 >= 2 && numberOfCow%10 <= 4 && numberOfCow/10 != 1 {
		fmt.Println("На "+place+" "+action, numberOfCow, animal+"ы")
	} else {
		fmt.Println("На "+place+" "+action, numberOfCow, animal)
	}
}
