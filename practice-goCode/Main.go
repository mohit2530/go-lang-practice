package main

import "practice-goCode/practice"

func main() {

	practice.GetNumberNineIfPresentInArray([4]int{1, 2, 3, 4})
	practice.GetNumberNineIfPresentInArray([4]int{1, 2, 9, 4})

	practice.FirstFourElementsIsSaidElement([4]int{1, 2, 3, 4}, 3)
	practice.FirstFourElementsIsSaidElement([4]int{1, 2, 3, 3}, 0)

	practice.ReturnTrueIfSequenceIsPresent([5]int{1, 2, 3, 4, 5}, [3]int{1, 2, 3})
	practice.ReturnTrueIfSequenceIsPresent([5]int{1, 2, 3, 4, 5}, [3]int{3, 4, 5})

	practice.NoTriples([5]int{1, 2, 3, 4, 5})
	practice.NoTriples([5]int{1, 2, 2, 2, 3})
	practice.NoTriples([5]int{1, 2, 3, 3, 3})

	practice.Has23([]int{1, 2, 3, 4, 5})
	practice.Has23([]int{1, 2, 4, 5})

	practice.MakeEndsMeet([4]int{1, 2, 3, 4})
	practice.MakeEndsMeet([4]int{5, 2, 3, 8})

	practice.MoreOneThanFour([]int{1, 2, 3, 4, 1, 1})
	practice.MoreOneThanFour([]int{1, 2, 3, 4, 1})

	practice.HasOneAndTwoAlso([]int{1, 1, 1, 1, 1, 2})
	practice.HasOneAndTwoAlso([]int{1, 1, 1, 1, 1, 0})

}
