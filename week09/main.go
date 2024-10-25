package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(6)
	fmt.Println(target)
}
