package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	char := "t"
	cant := 9
	repeated := Repeat(char, cant)
	expected := strings.Repeat(char, cant)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

 func BenchmarkRepeat(b *testing.B) {
	char := "t"
	cant := 9
 	for i := 0; i < b.N; i++ {
 		Repeat(char, cant)
 	}
 }
