package practice

import "fmt"

// FirstLast6 - return true if itemToFind is in the first or last element in the array
func FirstLast6(numbers []int, itemToFind int) bool {
	// if the first or last item is itemToFind
	if numbers[0] == itemToFind || numbers[len(numbers)-1] == itemToFind {
		fmt.Printf("number %d is present in %v \n", itemToFind, numbers)
		return true
	}
	fmt.Printf("number %d is not present in %v \n", itemToFind, numbers)
	return false
}

// SumOfItems - return the sum of all the items within an array
func SumOfItems(numbers []int) float64 {
	sumTotal := 0
	for _, element := range numbers {
		sumTotal += element
	}
	fmt.Printf("sum total of all items in an array is : %d \n", sumTotal)
	return float64(sumTotal)
}

// RotateLeft3 - for array of 3 items, items are rotated left. so {1,2,3} becomes {2,3,1}
func RotateLeft3(numbers []int) []int {
	var rotatedValues []int

	if len(numbers) == 3 {
		for index := range numbers {
			// if index + 1 doesn't exist in the numbers array
			// we don't want to try to append
			if index == len(numbers)-1 {
				rotatedValues = append(rotatedValues, numbers[0])
				break
			}
			rotatedValues = append(rotatedValues, numbers[index+1])
		}
	}
	fmt.Printf("for %v, rotated values are : %v \n", numbers, rotatedValues)
	return rotatedValues
}

// ReverseOrder - returns the same array in the reverse order
func ReverseOrder(numbers []int) []int {

	var reversedArr []int
	for index := range numbers {
		reversedArr = append(reversedArr, numbers[len(numbers)-(index+1)])
	}
	fmt.Printf("reversed array of the array %v is %v\n", numbers, reversedArr)
	return reversedArr
}
