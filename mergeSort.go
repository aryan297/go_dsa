package main

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
	temp := []int{}
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}

}
