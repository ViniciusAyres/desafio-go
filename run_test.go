package desafiogo

/*
	AS DESCRIBED IN THE RULES: DO NOT EDIT THIS FILE!!!
*/

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		data     []float64
		expected Result
	}{
		{
			[]float64{1, 2, 2, 2, 3},
			Result{
				Average:      2,
				Median:       2,
				Percentile90: 3,
				Percentile99: 3,
				Mode:         2,
				HasNumber47:  false,
			},
		},
		{
			[]float64{47},
			Result{
				Average:      47,
				Median:       47,
				Percentile90: 47,
				Percentile99: 47,
				Mode:         0,
				HasNumber47:  true,
			},
		},
		{
			getDataSlice(),
			Result{
				Average:      48.42944397485905,
				Median:       37.59935969855793,
				Percentile90: 90.80600630442046,
				Percentile99: 99.95519419404887,
				Mode:         0,
				HasNumber47:  false,
			},
		},
	}

	for i, tc := range testCases {
		if r := *Run(tc.data); !reflect.DeepEqual(tc.expected, r) {
			t.Fatalf("test #%d failed!\n got:\n %v\n expected:\n %v \n", i, r, tc.expected)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	n := 10000
	data := getRandDataSlice(n)

	b.Run("baseline Run", func(b *testing.B) {
		BaselineRun(data)
	})

	b.Run("implemented Run", func(b *testing.B) {
		Run(data)
	})
}