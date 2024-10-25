package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {

		fmt.Print("숫자 입력 : ")
		in := bufio.NewReader(os.Stdin)
		i, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i) // 줄바꿈등 제거. 파이썬 strip 함수와 비슷
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}
}
