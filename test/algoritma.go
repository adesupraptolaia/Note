/*
============ Basic Concepts

1A. Long method is when you write a lot of complex logic in a method, it's usually because the method has too many responsibilities. It makes the code hard to maintain and debug.
--- How to prevent it? Try to separate the method into several smaller methods which have one precise to do

1B. Dependency Injection is when a class independent of its depedencies, usually you make a interface to represent the depedency.
--- Why is it important? If you need to change/add your depedency from A to B, you just focus to change/add the depedency without change another code.

2 POST
-- do : validate request payload
-- don't : don't save the data before you know it saves

2 GET
-- do : validate user access
-- don't : don't give user request outside his access
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	sockMerchant()
	// countingValleys()
	// numberThree()
	// determineTotalOfLampIsOn()
}

func sockMerchant() {
	var numberOfSocks, sockNumber int
	fmt.Scanf("%d", &numberOfSocks)

	CollectAllMatchingSock := make(map[int]int)

	for i := 0; i < numberOfSocks; i++ {
		fmt.Scanf("%d", &sockNumber)

		CollectAllMatchingSock[sockNumber]++
	}

	var totalMatchingSock int
	for _, totalSock := range CollectAllMatchingSock {
		totalMatchingSock += int(totalSock / 2)
	}

}

func countingValleys() {
	var garySteps, garryPositition, result int
	var path string

	fmt.Scanf("%d", &garySteps)

	for i := 0; i < garySteps; i++ {
		fmt.Scanf("%s", &path)

		if path == "D" {
			garryPositition--
		} else if path == "U" {
			garryPositition++
		}

		if garryPositition == -1 && path == "D" {
			result++
		}
	}

}

func numberThree() {
	var number string

	fmt.Scanf("%s", &number)

	lengthOfNumber := len(number)

	for i := 1; i <= lengthOfNumber; i++ {
		num := string(number[i-1])

		fmt.Println(num + strings.Repeat("0", lengthOfNumber-i))
	}
}

// pseudo code for number 3
/*
number = string of input

for-loop i=1, i <= len(number), i++
	subNumber = string of number[i-1]

	println ==> subNumber + repeteadZero[len(number)-i]times
*/

func determineTotalOfLampIsOn() {
	primeNumbers := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}
	primeNumbersMap := make(map[int]bool)
	for _, prime := range primeNumbers {
		primeNumbersMap[prime] = true
	}

	var result int
	for n := 1; n <= 100; n++ {
		if isLampOn(n, primeNumbers, primeNumbersMap) {
			result++
		}
	}

}

// find all factor of num
// e.g 12 ==> 1, 2, 3, 4, 6, 12 ==> there are 6 factor of 12
// find by fa
// 12 = 2^2 * 3^1 ===> (2+1)*(1+1) = 6
func isLampOn(num int, primeNumbers []int, primeNumbersMap map[int]bool) bool {
	if primeNumbersMap[num] {
		return false
	}

	result := 1
	for _, prime := range primeNumbers {
		power := 0

		// find the biggest of power ==> num % prime^power == 0
		for num%prime == 0 {
			power++
			num /= prime
		}

		if power > 0 {
			result *= (power + 1)
		}

		if num == 1 {
			break
		}
	}

	return result%2 != 0
}
