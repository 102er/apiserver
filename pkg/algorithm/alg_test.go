package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 13, 24, 32, 33, 33, 33, 33, 33, 67, 89, 123, 451}
	fmt.Println(binarySearch(nums, 33))
}

func TestBinarySearchLeftBound(t *testing.T) {
	nums := []int{1, 13, 24, 32, 33, 33, 33, 33, 33, 67, 89, 123, 451}
	fmt.Println(binarySearchRightBound(nums, 500))
}
