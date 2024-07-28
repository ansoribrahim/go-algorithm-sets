package handler

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"

	api "simple-open-api-spec-example/generated"
)

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
