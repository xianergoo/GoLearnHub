package main

func isPalindrome(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}

	revertNumber := 0
	for x > revertNumber {
		revertNumber = revertNumber*10 + x%10
		x /= 10
	}
	return x == revertNumber || x == revertNumber/10

}

func main() {
	isPalindrome(12)
}
