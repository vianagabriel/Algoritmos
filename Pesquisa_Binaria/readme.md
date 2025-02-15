# Pesquisa Binária em Go

Este repositório contém uma implementação do algoritmo de **Pesquisa Binária** em Go. A pesquisa binária é uma técnica eficiente para encontrar um valor em uma lista ordenada, dividindo repetidamente o intervalo de pesquisa ao meio até encontrar o valor desejado.

## Índice

- [Descrição](#descrição)
- [Como funciona](#como-funciona)
- [Complexidade](#complexidade)
- [Exemplo de código](#exemplo-de-código)


## Descrição

A Pesquisa Binária reduz significativamente o número de comparações necessárias para encontrar um elemento em uma lista ordenada. Ao contrário da busca linear, que percorre toda a lista, a pesquisa binária divide a lista ao meio em cada iteração, tornando-se muito mais eficiente.

## Como funciona

1. **Passo 1**: Defina os índices `inicial` e `final`. No início, `inicial` é 0 e `final` é o índice do último elemento da lista.
2. **Passo 2**: Calcule o ponto médio `meio` da lista.
3. **Passo 3**: Compare o valor desejado com o valor no ponto médio:
    - Se o valor desejado for igual ao valor do ponto médio, o índice é retornado.
    - Se o valor desejado for menor, a busca continua na metade inferior da lista.
    - Se o valor desejado for maior, a busca continua na metade superior.
4. **Passo 4**: Repita os passos 2 e 3 até encontrar o valor ou até os índices `inicial` e `final` se cruzarem (o que indica que o valor não está presente).

## Complexidade

- **Tempo**: A complexidade de tempo da pesquisa binária é **O(log n)**, onde `n` é o número de elementos na lista. Isso significa que o tempo de execução cresce logaritmicamente com o aumento do tamanho da lista, tornando o algoritmo eficiente para listas grandes.


## Exemplo de código

Aqui está o trecho central da implementação em Go:

```go
func PesquisaBin(numeros []int, numAlvo int) (int, error) {
	inicio, final := 0, len(numeros)-1

	for inicio <= final {
		meio := (inicio + final) / 2
		if numeros[meio] == numAlvo {
			return meio, nil
		} else if numeros[meio] < numAlvo {
			inicio = meio + 1
		} else {
			final = meio - 1
		}
	}

	return -1, errors.New("numero não encontrado")
}
