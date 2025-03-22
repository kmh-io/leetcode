package binarysearch

func searchRange(nums []int, target int) []int {
	result, left, right := [2]int{-1, -1}, 0, len(nums)-1

	if left+right == 1 {
		if nums[left] == nums[right] {
			if nums[left] == target {
				return []int{0, 1}
			}
			return result[:]
		}

		if nums[left] == target {
			return []int{0, 0}
		}

		if nums[right] == target {
			return []int{1, 1}
		}

		return result[:]
	}

	for left < right {
		half := (left + right) / 2

		if nums[half] == target {
			leftPointer, rightPointer := half, half
			for rightPointer+1 <= right && nums[rightPointer+1] == target {
				rightPointer++
			}

			for leftPointer-1 >= left && nums[leftPointer-1] == target {
				leftPointer--
			}

			return []int{leftPointer, rightPointer}
		}

		if nums[half] > target {
			right = half
		} else {
			left = half + 1
		}
	}

	if left == right && nums[left] == target {
		return []int{left, left}
	}

	return result[:]
}
