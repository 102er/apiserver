package algorithm

/*
给定一个个元素有序的（升序）整型数组nums 和一个目标值 target，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
*/
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binarySearchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func binarySearchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}
