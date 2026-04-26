/*
Five philosophers sit around a table with five forks.
Each philosopher needs the two neighboring forks to eat.
Design a synchronization strategy that allows them to eat
without deadlock or starvation.

F0 | P0 | F1 | P1 | F2 | P2 | F3 | P3 | F4 | P4

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type DiningTable struct {
	numOfPhilosophers int
	forks             []sync.Mutex
}

func NewDiningTable(count int) *DiningTable {
	return &DiningTable{
		numOfPhilosophers: count,
		forks:             make([]sync.Mutex, count),
	}
}

func (d *DiningTable) Run() {

	var wg sync.WaitGroup

	for p := range d.numOfPhilosophers {

		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			d.philosopherStartEating(p)
		}(p)

	}

	wg.Wait()
	fmt.Println("All philosophers end eating")
}

func (d *DiningTable) philosopherStartEating(philosopherID int) {

	leftFork := philosopherID
	rightFork := (philosopherID + 1) % d.numOfPhilosophers

	fmt.Printf("P%d wants Left Fork\n", philosopherID)
	d.pickUpFork(leftFork)
	fmt.Printf("P%d got Left Fork\n", philosopherID)

	time.Sleep(100 * time.Millisecond) // force everyone to hold left fork first

	fmt.Printf("P%d wants Right Fork\n", philosopherID)
	d.pickUpFork(rightFork)
	fmt.Printf("P%d got Right Fork\n", philosopherID)

	d.startEating(philosopherID)

	d.returnFork(leftFork)
	fmt.Printf("P%d Released Left Fork\n", philosopherID)
	d.returnFork(rightFork)
	fmt.Printf("P%d Released Right Fork\n", philosopherID)
}

func (d *DiningTable) pickUpFork(fork int) {

	d.forks[fork].Lock()

}

func (d *DiningTable) returnFork(fork int) {

	d.forks[fork].Unlock()

}

func (d *DiningTable) startEating(philosopherID int) {

	fmt.Printf("Philosopher %d is eating now.\n", philosopherID)
	time.Sleep(100 * time.Second)

}

func main() {

	d := NewDiningTable(5)
	d.Run()
}
