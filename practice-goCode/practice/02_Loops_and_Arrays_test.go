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
