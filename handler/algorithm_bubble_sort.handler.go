package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"

	api "simple-open-api-spec-example/generated"
)

/*
===========================================================================================================================================================================================================================================================================================
Algorithm Bubble Sort
===========================================================================================================================================================================================================================================================================================
Summary
This Go function implements the Bubble Sort algorithm to sort a slice of integers in ascending order. The algorithm repeatedly compares adjacent elements and swaps them if they are in the wrong order. It continues until the entire slice is sorted.

# Code Analysis

Inputs
array: A slice of integers to be sorted.

Flow
1. Get the length of the array.
2. Loop through the array multiple times (outer loop).
3. For each iteration of the outer loop, iterate through the array up to the unsorted portion (inner loop).
4. Compare adjacent elements and swap them if they are in the wrong order.
5. Reduce the range of the inner loop with each outer loop iteration as the largest elements "bubble up" to the end.
6. Exit when the array is fully sorted.

Outputs
The input slice is sorted in ascending order.
*/
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			println("i = ", i, "j = ", j)
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (h Handler) PostAlgorithmBubleSort(c *gin.Context) {

	var req api.PostAlgorithmBubleSortJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: null.StringFrom(err.Error()).Ptr()})
		return
	}

	bubbleSort(req.Array)

	c.JSON(http.StatusOK, api.AlgorithmBubbleSortResponse{SortedArray: &req.Array})
}
