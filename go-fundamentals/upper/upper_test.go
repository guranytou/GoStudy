package upper

import (
	"fmt"
	"testing"
)

func TestUpper(t *testing.T) {
	word := Upper("Gopher")
	expected := "GOPHER"

	if word != expected {
		t.Errorf("outcome %q but got %q", expected, word)
	}
}

func ExampleUpper() {
	word := Upper("test")
	fmt.Println(word)
	// Output: TEST
}
