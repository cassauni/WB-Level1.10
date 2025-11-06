package main

import "fmt"

func main() {

	array := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10

	tempo := make(map[int][]float64)

	for _, value := range array {
		key := (int(value) / step) * step
		tempo[key] = append(tempo[key], value)
	}

	for k, value := range tempo {
		fmt.Println(k, value)
	}
}
