package array

import (
	"errors"
	"fmt"
	"gopi/utils"
	"strings"
)

type NDArray struct {
	data    []float64
	shape   []int
	strides []int
}

func NewNDArray(shape ...int) *NDArray {
	size := 1
	strides := make([]int, len(shape))

	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = size
		size *= shape[i]
	}

	return &NDArray{
		data:    make([]float64, size),
		shape:   shape,
		strides: strides,
	}
}

func (arr *NDArray) String() string {
	var builder strings.Builder

	return builder.String()
}

func (arr *NDArray) Shape() []int {
	return arr.shape
}

func (arr *NDArray) Get(indices ...int) (*NDArray, error) {
	err := validateIndices(indices, arr)

	if err != nil {
		return nil, err
	}

	// When the indices length is equal to the shape length, we are accessing a single element
	if len(indices) == len(arr.shape) {
		flatIdx, err := arr.flatIndex(indices...)
		if err != nil {
			return nil, err
		}

		return &NDArray{
			data:    []float64{arr.data[flatIdx]},
			shape:   []int{1},
			strides: []int{1},
		}, nil
	}

	newShape := arr.shape[len(indices):]
	newStrides := arr.strides[len(indices):]

	flatIdx, err := arr.flatIndex(indices...)
	if err != nil {
		return nil, err
	}

	size := 1
	for _, s := range newShape {
		size *= s
	}

	data := make([]float64, size)

	for i := 0; i < size; i++ {
		offset := 1
		subIdx := 0

		for j := len(newStrides) - 1; j >= 0; j-- {
			strideIndex := len(arr.strides) - len(newStrides) + j
			subIdx += (offset / arr.strides[strideIndex]) * arr.strides[strideIndex]
			offset *= arr.strides[strideIndex]
		}

		data[i] = arr.data[flatIdx+subIdx]
	}

	return &NDArray{
		data:    data,
		shape:   newShape,
		strides: newStrides,
	}, nil
}

func validateIndices(indices []int, arr *NDArray) error {
	indicesCnt := len(indices)
	shapeCnt := len(arr.shape)

	if indicesCnt > shapeCnt {
		return fmt.Errorf("too many indices for array: %d Expected Indices must be lower than %d", indicesCnt, shapeCnt)
	}

	length := utils.Min(indicesCnt, shapeCnt)

	for i := 0; i < length; i++ {
		index := indices[i]
		shape := arr.shape[i]

		if index < 0 || index >= shape {
			return fmt.Errorf("index out of bounds: %d Expected Index must be between 0 and %d", index, shape)
		}
	}

	return nil
}

func (arr *NDArray) flatIndex(indices ...int) (int, error) {
	if len(indices) > len(arr.shape) {
		err := fmt.Sprintf("Invalid number of indices: %d Expected Indices must be equal to %d\n", len(indices), len(arr.shape))
		return -1, errors.New(err)
	}

	flatIdx := 0

	for i, index := range indices {
		if index >= arr.shape[i] || index < 0 {
			return -1, fmt.Errorf("index out of bounds: %d Expected Index must be between 0 and %d", index, arr.shape[i])
		}

		flatIdx += index * arr.strides[i]
	}

	return flatIdx, nil
}
