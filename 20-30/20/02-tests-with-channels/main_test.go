package main

import (
	"reflect"
	"testing"
)

func TestGenerateInts(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "default testing case",
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := GenerateInts()
			var got []int

			for i := range cgot {
				got = append(got, i)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
