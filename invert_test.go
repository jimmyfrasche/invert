package invert

import "testing"

var tests = []struct {
	len     int
	in, out [][]int
}{
	{
		len: 10,
		in:  [][]int{{0, 10}},
		out: nil,
	},
	{
		len: 10,
		in:  nil,
		out: [][]int{{0, 10}},
	},
	{
		len: 10,
		in:  [][]int{{0, 5}},
		out: [][]int{{5, 10}},
	},
	{
		len: 10,
		in:  [][]int{{5, 10}},
		out: [][]int{{0, 5}},
	},
	{
		len: 10,
		in:  [][]int{{1, 9}},
		out: [][]int{{0, 1}, {9, 10}},
	},
	{
		len: 10,
		in:  [][]int{{0, 9}},
		out: [][]int{{9, 10}},
	},
	{
		len: 10,
		in:  [][]int{{1, 10}},
		out: [][]int{{0, 1}},
	},
	{
		len: 10,
		in:  [][]int{{1, 4}, {6, 9}},
		out: [][]int{{0, 1}, {4, 6}, {9, 10}},
	},
	{
		len: 10,
		in:  [][]int{{0, 4}, {6, 10}},
		out: [][]int{{4, 6}},
	},
}

func TestIndicies(t *testing.T) {
	for i, test := range tests {
		out := Indicies(test.in, test.len)
		if len(out) != len(test.out) {
			t.Errorf("%d: length mismatch got: %d, expected: %d", i, len(out), test.len)
			continue
		}
		for j, p := range test.out {
			q := out[j]
			if len(q) != 2 {
				t.Errorf("%d: vector %d malformed, got: %v", i, j, q)
				continue
			}
			if p[0] != q[0] || p[1] != q[1] {
				t.Errorf("%d: vector %d expected %v got %v", i, j, p, q)
				continue
			}
		}
	}
}
