package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: go run main.go")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	slice := []float64{}
	for scanner.Scan() {
		Strnbr := scanner.Text()
		number, err := strconv.ParseFloat(Strnbr, 64)
		if err != nil {
			fmt.Println("Error Parsing :", err)
			continue
		}
		slice = append(slice, number)
		if len(slice) > 1 {
			min, max := Guess_it(slice)
			fmt.Printf("%.0f %.0f\n", min, max)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func Average(data []float64) float64 {
	var sum float64
	length := len(data)
	var average float64
	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	average = sum / float64(length)
	return average
}

func Variance(data []float64) float64 {
	var mean float64
	var variance float64
	length := len(data)
	var sub float64
	var sq float64

	mean = Average(data)
	for _, n := range data {
		sub = n - mean
		sq += sub * sub
	}
	variance = sq / float64(length)
	return variance
}

func Standard_Deviation(data []float64) float64 {
	deviation := math.Sqrt(Variance(data))
	return deviation
}

func Guess_it(data []float64) (float64, float64) {
	var min float64
	var max float64

	start := len(data) - 4
	if start < 0 {
		start = 0
	}
	preciseData := data[start:]
	average := Average(preciseData)
	sd := Standard_Deviation(preciseData)

	min = average - sd
	max = average + sd

	return min, max
}
