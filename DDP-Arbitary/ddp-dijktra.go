package main

import (
	"fmt"
	"time"
)

const numGarfos = 5

func filosofo(id int, garcom chan int) {
	for i := 0; i < 3; i++ {
		// Tenta pegar dois garfos
		fmt.Printf("Filósofo %d pediu permissao p/ pegar 2 garfos...\n", id)
		garcom <- 1 
		garcom <- 1

		fmt.Printf("Garcom cedeu filosofo %d  p/ pegar os 2 garfos.\n", id)
		fmt.Printf("Filósofo %d está comendo.\n", id)
		time.Sleep(time.Second * 2) 

		// Devolve os dois garfos
		<-garcom 
		<-garcom 
		
		fmt.Printf("Filósofo %d terminou de comer e devolveu os garfos.\n", id)
	

		// O filósofo volta a pensar 
		time.Sleep(time.Second * 2) 
	}
}

func main() {

	garcom := make(chan int, numGarfos)


	for i := 1; i <= 5; i++ {
	 filosofo(i, garcom)
	}

	// Evita que o programa principal finalize antes das goroutines
	select {}
}
