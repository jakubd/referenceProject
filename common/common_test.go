package common

import "testing"

func TestAddNumbers(t *testing.T) {
	answerShouldBe := 2
	res := AddNumbers(1,1)

	if res != answerShouldBe {
		t.Errorf("add number failed should be %d but is %d", answerShouldBe, res)
	}
}