package main

import "fmt"

func main() {
	array := []int{8, 4, 2, 5, 3}
	fmt.Println(array)
	Ordenacao(array)
	fmt.Println(array)

}

func Ordenacao(arr []int) {
	for i := 0; i < len(arr); i++ {
		posAnterior := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[posAnterior] {
				posAnterior = j
			}
		}

		temp := arr[i]
		arr[i] = arr[posAnterior]
		arr[posAnterior] = temp

	}
}
