package leetcode

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// assume the length of nums1 is less
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays1(nums2, nums1)
	}
	low, high, k, n1Mid, n2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)/2, 0, 0
	for low <= high {
		// nums1 .....nums1[n1Mid-1] | nums1[n1Mid].....
		// nums2 .....nums2[n2Mid-1] | nums2[n2Mid].....
		n1Mid = low + (high-low)/2
		n2Mid = k - n1Mid
		if n1Mid > 0 && nums1[n1Mid-1] > nums2[n2Mid] {
			high = n1Mid - 1
		} else if n1Mid < len(nums1) && nums1[n1Mid] < nums2[n2Mid-1] {
			low = n1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if n1Mid == 0 {
		midLeft = nums2[n2Mid-1]
	} else if n2Mid == 0 {
		midLeft = nums1[n1Mid-1]
	} else {
		midLeft = max(nums1[n1Mid-1], nums2[n2Mid-1])
	}
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(midLeft)
	}
	if n1Mid == len(nums1) {
		midRight = nums2[n2Mid]
	} else if n2Mid == len(nums2) {
		midRight = nums1[n1Mid]
	} else {
		midRight = min(nums1[n1Mid], nums2[n2Mid])
	}
	return float64((midLeft + midRight)) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
