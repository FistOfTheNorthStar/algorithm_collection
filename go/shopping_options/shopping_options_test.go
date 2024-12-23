package shopping

import "testing"

func TestGetNumberOfOptions(t *testing.T) {
	tests := []struct {
		name          string
		priceOfJeans  []int
		priceOfShoes  []int
		priceOfSkirts []int
		priceOfTops   []int
		dollars       int
		want          int
		wantErr       bool
	}{
		{
			name:          "basic test",
			priceOfJeans:  []int{2, 3},
			priceOfShoes:  []int{4},
			priceOfSkirts: []int{2, 3},
			priceOfTops:   []int{1, 2},
			dollars:       10,
			want:          4,
			wantErr:       false,
		},
		{
			name:          "test with big prices",
			priceOfJeans:  []int{2, 10000, 3},
			priceOfShoes:  []int{2000002, 4},
			priceOfSkirts: []int{2, 3000000, 3},
			priceOfTops:   []int{1, 2},
			dollars:       10,
			want:          4,
			wantErr:       false,
		},
		{
			name:          "invalid budget",
			priceOfJeans:  []int{2, 3},
			priceOfShoes:  []int{4},
			priceOfSkirts: []int{2, 3},
			priceOfTops:   []int{1, 2},
			dollars:       0,
			want:          0,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNumberOfOptions(tt.priceOfJeans, tt.priceOfShoes,
				tt.priceOfSkirts, tt.priceOfTops, tt.dollars)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
