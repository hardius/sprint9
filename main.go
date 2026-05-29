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
		return []int{}
	}

	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	if size == 1 {
		return []int{generator.Int()}
	}

	array := make([]int, size)
	for range array {
		array = append(array, generator.Int())
	}
	return array
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
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

	arrMax := []int{}

	var wg sync.WaitGroup
	wg.Add(8)

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		go func(slice []int) {
			arrMax = append(arrMax, maximum(slice))

			wg.Done()
		}(data[start:end])
	}

	remainder := len(data) % CHUNKS
	if remainder != 0 {
		start := chunkSize * CHUNKS
		end := start + remainder
		arrMax = append(arrMax, maximum(data[start:end]))
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
