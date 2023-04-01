package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	rand.Seed(time.Now().UnixNano())
	problemsDoneEasy := []int{374, 704, 217, 242, 1, 125, 121, 70, 1929, 412, 226, 58, 104, 21, 234, 1299, 206, 13, 100, 35, 572, 101, 680, 20, 367}
	problemsDoneMedium := []int{49, 167, 98, 560, 33, 19, 567, 155, 3, 5, 198, 2}
	easy := problemsDoneEasy[rand.Intn(len(problemsDoneEasy))]
	medium := problemsDoneMedium[rand.Intn(len(problemsDoneMedium))]
	fmt.Println("Random done easy problem:", easy)
	fmt.Println("Random done medium problem:", medium)
}
