package backup

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formacaoDeTime' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY pontuacao
 *  2. INTEGER tamanho_do_time
 *  3. INTEGER k
 */

func formacaoDeTime(pontuacao []int32, tamanho_do_time int32, k int32) int64 {
	var result int64
	var escolhido int32
	var lorR int32

	time := make([]int32, tamanho_do_time)

	// fmt.Println("--------------------")
	// fmt.Println("        Inicio      ")
	// fmt.Println("--------------------")
	// fmt.Println(pontuacao)
	// fmt.Println(tamanho_do_time)
	// fmt.Println(k)

	// fmt.Println("--------------------")
	// fmt.Println("         Loop       ")
	// fmt.Println("--------------------")

	for i := int32(0); i < tamanho_do_time; i++ {
		// fmt.Println("--------------------")

		// fmt.Println(pontuacao)
		if int32(len(pontuacao)) >= k {
			escolhido, lorR = pegarMaior(pontuacao[:k], pontuacao[int32(len(pontuacao))-k:])
			// fmt.Println(escolhido)
			// fmt.Println("len", len(pontuacao))

			pontuacao = removeElement(pontuacao, escolhido, lorR)
			// fmt.Println(pontuacao)

			time[i] = escolhido
			// fmt.Println(time)

			result += int64(escolhido)
		} else {
			// escolhido = pegarMaior(pontuacao[:k], pontuacao[int32(len(pontuacao)):])
			// fmt.Println(escolhido)
			// fmt.Println("len", len(pontuacao))
			for _, v := range pontuacao {
				result += int64(v)
			}

			return result
		}

		// if int32(len(pontuacao)) <= k {
		//     for _, v := range pontuacao{
		//         result += int64(v)
		//     }
		// } else {
		// pontuacao = removeElement(pontuacao, escolhido)
		// fmt.Println(pontuacao)

		// time[i] = escolhido
		// fmt.Println(time)

		// result += int64(escolhido)
		// }
	}

	// result += int64(escolhido)

	// fmt.Println(pontuacao[:k-1])
	// fmt.Println(pontuacao[pl - k + 1:])

	// fmt.Println("--------------------")
	// fmt.Println("       Result       ")
	// fmt.Println(escolhido)
	// fmt.Println("--------------------")
	// fmt.Println(removeElement(pontuacao, escolhido))

	return result
}

func pegarMaior(xLeft, xRight []int32) (int32, int32) {
	maior := int32(0)

	lorR := int32(0)

	const (
		left = iota
		right
	)

	for k, v := range xLeft {
		// fmt.Println("--------------------")
		// // fmt.Println(k, v)
		// fmt.Println(xLeft, xRight)
		// fmt.Println(xLeft[k], xRight[k])

		if xLeft[k] >= maior {
			maior = int32(v)
			lorR = left
		}

		if xRight[k] > maior {
			maior = int32(xRight[k])
			lorR = right
		}
		// fmt.Println(maior)
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

func ab() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	pontuacaoCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var pontuacao []int32

	for i := 0; i < int(pontuacaoCount); i++ {
		pontuacaoItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		pontuacaoItem := int32(pontuacaoItemTemp)
		pontuacao = append(pontuacao, pontuacaoItem)
	}

	tamanho_do_timeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tamanho_do_time := int32(tamanho_do_timeTemp)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := formacaoDeTime(pontuacao, tamanho_do_time, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
