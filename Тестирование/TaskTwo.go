package math

import "testing"

type testpair struct {
	values  []float64
	min     float64
	max     float64
	average float64
}

var tests = []testpair{
	{[]float64{6, 10, 2, 7}, 2.0, 10.0, 6.25},
	{[]float64{-8, 1, 8, -2, 4, -9}, -9.0, 8.0, -1.0},
	{[]float64{-1, 1}, -1.0, 1.0, 0.0},
	{[]float64{}, 0.0, 0.0, 0.0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		v, ok := Min(pair.values)
		if !ok {
			t.Error("Len array 0")
		}
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		v, ok := Max(pair.values)
		if !ok {
			t.Error("Len array 0")
		}
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
