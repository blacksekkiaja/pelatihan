package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) EstimatedDistace() float64 {
	costPerMil := 1.5
	count := car.FuelIn / costPerMil

	return count
}

func main() {
	Cars := Car{
		Type:   "Super Car",
		FuelIn: 20,
	}

	EstimatedDistance := Cars.EstimatedDistace()
	fmt.Printf("Perkiraan jarak tempuh mobil %s adalah : %.2f /mil.\n", Cars.Type, EstimatedDistance)
}
