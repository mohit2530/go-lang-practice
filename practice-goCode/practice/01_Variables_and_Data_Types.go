package practice

import (
	"fmt"
)

// GetNumberNineIfPresentInArray - return the number of 9's in the arraylist
func GetNumberNineIfPresentInArray(numbers [4]int) int {
	count := 0
	for _, element := range numbers {
		if element == 9 {
			count++
		}
	}
	fmt.Printf("found %d number nine(s) in the arraylist\n", count)
	return count
}

// FirstFourElementsIsSaidElement - return true if the element exists in the arraylist
func FirstFourElementsIsSaidElement(numbers [4]int, searchFor int) bool {

	for _, element := range numbers {
		if element == searchFor {
			fmt.Printf("item found. searching for : %d, found : %d \n", searchFor, searchFor)
			return true
		}
	}
	fmt.Printf("item not found. searching for : %d, not found : %d \n", searchFor, searchFor)
	return false
}

// ReturnTrueIfSequenceIsPresent - returns true if the given sequence is found
func ReturnTrueIfSequenceIsPresent(numbers [5]int, sequence [3]int) bool {

	for index, element := range numbers {
		// if the index already exceeds the length of the array to traverse, there is no need to check
		if index > len(numbers)-3 {
			fmt.Printf("index already exceeds sequence. index: %d, sequence %v. returns - %v\n", index, sequence, false)
			return false
		}
		// checking if the sequence has the numbers provided
		if element == sequence[0] && numbers[index+1] == sequence[1] && numbers[index+2] == sequence[2] {
			fmt.Printf("sequence found. index: %d, sequence %v. returns - %v\n", index, sequence, true)
			return true
		}
	}
	fmt.Printf("sequence was not found. Returns : %v\n", false)
	return false
}

// NoTriples - returns true if there are no triples
func NoTriples(numbers [5]int) bool {

	for index, element := range numbers {

		if index > len(numbers)-3 {
			return false
		}

		if element == numbers[index+1] && element == numbers[index+2] {
			fmt.Printf("triples found for the element : %d \n", element)
			return false
		}
	}
	return true
}

// MakeEndsMeet - returns an array with the first and last numbers together
func MakeEndsMeet(numbers [4]int) []int {

	var makeEnds []int
	makeEnds = append(makeEnds, numbers[0], numbers[3])
	fmt.Printf("previous array : %v . new array : %v \n", numbers, makeEnds)
	return makeEnds
}

// Has23 - returns true if the given array has 2 or 3
func Has23(numbers []int) bool {

	if len(numbers) < 2 {
		fmt.Printf("the array has length less than 2.\n")
		return false
	}
	for index, element := range numbers {
		if element == 2 || element == 3 {
			fmt.Printf("contains the value %d or %d at index : %d\n", 2, 3, index)
			return true
		}
	}
	return false
}

// MoreOneThanFour - returns true if the count of 1 is greater than the count of 4
func MoreOneThanFour(numbers []int) bool {
	oneCount, fourCount := 0, 0
	for _, element := range numbers {
		if element == 1 {
			oneCount++
		}
		if element == 4 {
			fourCount++
		}
	}
	fmt.Printf("oneCount : %d and fourCount : %d \n", oneCount, fourCount)
	return oneCount > fourCount
}

// HasOneAndTwoAlso - returns true if value 1 is present and 2 is also present somewhere else
func HasOneAndTwoAlso(numbers []int) bool {

	for index, element := range numbers {
		if element == 1 {
			// length of numbers array - index of the found value of 1
			for i := 0; i < len(numbers)-index; i++ {
				if numbers[i] == 2 {
					fmt.Printf("found both values. %v \n", true)
					return true
				}
			}
		}
	}
	fmt.Printf("did not find both values. %v \n", false)
	return false
}
