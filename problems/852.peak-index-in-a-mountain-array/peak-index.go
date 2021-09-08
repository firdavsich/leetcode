package peakindex

func peakIndexInMountainArray(arr []int) int {
	peak := 0
	for i, v := range arr {
		if v > arr[peak] {
			peak = i
		}
	}
	return peak
}
