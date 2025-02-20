# Ordenação por Seleção

A ordenação por seleção (Selection Sort) é um algoritmo de ordenação simples e ineficiente para grandes volumes de dados. Ele funciona dividindo o array em duas partes: a parte ordenada e a parte não ordenada. Em cada iteração, o menor elemento da parte não ordenada é selecionado e trocado com o primeiro elemento da parte não ordenada, expandindo assim a parte ordenada.

## Como funciona

1. Percorre a lista e encontra o menor elemento.
2. Troca esse menor elemento com o primeiro elemento da parte não ordenada.
3. Move o índice inicial da parte não ordenada uma posição à frente.
4. Repete esse processo até que toda a lista esteja ordenada.

## Implementação em Go

Abaixo está a implementação do algoritmo de ordenação por seleção em Go:

```go
package main

import "fmt"

func main() {
	array := []int{8, 4, 2, 5, 3}
	fmt.Println("Array antes da ordenação:", array)
	Ordenacao(array)
	fmt.Println("Array depois da ordenação:", array)
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
```

## Explicação do Código

1. O array inicial `{8, 4, 2, 5, 3}` é impresso antes da ordenação.
2. A função `Ordenacao` percorre o array, encontrando o menor valor e trocando-o com a posição inicial da parte não ordenada.
3. Após a execução da função, o array estará ordenado e é impresso novamente.

## Complexidade

- **Melhor caso (array já ordenado):** O(n²)
- **Pior caso (array em ordem reversa):** O(n²)
- **Caso médio:** O(n²)

Devido à sua complexidade quadrática, o Selection Sort não é eficiente para grandes conjuntos de dados, mas pode ser útil para pequenas listas devido à sua simplicidade e estabilidade.

