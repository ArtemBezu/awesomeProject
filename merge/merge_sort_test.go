package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name    string
		reverse bool
		input   []int
		result  []int
	}{
		{"a", false, []int{1, 2, 5, 1}, []int{1, 1, 2, 5}},
		{"a", false, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"a", false, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"a", false, []int{0}, []int{0}},
		{"a", true, []int{1, 2, 5, 1}, []int{5, 2, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("input %v expect %v", tt.input, tt.result), func(t *testing.T) {
			if !reflect.DeepEqual(tt.result, Sort(tt.input, tt.reverse)) {
				t.Errorf("expect %v got %v", tt.input, tt.result)
			}
		})
	}
}
