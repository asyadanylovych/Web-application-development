package main

import (
	"fmt"
	"math"
)

// ProfitResult зберігає результати розрахунків
type ProfitResult struct {
	Sigma           float64
	Probability     float64
	EnergyNoPenalty float64
	EnergyPenalty   float64
	ProfitNoPenalty float64
	Penalty         float64
	TotalProfit     float64
}

// pdf обчислює значення функції щільності ймовірності нормального розподілу
func pdf(x, mean, sigma float64) float64 {
	return (1.0 / (sigma * math.Sqrt(2.0*math.Pi))) * math.Exp(-0.5*math.Pow((x-mean)/sigma, 2.0))
}

// integrateNormal виконує чисельне інтегрування методом трапецій
func integrateNormal(mean, sigma, a, b float64, nSteps int) float64 {
	h := (b - a) / float64(nSteps)
	sum := 0.5 * (pdf(a, mean, sigma) + pdf(b, mean, sigma))
	
	for i := 1; i < nSteps; i++ {
		sum += pdf(a+float64(i)*h, mean, sigma)
	}
	return sum * h
}

// calculateProfitForSigma розраховує фінансові показники для заданої похибки
func calculateProfitForSigma(Pc, sigma, B float64) ProfitResult {
	hours := 24.0
	lowerBound := Pc - 0.25 
	upperBound := Pc + 0.25
	nSteps := 10000

	// Ймовірність потрапляння в діапазон без штрафів
	probability := integrateNormal(Pc, sigma, lowerBound, upperBound, nSteps)

	// Розрахунок енергії (МВт-год)
	energyNoPenalty := Pc * hours * probability
	energyPenalty := Pc * hours * (1.0 - probability)

	// Розрахунок грошових показників (тис. грн)
	profitNoPenalty := energyNoPenalty * B
	penalty := energyPenalty * B
	totalProfit := profitNoPenalty - penalty

	return ProfitResult{
		Sigma:           sigma,
		Probability:     probability,
		EnergyNoPenalty: energyNoPenalty,
		EnergyPenalty:   energyPenalty,
		ProfitNoPenalty: profitNoPenalty,
		Penalty:         penalty,
		TotalProfit:     totalProfit,
	}
}

func main() {
	// Вхідні дані згідно з контрольним прикладом [cite: 502, 503, 506]
	Pc := 5.0
	sigma1 := 1.0
	sigma2 := 0.25
	B := 7.0

	fmt.Println("Розрахунок ефекту від вдосконалення системи прогнозування")
	fmt.Printf("Вхідні дані:\n• Pc = %.1f МВт\n• σ₁ = %.1f МВт\n• σ₂ = %.2f МВт\n• B = %.1f грн/кВт·год\n\n", 
		Pc, sigma1, sigma2, B)

	res1 := calculateProfitForSigma(Pc, sigma1, B)
	res2 := calculateProfitForSigma(Pc, sigma2, B)

	printResult("ДО ВДОСКОНАЛЕННЯ", res1, Pc)
	printResult("ПІСЛЯ ВДОСКОНАЛЕННЯ", res2, Pc)

	improvement := res2.TotalProfit - res1.TotalProfit
	fmt.Printf("\n>>> ПРИБУТОК ВІД ВДОСКОНАЛЕННЯ СИСТЕМИ: %.1f тис. грн\n", improvement)
}

func printResult(title string, res ProfitResult, Pc float64) {
	fmt.Printf("--- %s (σ = %.2f МВт) ---\n", title, res.Sigma)
	fmt.Printf("Ймовірність у межах [%.2f; %.2f]: %.1f%%\n", Pc-0.25, Pc+0.25, res.Probability*100)
	fmt.Printf("Енергія без штрафу: %.1f МВт-год\n", res.EnergyNoPenalty)
	fmt.Printf("Енергія зі штрафом: %.1f МВт-год\n", res.EnergyPenalty)
	fmt.Printf("Прибуток без штрафу: %.1f тис. грн\n", res.ProfitNoPenalty)
	fmt.Printf("Штраф за небаланс: %.1f тис. грн\n", res.Penalty)
	fmt.Printf("ЗАГАЛЬНИЙ ПРИБУТОК: %.1f тис. грн\n\n", res.TotalProfit)
}