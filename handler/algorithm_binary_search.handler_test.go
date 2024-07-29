package handler

import (
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBinarySearch_TargetInMiddle(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 11, 13}
	target := 7

	result := binarySearch(array, target)

	expected := 3
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_TargetPresentAtBeginning(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 1

	result := binarySearch(array, target)

	assert.Equal(t, 0, result)
}

func TestBinarySearch_TargetPresentAtEnd(t *testing.T) {
	array := []int{1, 3, 5, 7, 9, 11}
	target := 11

	result := binarySearch(array, target)

	expected := 5
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_TargetPresent(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}
	target := 5

	result := binarySearch(array, target)

	assert.Equal(t, 2, result)
}

func TestBinarySearch_TargetNotPresent(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}
	target := 4

	result := binarySearch(array, target)

	if result != -1 {
		t.Errorf("Expected result to be -1, but got %d", result)
	}
}

func TestBinarySearch_EmptyArray(t *testing.T) {
	array := []int{}
	target := 5

	result := binarySearch(array, target)

	if result != -1 {
		t.Errorf("Expected result to be -1, got %d", result)
	}
}

func TestBinarySearch_OneElementTarget(t *testing.T) {
	array := []int{5}
	target := 5

	result := binarySearch(array, target)

	if result != 0 {
		t.Errorf("Expected index 0, but got %v", result)
	}
}

func TestBinarySearch_OneElementNotTarget(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}
	target := 6

	result := binarySearch(array, target)

	expected := -1
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_DuplicateElements(t *testing.T) {
	array := []int{1, 2, 2, 3, 4, 5, 6}
	target := 2

	result := binarySearch(array, target)

	expected := 1
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}
func TestBinarySearch_DuplicateElementsExcludingTarget(t *testing.T) {
	array := []int{1, 2, 2, 3, 4, 5}
	target := 2

	result := binarySearch(array, target)

	assert.Equal(t, 2, result)
}

func TestBinarySearch_NegativeNumbersAndNegativeTarget(t *testing.T) {
	array := []int{-10, -5, -3, -1}
	target := -5

	result := binarySearch(array, target)

	assert.Equal(t, 1, result)
}

func TestBinarySearch_NegativeNumbersPositiveTarget(t *testing.T) {
	array := []int{-10, -5, -3, -1}
	target := 5

	result := binarySearch(array, target)

	assert.Equal(t, -1, result)
}

func TestBinarySearchDescendingOrder(t *testing.T) {
	array := []int{9, 7, 5, 3, 1}
	target := 5

	result := binarySearch(array, target)

	expected := 2
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_MaxIntegerValues(t *testing.T) {
	array := []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	target := math.MaxInt32

	result := binarySearch(array, target)

	assert.Equal(t, 1, result)
}

func TestBinarySearch_MinimumIntegerValues(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 1

	result := binarySearch(array, target)

	expected := 0
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_ArrayWithNegativeAndPositiveNumbers(t *testing.T) {
	array := []int{-5, -3, 0, 2, 4, 6}
	target := 2
	expected := 3

	result := binarySearch(array, target)

	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}

func TestBinarySearch_AllIdenticalElements(t *testing.T) {
	array := []int{5, 5, 5, 5, 5}
	target := 5

	result := binarySearch(array, target)

	assert.Equal(t, 2, result)
}
