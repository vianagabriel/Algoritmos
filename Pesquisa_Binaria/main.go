package main

import (
	"errors"
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 15, 16, 18, 22, 33, 57, 69, 88, 100, 112, 511, 805, 1200, 5522}

	indice, err := PesquisaBin(numeros, 444)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Número encontrado no índice:", indice)
	}
}

func PesquisaBin(numeros []int, numAlvo int) (int, error) {
	inicio := 0
	final := len(numeros) - 1

	for i := 0; i <= final; i++ {
		posMeio := (inicio + final) / 2
		palpite := numeros[posMeio]

		if palpite == numAlvo {
			return posMeio, nil
		}

		if palpite > numAlvo {
			final = posMeio - 1
		} else {
			inicio = posMeio + 1
		}
	}

	return -1, errors.New("numero não encontrado")

}
