package vma

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	isInc := arr[0] < arr[1]
	if !isInc {
		return false
	}

	switchCnt := 0
	for i := 1; i+1 < len(arr); i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if switchCnt > 1 {
			break
		}
		if arr[i+1] > arr[i] && isInc {
			continue
		}
		if arr[i+1] < arr[i] && !isInc {
			continue
		}
		if arr[i+1] < arr[i] && isInc {
			isInc = false
			switchCnt++
			continue
		}
		if arr[i+1] > arr[i] && !isInc {
			switchCnt++
			continue
		}
	}
	return switchCnt == 1
}

func ValidMountainArray(arr []int) bool {
	return validMountainArray(arr)
}
