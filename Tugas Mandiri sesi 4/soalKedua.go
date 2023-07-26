package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (student Student) Average() int {
	getTotalNilai := 0
	for _, v := range student.score {
		getTotalNilai += v
	}
	return getTotalNilai
}

func (student Student) Minimum() (int, string) {
	minScore := student.score[0]
	minIndex := 0

	for i, v := range student.score {
		if v < minScore {
			minScore = v
			minIndex = i
		}
	}
	return minScore, student.name[minIndex]
}

func (student Student) Maximum() (int, string) {
	maxScore := student.score[0]
	maxIndex := 0

	for i, v := range student.score {
		if v > maxScore {
			maxScore = v
			maxIndex = i
		}
	}
	return maxScore, student.name[maxIndex]
}
func main() {
	student := Student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umar", "Yopan"},
		score: []int{80, 75, 70, 75, 60},
	}

	average := student.Average() / int(len(student.score))
	minScore, minName := student.Minimum()
	maxScore, maxName := student.Maximum()

	fmt.Println("Average score :", average)
	fmt.Println("Min score students : ", minName, minScore)
	fmt.Println("Max score students : ", maxName, maxScore)
}
