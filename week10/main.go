package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	if n <= 1 {
		isPrime = false
	} else if n == 2{
		isPrime = true
	} else if n % 2 == 0{ // All even numbers except 2 are not prime numbers 
		isPrime = false
	} else { //odd number
		j := 3 // start value
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break //performance up
			}
			fmt.Printf("%d ", j)
			j = j + 2 //increment
		}
	}
	
	

	if isPrime {
		fmt.Printf("%d is prime number", n)
	} else {
		fmt.Printf("%d is NOT prime number", n)
	}
}
