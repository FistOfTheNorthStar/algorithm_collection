package algorithms

import (
	"reflect"
	"testing"
)

func TestRotateMat90Clock(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "4x4 matrix",
			matrix: [][]int{
				{9, 10, 11, 12},
				{16, 17, 18, 19},
				{23, 24, 25, 26},
				{30, 31, 32, 33},
			},
			want: [][]int{
				{30, 23, 16, 9},
				{31, 24, 17, 10},
				{32, 25, 18, 11},
				{33, 26, 19, 12},
			},
		},
		{
			name: "5x5 matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			want: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("Rotate() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
