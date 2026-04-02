package main

import (
	"fmt"
	"lockfree/internal/stack"
	"sync"
)

func main() {
	s := stack.NewTreiberStack[int]()

	const numGoroutines = 100
	const opsPerGoroutine = 1000

	var wg sync.WaitGroup

	fmt.Printf("Запуск: %d горутин по %d операций каждая...\n", numGoroutines, opsPerGoroutine)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				s.Push(id*1000 + j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Все операции завершены успешно!")
}
