package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    
    n := len(arr)
    xRatios := [3]float64{}
    
    const(
        pos = 0;
        neg = 1;
        zer = 2;
    )
    
    for i := 0; i < int(n); i++ {
        switch  {
            case arr[i] > 0:
            xRatios[pos]++
        
            case arr[i] < 0:
            xRatios[neg]++
            
            default:
            xRatios[zer]++
        }
    }
    
    for _, num := range xRatios {
        if num != 0 {
            fmt.Printf("%.6f\n", num/float64(n))
        } else {
            fmt.Printf("%.6f\n", float64(0))
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
