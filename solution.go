package main

// Solution is a simple bruteforce way to identify how many ways
// one can make @a aesthetic
func Solution(a []int) int {
	if isAesthetic(a) {
		return 0
	}
	numberOfWays := 0
	for i := 0; i < len(a)-2; i++ {
		//copy and remove
		temp := remove(a, i)
		//check if aesthetic
		if isAesthetic(temp) {
			numberOfWays++
		}
	}
	if numberOfWays == 0 {
		return -1
	}
	return numberOfWays
}

// isAesthetic returns true if @a is aesthetic and false if not
func isAesthetic(a []int) bool {
	for i := 1; i < len(a)-1; i++ {
		curr := a[i]
		prev := a[i-1]
		next := a[i+1]
		if curr >= prev && curr <= next ||
			curr <= prev && curr >= next {
			return false
		}
	}
	return true
}

// remove will copy the array @s and remove the @index
func remove(s []int, index int) []int {
	return append(append([]int{}, s[:index]...), s[index+1:]...)
}
