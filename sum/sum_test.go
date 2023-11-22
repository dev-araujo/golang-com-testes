package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expectation := 15

	if result != expectation {
		t.Errorf("resultado %d, expectativa %d, dados %v", result, expectation, numbers)
	}
}
