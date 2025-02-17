package main

import (
	"fmt"
)

func main() {
	array := []int{102, 23, 67, 74, 22, 5, 10, 43, 2, 1}

	fmt.Println(array)
	OrdenacaoSelecao(array)
	fmt.Println(array)
}

func OrdenacaoSelecao(array []int) {
	for i := 0; i < len(array); i++ {
		numAnterior := i

		for j := numAnterior + 1; j < len(array); j++ {
			if array[j] < array[numAnterior] {
				numAnterior = j
			}
		}
		Temp := array[i]
		array[i] = array[numAnterior]
		array[numAnterior] = Temp
	}

}
