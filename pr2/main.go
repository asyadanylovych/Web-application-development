package main

import (
	"fmt"
	"math"
)

func calculateEmission(k float64, Q float64, B float64) float64 {
	return 0.000001 * k * Q * B
}

func controlExample() {
	fmt.Println("\n--- ПЕРЕВІРКА ---")

	bC := 1096363.0
	bO := 70945.0

	kC := 150.0
	QC := 20.47
	eC := calculateEmission(kC, QC, bC)

	kO := 0.57
	QO := 39.48
	eO := calculateEmission(kO, QO, bO)

	fmt.Printf("k вугілля: %.2f (Очікувано: 150.00)\n", kC)
	fmt.Printf("Викид вугілля: %.2f (Очікувано: 3366.00)\n", eC)
	fmt.Printf("Викид мазуту: %.2f (Очікувано: 1.60)\n", eO)

	if math.Abs(eC-3366.0) < 1 && math.Abs(eO-1.60) < 0.1 {
		fmt.Println("Статус: ПЕРЕВІРКУ ПРОЙДЕНО")
	} else {
		fmt.Println("Статус: ПОМИЛКА")
	}
}

func main() {
	fmt.Println("=== КАЛЬКУЛЯТОР ВИКИДІВ (GO) ===")

	// Дані варіанту 8
	bC := 412407.75
	bO := 175657.21
	bG := 195337.23

	// Вугілля 
	kC := (1000000.0 / 20.47) * 0.8 * (25.2 / (100.0 - 1.5)) * (1.0 - 0.985)
	eC := calculateEmission(kC, 20.47, bC)

	// Мазут
	kO := 0.57
	eO := calculateEmission(kO, 39.48, bO)

	// Газ 
	kG := 0.0
	eG := calculateEmission(kG, 33.08, bG)

	fmt.Println("\nРЕЗУЛЬТАТИ:")
	fmt.Printf("Показник вугілля: %.2f\n", kC)
	fmt.Printf("Викид вугілля: %.2f т\n", eC)
	fmt.Printf("Показник мазуту: %.2f\n", kO)
	fmt.Printf("Викид мазуту: %.2f т\n", eO)
	fmt.Printf("Газ: %.2f т\n", eG)
	fmt.Printf("РАЗОМ: %.2f т\n", eC+eO+eG)

	// Перевірка
	controlExample()
}