package practice

import "testing"

func TestFirstLast6(t *testing.T) {

	shouldReturnTrue := FirstLast6([]int{1, 2, 3}, 3)
	shouldReturnFalse := FirstLast6([]int{2, 3, 4}, 5)

	if !shouldReturnTrue {
		t.Errorf("%d was not found in the first / last item of the array", 3)
	}

	if shouldReturnFalse {
		t.Errorf("%d was not found in the first / last item of the array", 5)
	}

}

func TestSumOfItems(t *testing.T) {

	shouldBe10 := SumOfItems([]int{1, 2, 3, 4})

	if shouldBe10 != 10 {
		t.Errorf("the sum of items was not %d", 10)
	}

}

func TestSumOfItems2(t *testing.T) {

	shouldBe14 := SumOfItems([]int{1, 2, 3, 4, 2, 1, 1, 2})

	if shouldBe14 == 14 {
		t.Errorf("the sum of items was %d", 16)
	}
}

func TestRotateLeft3(t *testing.T) {

	expected := []int{3, 2, 1}
	actual := RotateLeft3(expected)

	if len(actual) != len(expected) {
		t.Errorf("actual and expected length do not match. Actual : %v, Expected : %v", actual, expected)
	}

}

func TestReverseOrder(t *testing.T) {

	expected := []int{3, 2, 1}
	actual := ReverseOrder([]int{1, 2, 3})

	if len(expected) != len(actual) {
		t.Errorf("array was not reversed. Expected : %v, Got : %v", expected, actual)
	}

}

func TestMaxEnds3Part01(t *testing.T) {

	expected := []int{4, 4, 4}
	actual := MaxEnds3([]int{1, 2, 4})

	if len(expected) != len(actual) {
		t.Errorf("highest not return in all across. Expected : %v, Got : %v", expected, actual)
	}
}

func TestMaxEnds3Part02(t *testing.T) {
	expected := []int{3, 3, 3}
	actual := MaxEnds3([]int{3, 2, 2})

	if len(actual) != len(expected) {
		t.Errorf("highest not returned in all across. Expected : %v, Got : %v", expected, actual)
	}
}

func TestMiddleWay(t *testing.T) {
	expected := []int{1, 1}
	actual := MiddleWay([]int{2, 1, 3}, []int{3, 1, 2})

	if len(actual) != len(expected) {
		t.Errorf("returning array is not the middle values of both array. "+
			"Expected : %v, got : %v", expected, actual)
	}

	expected01 := []int{2, 4}
	actual01 := MiddleWay([]int{5, 2, 9}, []int{1, 4, 5})

	if len(actual01) != len(expected01) {
		t.Errorf("middle way array is not the middle item of both arrays. "+
			"Expected : %v, Actual : %v", expected01, actual01)
	}

}

func TestSumOf28(t *testing.T) {
	actual := SumOf28([]int{1, 3, 2, 4, 5, 3, 2, 7, 2, 8, 7, 2})

	if !actual {
		t.Errorf("sum of all the %d's should be %d", 2, 8)
	}
}
