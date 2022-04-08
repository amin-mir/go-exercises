package slice

import (
	"reflect"
	"testing"
)

func TestSliceReverse(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "should reverse a slice with three elements",
			in:   []int{1, 2, 3},
			out:  []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := Reverse(tt.in)
			reflect.DeepEqual(tt.out, out)
		})
	}
}
