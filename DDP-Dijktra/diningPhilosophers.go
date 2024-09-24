package main

import (
    "fmt"
    "time"
)

const timesToEat = 3

// Philosopher 
type Philosopher struct {   
    id              int
    leftFork, rightFork chan bool
}

// run 
func (p *Philosopher) run(done chan<- bool) {
    for i := 0; i < timesToEat; i++ {
        fmt.Printf("Philosopher %d is thinking.\n", p.id)
        time.Sleep(time.Second)

        <-p.leftFork 
        <-p.rightFork 
        fmt.Printf("Philosopher %d is eating.\n", p.id)
        time.Sleep(time.Second)

        p.leftFork <- true 
        p.rightFork <- true 
    }
    done <- true 
}

func main() {
    const numPhilosophers = 5

    startTime := time.Now() 

    forks := make([]chan bool, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        forks[i] = make(chan bool, 1)
        forks[i] <- true 
    }

    done := make(chan bool, numPhilosophers) 
    philosophers := make([]Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = Philosopher{
            id: i + 1,
            leftFork: forks[i],
            rightFork: forks[(i+1)%numPhilosophers],
        }
        go philosophers[i].run(done)
    }

    // Wait for all philosophers
    for i := 0; i < numPhilosophers; i++ {
        <-done
    }

    executionTime := time.Since(startTime) 
    fmt.Printf("Total execution time: %s\n", executionTime)
}
