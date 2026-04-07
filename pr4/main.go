package main

import (
	"fmt"
	"math"
)

// Функція термічної перевірки
func checkThermal(ik3A, t, allowed float64) string {
	if t <= 0 || allowed <= 0 {
		return "Термічна перевірка: некоректні дані"
	}

	actual := ik3A * ik3A * t

	if actual <= allowed {
		return fmt.Sprintf("Термічно: Стійкий (I²t = %.2e ≤ %.2e)", actual, allowed)
	}
	return fmt.Sprintf("Термічно: НЕ стійкий (I²t = %.2e > %.2e)", actual, allowed)
}

func main() {
	var voltageKv, impedance, timeSec, allowedI2t float64

	fmt.Println("=== Калькулятор струмів короткого замикання ===")
	fmt.Println("Усі розрахунки виконуються в СІ (В, А, Ом)\n")

	// Ввід з перевіркою
	fmt.Print("Лінійна напруга U_LL (кВ): ")
	if _, err := fmt.Scan(&voltageKv); err != nil {
		fmt.Println("Помилка введення!")
		return
	}

	fmt.Print("Сумарний опір Z (Ом): ")
	if _, err := fmt.Scan(&impedance); err != nil {
		fmt.Println("Помилка введення!")
		return
	}

	fmt.Print("Час t (с): ")
	if _, err := fmt.Scan(&timeSec); err != nil {
		fmt.Println("Помилка введення!")
		return
	}

	fmt.Print("Дозволене I²·t (A²·s): ")
	if _, err := fmt.Scan(&allowedI2t); err != nil {
		fmt.Println("Помилка введення!")
		return
	}

	// Перевірка базових значень
	if voltageKv <= 0 || impedance <= 0 {
		fmt.Println("Помилка: U_LL і Z повинні бути > 0")
		return
	}

	// Переведення в В
	uV := voltageKv * 1000.0

	// 1. Трифазний КЗ
	ik3A := uV / (math.Sqrt(3.0) * impedance)
	ik3kA := ik3A / 1000.0

	// 2. Однофазний КЗ
	ik1A := (uV / math.Sqrt(3.0)) / impedance
	ik1kA := ik1A / 1000.0

	// 3. Термічна стійкість
	thermalResult := checkThermal(ik3A, timeSec, allowedI2t)

	// 4. Динамічна стійкість
	kDyn := 1.8
	iDynA := ik3A * kDyn
	iDynKA := iDynA / 1000.0

	// Вивід
	fmt.Println("\n=== РЕЗУЛЬТАТИ ===")
	fmt.Printf("I_k(3φ) = %.2f A (%.3f kA)\n", ik3A, ik3kA)
	fmt.Printf("I_k(1φ) = %.2f A (%.3f kA)\n", ik1A, ik1kA)
	fmt.Println(thermalResult)
	fmt.Printf("Динамічна дія: I_dyn = %.2f A (%.2f kA), k = %.1f\n", iDynA, iDynKA, kDyn)
}
