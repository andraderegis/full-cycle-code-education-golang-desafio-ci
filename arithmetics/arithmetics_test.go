package arithmetics

import (
	"testing"
)

func TestExpectedSum(t *testing.T) {
	const sumExpected int = 10
	sumResult := Sum(5, 5)

	if sumResult != sumExpected {
		t.Errorf("Expected result is %d, but result is %d", sumExpected, sumResult)
	}
}
