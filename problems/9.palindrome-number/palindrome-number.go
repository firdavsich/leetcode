package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var arr []int
	for x > 0 {

		arr = append(arr, x%10)
		x /= 10
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}
