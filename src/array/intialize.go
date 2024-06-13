// Package array provides functions for creating and manipulating NDArrays.
package array

import (
	"fmt"
	"gopi/utils"
	"math"
	"runtime"
	"sync"
)

// Zeroes creates a new NDArray with the given shape, filled with zeroes.
func Zeroes(shape ...int) *NDArray {
	array := NewNDArray(shape...)
	return array
}

// Fill creates a new NDArray with the given shape, filled with the provided value.
func Fill(value float64, shape ...int) *NDArray {
	arr := NewNDArray(shape...)
	n := len(arr.data)
	numWorkers := runtime.NumCPU()
	chukSize := (n + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		fmt.Println("Starting worker", i)
		go func(workerId int) {
			defer wg.Done()

			start := workerId * chukSize
			end := start + chukSize

			if end > n {
				end = n
			}

			for j := start; j < end; j++ {
				arr.data[j] = value
			}

			fmt.Println("Worker", workerId, "done")
		}(i)
	}

	wg.Wait()
	return arr
}

// Ones creates a new NDArray with the given shape, filled with ones.
func Ones(shape ...int) *NDArray {
	return Fill(1, shape...)
}

// Eye creates a new square NDArray with the given number of rows (and columns),
// with ones on the diagonal and zeroes elsewhere.
func Eye(rows int) *NDArray {
	arr := NewNDArray(rows, rows)
	for i := 0; i < rows; i++ {
		arr.data[i*rows+i] = 1
	}
	return arr
}

// EyeWithCols creates a new NDArray with the given number of rows and columns,
// with ones on the diagonal and zeroes elsewhere.
func EyeWithCols(rows int, cols int) *NDArray {
	arr := NewNDArray(rows, cols)
	min := utils.Min(rows, cols)

	for i := 0; i < min; i++ {
		arr.data[i*cols+i] = 1
	}

	return arr
}

// Identity creates a new square NDArray with the given size,
// with ones on the diagonal and zeroes elsewhere. It's equivalent to Eye(size).
func Identity(size int) *NDArray {
	return Eye(size)
}

func Arrange(start float64, stop float64, step float64) *NDArray {
	if step == 0 {
		panic("Step cannot be zero")
	}
	if start == stop {
		panic("Start and stop cannot be equal")
	}

	if start > stop && step > 0 {
		step = -step
	} else if start < stop && step < 0 {
		step = -step
	}

	size := int(math.Ceil(math.Abs(stop-start))/math.Abs(step)) + 1

	arr := NewNDArray(size)
	numWorkers := runtime.NumCPU()
	chukSize := (size + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			defer wg.Done()

			startIdx := workerId * chukSize
			endIdx := startIdx + chukSize

			if endIdx > size {
				endIdx = size
			}

			value := start + float64(startIdx)*step

			for j := startIdx; j < endIdx; j++ {
				arr.data[j] = value
				value += step
			}
		}(i)
	}

	wg.Wait()

	return arr
}
