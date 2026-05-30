package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return nil
	}

	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	if size == 1 {
		return []int{generator.Int()}
	}

	array := make([]int, size)
	for i := range array {
		array[i] = generator.Int()
	}
	return array
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if data == nil {
		return 0
	}

	max := data[0]
	for i := range data {
		if max < data[i] {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	if chunkSize < 2 {
		return maximum(data)
	}

	arrMax := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		start := i * chunkSize

		var end int
		if i == CHUNKS-1 {
			remainder := len(data) % CHUNKS
			end = start + chunkSize + remainder
		} else {
			end = start + chunkSize
		}

		go func(slice []int) {
			defer wg.Done()

			arrMax[i] = maximum(slice)
		}(data[start:end])
	}

	wg.Wait()
	return maximum(arrMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	genArr := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(genArr)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(genArr)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(genArr)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
