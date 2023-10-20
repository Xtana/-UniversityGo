package main

import "fmt"

func main() {
	errorMessage := "Введено неверное число итераций"
	iterationsNumber := 100

	if сheckingIterationInput(iterationsNumber) {
		piWallis := piCalculationByWallis(iterationsNumber)
		piLeibniz := piCalculationByLeibniz(iterationsNumber)
		printPi(iterationsNumber, piWallis, piLeibniz)
	} else {
		fmt.Print(errorMessage)
	}
}

func printPi(iterationsNumber int, piWallis float64, piLeibniz float64) {
	fmt.Printf("Число итераций: %d\n"+
		"Число ПИ по формуле Валлиса - %.5f\n"+
		"Число ПИ по формуле Лейбница - %.5f\n", iterationsNumber, piWallis, piLeibniz)
}

func сheckingIterationInput(iterationsNumber int) bool {
	if iterationsNumber <= 0 {
		return false
	} else {
		return true
	}
}

func piCalculationByWallis(iterationsNumber int) float64 {
	pi := 1.0
	numerator := 0.0
	denominator := 1.0
	for i := 0; i < iterationsNumber; i++ {
		if i%2 == 0 {
			numerator += 2
		} else {
			denominator += 2
		}
		pi *= (numerator / denominator)
	}
	return pi * 2
}

func piCalculationByLeibniz(iterationsNumber int) float64 {
	pi := 0.0
	numerator := 1.0
	denominator := -1.0
	for i := 0; i < iterationsNumber; i++ {
		denominator += 2
		if i%2 == 0 {
			pi += (numerator / denominator)
		} else {
			pi -= (numerator / denominator)
		}
	}
	return pi * 4
}
