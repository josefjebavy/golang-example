package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulace dlouhotrvajícího úkolu
func longRunningTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulace náhodné doby zpracování mezi 1-3 sekundami
	duration := time.Duration(rand.Intn(3)+1) * time.Second
	fmt.Printf("Task %d started, will take %v\n", id, duration)

	time.Sleep(duration) // Simulace zpracování
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Synchronizační skupina pro čekání na dokončení všech gorutin
	var wg sync.WaitGroup

	// Počet paralelních úkolů
	numTasks := 5

	// Spuštění úkolů paralelně pomocí gorutin
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go longRunningTask(i, &wg)
	}

	// Čekání na dokončení všech úkolů
	wg.Wait()

	fmt.Println("All tasks completed")
}

