package main

import (
	"fmt"
	"time"
)

const numGarfos = 5

func filosofo(id int, garcom chan int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Filósofo %d pediu permissao p/ pegar garfos\n", id)
		garcom <- 1
		garcom <- 1

		fmt.Printf("Garcom cedeu filosofo %d p/ pegar os garfos.\n", id)
		fmt.Printf("Filósofo %d está comendo pela %d vez.\n", id, i+1)
		time.Sleep(time.Second * 2)

		<-garcom
		<-garcom
		fmt.Printf("Filósofo %d terminou de comer e devolveu os garfos.\n", id)

		// O filósofo volta a pensar
		time.Sleep(time.Second * 2)
	}
}

func main() {
	const numPhilosophers = 5
	garcom := make(chan int, numGarfos)

	for i := 1; i <= numPhilosophers; i++ {
		go filosofo(i, garcom)
	}

	// Evita que o programa principal finalize antes das goroutines
	select {}
}
