package internal

import (
	"reflect"
	"testing"
)

func TestRotateTilesToRight(t *testing.T) {
	type testCase struct {
		name     string
		count    int
		expected Tiles
	}

	initial := Tiles{
		{1, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}

	cases := []testCase{
		{
			name:  "0 rotations",
			count: 0,
			expected: Tiles{
				{1, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name:  "1 rotation",
			count: 1,
			expected: Tiles{
				{0, 0, 1, 1},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name:  "2 rotations",
			count: 2,
			expected: Tiles{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 1, 0, 1},
				{0, 0, 0, 1},
			},
		},
		{
			name:  "3 rotations",
			count: 3,
			expected: Tiles{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{1, 1, 0, 0},
			},
		},
		{
			name:  "4 rotations",
			count: 4,
			expected: Tiles{
				{1, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}

	size := Size{Width: 4, Height: 4}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := rotateTilesToRight(initial, size, tc.count)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("%s:\ngot:\n%v\nwant:\n%v", tc.name, got, tc.expected)
			}
		})
	}
}
