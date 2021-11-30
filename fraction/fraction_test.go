package main

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		{"", 6, []string{"1/2", "1/3", "2/3", "1/4", "3/4", "1/5", "2/5", "3/5", "4/5", "1/6", "5/6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions()\ngot  %v\nwant %v", got, tt.want)
			}
		})
	}
}
