package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
)

func main() {
	place := "лугу"
	action := "пасется"
	animal := "коров"
	numberOfAnimal := rand.Intn(100)
	resString := "На " + place + " " + action + " " + strconv.Itoa(numberOfAnimal) + " " + animal

	ending := getEndingBySwitchCase(numberOfAnimal)
	fmt.Println(resString + ending)
}

func getEndingBySwitchCase(numberOfAnimal int) string {
	sliceBetweenTowFore := []int{2, 3, 4}
	switch {
	case slices.Index(sliceBetweenTowFore, numberOfAnimal%10) != -1 && (numberOfAnimal <= 11 || numberOfAnimal >= 20):
		return "ы"
	case numberOfAnimal%10 == 1 && numberOfAnimal != 11:
		return "а"
	default:
		return ""
	}
}

func getEndingByIfElse(numberOfAnimal int) string {
	if numberOfAnimal%10 == 1 && numberOfAnimal != 11 {
		return "а"
	} else if numberOfAnimal%10 >= 2 && numberOfAnimal%10 <= 4 && numberOfAnimal/10 != 1 {
		return "ы"
	} else {
		return ""
	}
}
