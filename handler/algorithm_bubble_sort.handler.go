package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"

	api "simple-open-api-spec-example/generated"
)

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
