package main

import (
	"fmt"
	"nu/utils"
)

/*
 * Complete the 'formacaoDeTime' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY pontuacao
 *  2. INTEGER tamanhoDoTime
 *  3. INTEGER k
 */

const (
	left = iota
	right
)

func formacaoDeTime(pontuacao []int32, tamanhoDoTime int32, k int32) int64 {
	var result int64
	var escolhido int32
	var lorR int32

	for i := int32(0); i < tamanhoDoTime; i++ {
		if int32(len(pontuacao)) >= k {
			escolhido, lorR = pegarMaiorLeftRight(pontuacao[:k], pontuacao[int32(len(pontuacao))-k:])
		} else {
			quickSort(&pontuacao, 0, len(pontuacao)-1)
			escolhido = pegarMaior(pontuacao)
			lorR = left
		}

		pontuacao = removeElement(pontuacao, escolhido, lorR)
		result += int64(escolhido)
	}

	return result
}

func quickSort(arr *[]int32, menor, maior int) {

	if menor < maior {
		index := dividir(arr, menor, maior)

		quickSort(arr, menor, index-1)

		quickSort(arr, index+1, maior)
	}
}

func dividir(arr *[]int32, menor, maior int) int {
	pivot := (*arr)[maior]

	i := menor - 1
	for j := menor; j <= maior-1; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[maior] = (*arr)[maior], (*arr)[i+1]

	return i + 1
}

func pegarMaior(arr []int32) int32 {
	maior := int32(0)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= maior {
			maior = arr[i]
		}
	}

	return maior
}

func pegarMaiorLeftRight(xLeft, xRight []int32) (int32, int32) {
	maior := int32(0)
	lorR := int32(0)

	for i := 0; i < len(xLeft); i++ {
		if xLeft[i] >= maior {
			maior = xLeft[i]
			lorR = left
		}

		if xRight[i] > maior {
			maior = xRight[i]
			lorR = right
		}
	}

	return maior, lorR
}

func removeElement(arr []int32, element, lorR int32) []int32 {
	index := 0

	if lorR == 0 {
		for key, value := range arr {
			if value == element {
				index = key
				break
			}
		}
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] == element {
				index = i
				break
			}
		}
	}

	return append(arr[:index], arr[index+1:]...)
}

func main() {
	pontuacao := utils.GetPontuacao3()
	tamanhoDoTime := utils.GetTamanhoTime3()
	k := utils.GetK3()

	result := formacaoDeTime(pontuacao, tamanhoDoTime, k)

	fmt.Println(result)
}
