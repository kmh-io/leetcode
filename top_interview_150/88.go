package topinterview150

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	if nums2[n-1] < nums1[0] {
		copy(nums1[n:], nums1[:m])
		copy(nums1[:n], nums2)
		return
	}

	i, j := 0, 0
	for ; i < m+n && j < n; i = i + 1 {
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
		}
	}

	if j < n {
		copy(nums1[m+j:], nums2[j:])
	}
}
