package testcases

import (
	"testing"

	"testcases"
)

func TestEven(t *testing.T) {
	n := 2
	expected := "YES"
	res := testcases.CheckEven(n)

	if expected != res {
		t.Errorf("Expected %v, Got: %v", expected, res)
	}
}
