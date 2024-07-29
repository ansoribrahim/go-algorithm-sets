package handler

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"

	api "simple-open-api-spec-example/generated"
)

/*
===========================================================================================================================================================================================================================================================================================
Algorithm Binary Search
===========================================================================================================================================================================================================================================================================================
Summary
This Go function implements the binary search algorithm to find the index of a target value within a sorted integer array. If the target is not found, it returns -1.
Example Usage
array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
target := 5
index := binarySearch(array, target)
fmt.Println(index) // Output: 4

# Code Analysis

Inputs
array: A sorted slice of integers.
target: An integer value to search for within the array.

Flow
1. Initialize left and right pointers to the start and end of the array.
2. Enter a loop that continues as long as left is less than or equal to right.
3. Calculate the middle index mid.
4. If the middle element equals the target, return mid.
5. If the middle element is less than the target, adjust left to mid + 1.
6. If the middle element is greater than the target, adjust right to mid - 1.
7. If the loop exits without finding the target, return -1.

Outputs
The index of the target value if found, otherwise -1.
*/
func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func (h Handler) AlgorithmBinarySearch(c *gin.Context) {
	var req api.AlgorithmBinarySearchJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: null.StringFrom(err.Error()).Ptr()})
		return
	}

	if !sort.IntsAreSorted(req.Array) {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: null.StringFrom("Invalid input: array is not sorted").Ptr()})
		return
	}

	index := binarySearch(req.Array, req.Target)
	c.JSON(http.StatusOK, api.AlgorithmBinarySearchResponse{Index: &index})
}
