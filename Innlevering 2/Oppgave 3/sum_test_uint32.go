package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 2, 3},
	{4, 5, 9},
	{120, 1, 119},
}

func TestSumuint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
