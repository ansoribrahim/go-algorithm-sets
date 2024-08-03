package handler

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBubbleSort_SortsArrayInAscendingOrder(t *testing.T) {
	arr := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	bubbleSort(arr)

	assert.Equal(t, expected, arr)
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	bubbleSort(arr)

	expected := []int{1, 2, 3, 4, 5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, arr[i])
		}
	}
}

func TestBubbleSortWithDuplicates(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 64}

	bubbleSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 64}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, arr[i])
		}
	}
}

func TestBubbleSortWithNegativeIntegers(t *testing.T) {
	arr := []int{3, -5, 2, -8, 1, -4}

	bubbleSort(arr)

	expected := []int{-8, -5, -4, 1, 2, 3}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d but got %d at index %d", expected[i], arr[i], i)
		}
	}
}

func TestBubbleSortSingleElement(t *testing.T) {
	arr := []int{5}
	bubbleSort(arr)
	assert.Equal(t, []int{5}, arr)
}

func TestBubbleSortEmptyArray(t *testing.T) {
	arr := []int{}
	bubbleSort(arr)
	// No errors should occur when sorting an empty array
}

func TestBubbleSort_AllIdenticalElements(t *testing.T) {
	arr := []int{5, 5, 5, 5, 5}

	bubbleSort(arr)

	expected := []int{5, 5, 5, 5, 5}
	assert.Equal(t, expected, arr)
}

func TestBubbleSort_MaxIntegerValues(t *testing.T) {
	// Prepare input array with maximum integer values
	arr := []int{2147483647, 2147483647, 2147483647, 2147483647, 2147483647}

	// Call the function to test
	bubbleSort(arr)

	// Validate the sorted array
	expected := []int{2147483647, 2147483647, 2147483647, 2147483647, 2147483647}
	assert.Equal(t, expected, arr)
}

func TestBubbleSort_MinimumIntegerValues(t *testing.T) {
	arr := []int{1, 0, -1}
	bubbleSort(arr)
	expected := []int{-1, 0, 1}
	for i, v := range arr {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, v)
		}
	}
}
