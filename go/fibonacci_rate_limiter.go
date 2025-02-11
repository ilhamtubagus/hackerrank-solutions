package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	a, b := 0, 1
	for {
		<-requestChan // Wait for a request

		// Rate limiting
		time.Sleep(11 * time.Millisecond)

		// Calculate next Fibonacci number
		next := (a + b) % 1000000000

		// Send the result
		resultChan <- next

		// Update a and b for the next iteration
		a, b = b, next
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skip := int32(skipTemp)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	total := int32(totalTemp)

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go ModuloFibonacciSequence(requestChan, resultChan)
	for i := int32(0); i < skip+total; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		new := <-resultChan
		if i < skip {
			continue
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
		fmt.Println(new)
	}
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
