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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

const (
    H = iota 
    M
    S
)

func timeConversion(s string) string {
    xTime := strings.Split(s, ":")
    
    switch {
        case strings.Contains(s, "AM"):
        xTime[S] = strings.ReplaceAll(xTime[2], "AM", "")

        if strings.Contains(xTime[0], "12") {
            xTime[H] = "00"
        }        
        
        case strings.Contains(s, "PM"):
        xTime[S] = strings.ReplaceAll(xTime[S], "PM", "")

        if strings.Contains(xTime[0], "12") {
            break
        }
        
        iTime, err := strconv.Atoi(xTime[0])
        checkError(err)            
        xTime[H] = fmt.Sprintf("%d", iTime + 12)
    }
        
    return strings.Join(xTime, ":")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

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
