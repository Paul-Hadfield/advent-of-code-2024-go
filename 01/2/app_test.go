package main

import (
	"reflect"
	"testing"
)

func TestGetCounts(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   map[int]int
	}{
		{
			name:   "counts duplicates correctly",
			values: []int{1, 2, 2, 3, 2, 1},
			want: map[int]int{
				1: 2,
				2: 3,
				3: 1,
			},
		},
		{
			name:   "empty slice",
			values: []int{},
			want:   map[int]int{},
		},
		{
			name:   "single value",
			values: []int{5},
			want: map[int]int{
				5: 1,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := getCounts(tc.values); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("getCounts(%v) = %v, want %v", tc.values, got, tc.want)
			}
		})
	}
}

func TestTotalSimilarities(t *testing.T) {
	tests := []struct {
		name       string
		leftValues []int
		rightCounts map[int]int
		want       int
	}{
		{
			name:       "values with matches and missing entries",
			leftValues: []int{1, 2, 3},
			rightCounts: map[int]int{
				1: 1,
				2: 2,
			},
			want: 5, // 1*1 + 2*2 + 3*0
		},
		{
			name:       "multiple occurrences counted",
			leftValues: []int{5, 5, 7},
			rightCounts: map[int]int{
				5: 3,
				7: 1,
			},
			want: 5*3 + 5*3 + 7*1,
		},
		{
			name:       "no overlapping values",
			leftValues: []int{1, 2, 3},
			rightCounts: map[int]int{
				4: 2,
				5: 3,
			},
			want: 0,
		},
		{
			name:       "empty inputs",
			leftValues: []int{},
			rightCounts: map[int]int{},
			want:       0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := totalSimilarities(tc.leftValues, tc.rightCounts); got != tc.want {
				t.Fatalf("totalSimilarities(%v, %v) = %d, want %d", tc.leftValues, tc.rightCounts, got, tc.want)
			}
		})
	}
}
