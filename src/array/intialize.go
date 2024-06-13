// Package array provides functions for creating and manipulating NDArrays.
package array

import "gopi/utils"

// Zeroes creates a new NDArray with the given shape, filled with zeroes.
func Zeroes(shape ...int) *NDArray {
	array := NewNDArray(shape...)
	return array
}

// Fill creates a new NDArray with the given shape, filled with the provided value.
func Fill(value float64, shape ...int) *NDArray {
	arr := NewNDArray(shape...)

	for i := range arr.data {
		arr.data[i] = value
	}

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
