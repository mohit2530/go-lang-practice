package practice

import "testing"

func TestGetNumberNineIfPresentInArray(t *testing.T) {

	expected := 0
	req := [4]int{1, 2, 3, 4}
	hasNumberNines := GetNumberNineIfPresentInArray(req)

	if hasNumberNines != expected {
		t.Errorf("expected %d. but got %v", expected, hasNumberNines)
	}
}

func TestGetNumberNineIfPresentInArray2(t *testing.T) {

	expected := 1
	req := [4]int{1, 9, 2, 3}
	hasNumberNines := GetNumberNineIfPresentInArray(req)
	if hasNumberNines != expected {
		t.Errorf("expected %d. but got %v", expected, hasNumberNines)
	}

}

func TestFirstFourElementsIsSaidElement(t *testing.T) {

	firstReqArr, firstReqInt := [4]int{1, 2, 3, 4}, 4
	secondReqArr, secondReqInt := [4]int{1, 2, 3, 4}, 5

	isValid := FirstFourElementsIsSaidElement(firstReqArr, firstReqInt)
	if !isValid {
		t.Errorf("expected %v. but got %v", isValid, !isValid)
	}

	isSecondValid := FirstFourElementsIsSaidElement(secondReqArr, secondReqInt)
	if isSecondValid {
		t.Errorf("expected %v. but got %v", isValid, !isValid)
	}

}

func TestReturnTrueIfSequenceIsPresent(t *testing.T) {

	firstReqArr, firstReqSequence := [5]int{1, 2, 3, 4, 5}, [3]int{1, 2, 3}
	secondReqArr, secondReqSequence := [5]int{1, 2, 3, 4, 5}, [3]int{5, 6, 7}

	isFirstValid := ReturnTrueIfSequenceIsPresent(firstReqArr, firstReqSequence)
	isSecondValid := ReturnTrueIfSequenceIsPresent(secondReqArr, secondReqSequence)

	if !isFirstValid {
		t.Errorf("expected %v. but got %v", isFirstValid, !isFirstValid)
	}

	if isSecondValid {
		t.Errorf("expected %v. but got %v", isSecondValid, !isSecondValid)
	}

}

func TestNoTriples(t *testing.T) {

	shouldBeTrue := NoTriples([5]int{1, 2, 3, 4, 5})
	shouldBeFalse := NoTriples([5]int{2, 2, 2, 3, 1})

	if !shouldBeTrue {
		t.Errorf("expected %v. but got %v", shouldBeTrue, !shouldBeTrue)
	}

	if shouldBeFalse {
		t.Errorf("expected %v. but got %v", shouldBeFalse, !shouldBeFalse)

	}

}

func TestMakeEndsMeet(t *testing.T) {

	shouldReturn := []int{1, 4}
	actualReturn := MakeEndsMeet([4]int{1, 2, 3, 4})

	if len(shouldReturn) != len(actualReturn) {
		t.Errorf("actual return is not equal to expected return. Actual : %v, "+
			"Expected : %v", actualReturn, shouldReturn)
	}

}

func TestHas23(t *testing.T) {
	actualReturn01 := Has23([]int{1, 2})
	actualReturn02 := Has23([]int{1, 3})

	if !actualReturn01 || !actualReturn02 {
		t.Errorf("actual return is not equal to expected return. "+
			"Actual01 : %v, Actual02 : %v, Expected : %v", actualReturn01, actualReturn02, true)
	}

}

func TestMoreOneThanFour(t *testing.T) {

	actualReturn := MoreOneThanFour([]int{1, 2, 3})
	falselyReturn := MoreOneThanFour([]int{2, 3, 4})

	if !actualReturn || falselyReturn {
		t.Errorf("actual return is not equal to expected return. "+
			"ActualReturn : %v, Expected : %v", actualReturn, true)
		t.Errorf("actual return is not equal to expected return. "+
			"ActualReturn : %v, Expected : %v", falselyReturn, false)
	}

}

func TestHasOneAndTwoAlso(t *testing.T) {

	shouldBeFalse := HasOneAndTwoAlso([]int{1, 4, 5, 7, 3})
	shouldBeFalseWithNo1 := HasOneAndTwoAlso([]int{3, 5, 4, 6, 2})
	shouldBeFalseWithNo2 := HasOneAndTwoAlso([]int{1, 4, 6, 3, 1})

	if shouldBeFalse {
		t.Errorf("should be false. ")
	}
	if shouldBeFalseWithNo1 {
		t.Errorf("should be false with no number ones. ")
	}
	if shouldBeFalseWithNo2 {
		t.Errorf("should be false with no number twos.")
	}

}
