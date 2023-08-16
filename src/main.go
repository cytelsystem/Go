package main

import (
	"fmt"
)

func par(c chan int) {
	for {
		val := <-c
		fmt.Printf("Par: %d\n", val)
	}
}

func impar(c chan int) {
	for {
		val := <-c
		fmt.Printf("Impar: %d\n", val)
	}
}

func main() {
	paresChan := make(chan int, 10)
	imparesChan := make(chan int, 10)

	go par(paresChan)
	go impar(imparesChan)

	numeros := []int{2, 7, 10, 15, 22, 31, 42, 53}

	for _, num := range numeros {
		if num%2 == 0 {
			paresChan <- num
		} else {
			imparesChan <- num
		}
	}

	close(paresChan)
	close(imparesChan)

	// Espera indefinida para que las goroutines no terminen prematuramente
	select {}
}
