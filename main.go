package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
// ВАРИАНТ 1
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size < 2 {
		return nil
	}

	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		sliceInt      = make([]int, 0, size)
		numGoroutines int
		capacity      int
	)
	//Заполняем слайс одной горутиной если размер меньше CHUNK
	if size < CHUNKS {
		numGoroutines = 1
		capacity = size
	} else {
		numGoroutines = CHUNKS
		capacity = size / numGoroutines
	}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			addCap := 0
			//Добавляем остаток к последней горутине
			if id == numGoroutines-1 {
				addCap = size % numGoroutines
			}

			tempBuf := make([]int, 0, capacity+addCap)
			for len(tempBuf) != cap(tempBuf) {
				tempBuf = append(tempBuf, rand.Int())
			}

			mu.Lock()
			sliceInt = append(sliceInt, tempBuf...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return sliceInt
}

/* Вариант 2 с использованием канала, остаток распределяется хаотично, а не в последнюю горутину + создаем горутины динамически
func generateRandomElements(size int) []int {
	// ваш код здесь

	if size < 2 {
		return nil
	}

	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		arrInt        = make([]int, 0, size)
		numGoroutines int
		capacity      int
	)

	//Динамическое создание горутин в зависимости от степени 2 с учетом того, что одна горутина может взять на себя минимум 10к значений.
	numGoroutines = int(math.Log2(float64(size) / 10000))
	if numGoroutines < 1 {
		numGoroutines = 1
		capacity = size
	} else {
		capacity = size / numGoroutines
	}

	wg.Add(numGoroutines + 1)

	ch := make(chan int, size%numGoroutines)
	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 0; i < size%numGoroutines; i++ {
			ch <- 1
		}
	}()

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			addCap := 0
			for range ch {
				addCap++
			}
			tempBuf := make([]int, 0, capacity+addCap)
			for len(tempBuf) != cap(tempBuf) {
				tempBuf = append(tempBuf, rand.Int())
			}
			mu.Lock()
			arrInt = append(arrInt, tempBuf...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return arrInt
}
*/ //

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	switch {
	case data == nil:
		return -1
	case len(data) < 2:
		return -1
	case len(data) != cap(data):
		return -1
	case slices.Min(data) < 0:
		return -1
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var (
		sliceChunks = make([][]int, CHUNKS)
		maxValues   [CHUNKS]int64
		chunkLen    = len(data) / CHUNKS
		wg          sync.WaitGroup
	)

	for i := 0; i < CHUNKS; i++ {
		if i == CHUNKS-1 {
			sliceChunks[i] = data[i*chunkLen:]
			break
		}
		sliceChunks[i] = data[i*chunkLen : (i+1)*chunkLen]
	}

	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(i int) {
			defer wg.Done()
			max := slices.Max(sliceChunks[i])
			atomic.StoreInt64(&maxValues[i], int64(max)) //Нужен ли вообще здесь мьютекс или атомарная функция? Мы пишем ведь в уникальный индекс
			//maxValues[i] = int64(max)
		}(i)
	}

	wg.Wait()

	return int(slices.Max(maxValues[0:]))
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	sliceInt := generateRandomElements(SIZE)
	if sliceInt == nil {
		log.Println("size cant be less than 2")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(sliceInt)
	finish := time.Since(start)
	elapsed := int(finish.Microseconds())
	if max == -1 {
		log.Println("incorrect slice")
		return
	}
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(sliceInt)
	finish = time.Since(start)
	elapsed = int(finish.Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
