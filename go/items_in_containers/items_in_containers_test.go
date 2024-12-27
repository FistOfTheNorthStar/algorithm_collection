package algorithms

import (
	"reflect"
	"testing"
)

func TestNumberOfItemsInContainer(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		startIndices []int
		endIndices   []int
		want         []int
		wantErr      bool
	}{
		{
			name:         "basic test",
			s:            "|**|*|*",
			startIndices: []int{1, 1},
			endIndices:   []int{5, 6},
			want:         []int{2, 3},
			wantErr:      false,
		},
		{
			name:         "invalid character",
			s:            "|**|x|*",
			startIndices: []int{1},
			endIndices:   []int{6},
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "invalid indices size",
			s:            "|**|*|*",
			startIndices: []int{},
			endIndices:   []int{5},
			want:         nil,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NumberOfItemsInContainer(tt.s, tt.startIndices, tt.endIndices)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberOfItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberOfItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
