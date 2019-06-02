package golang_lesson1

import "testing"

func TestPrintTime(t *testing.T) {
	err := PrintTime()

	if err != nil {
		t.Fatal(err)
	}
}
