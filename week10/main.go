package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input your number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true

	// bug fix
	if n <= 1 { //A prime number is a natural number greater than 1 that is not a product of two smaller natural numbers.
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false
			}
			j++
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is NOT prime number", n)
	}
}
