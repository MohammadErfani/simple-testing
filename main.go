package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneChan := make(chan bool)
	defer close(doneChan)
	go readUserInput(os.Stdin, doneChan)
	<-doneChan
	fmt.Println("Goodbye")
}
func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("--------------")
	fmt.Println("Enter a number, and we'll tell you if it is a prime number. print q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(reader io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(reader)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a number", false
	}
	_, msg := isPrime(numToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	if n < 0 {
		return false, "Negative numbers are not prime"
	}
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime becuase it id divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}
