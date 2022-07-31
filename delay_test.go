package delay

import "testing"

func TestGenerator(t *testing.T) {
	r := Generator()

	if r != 15 {
		t.Errorf("result is not 15")
	}
}
