package arithmetics

import (
	utils "golang-fullcycle-desafio-ci/utils"
	"testing"
)

func TestExpectedSum(t *testing.T) {
	const sumExpected int = 10
	sumResult := utils.Sum(5, 5)

	if sumResult != sumExpected {
		t.Errorf("Expected result is %d, but result is %d", sumExpected, sumResult)
	}
}
