package ch001

import (
	"testing"

	"github.com/matryer/is"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			data:     []int{7, 2},
			expected: 5,
		},
		{
			data:     []int{2, 7, 3},
			expected: -2,
		},
		{
			data:     []int{1000, 1000, 1000, 1000, 1000},
			expected: 1000,
		},
		{
			data:     []int{823, 912, 345, 100000, 867, 222, 991, 3, 40000},
			expected: -58111,
		},
		{
			data:     []int{23, 35, 12, 100000, 99234, 86123, 3245},
			expected: -83644,
		},
		{
			data:     []int{23, 35, 12, 100000, 99234, 86123, 3245, 1},
			expected: 83645,
		},
		{
			data:     []int{7, 7, 7, 7, 7, 7, 80, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			expected: -66,
		},
		{
			data:     []int{7, 7, 7, 7, 7, 7, 7, 80, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			expected: 73,
		},
		{
			data:     []int{91, 56, 23, 45, 87, 65, 45, 45, 78, 23, 20, 41, 17, 54, 51, 51, 94, 62, 74, 42, 76, 76},
			expected: 96,
		},
		{
			data:     []int{92834, 95461, 15911, 56189, 6369, 80545, 31811, 51263, 30076, 68867, 36905, 32499, 59799, 334, 82991, 46636, 98741, 66601},
			expected: 42958,
		},
		{
			data:     []int{129},
			expected: 129,
		},
		{
			data:     []int{35463, 88121, 4362, 94457, 86235, 83680, 72686, 6003, 93069, 2015, 10436, 2139, 93162, 30380, 19067, 76335, 78941, 48620, 55887, 15679},
			expected: 101879,
		},
		{
			data:     []int{19335, 97643, 11468, 86267, 79718, 59584, 12129, 52642, 86575, 62307, 11545, 52658, 72377, 39986, 74850, 1992, 86928},
			expected: 1846,
		},
		{
			data:     []int{91883, 97793, 54567, 64714, 98624},
			expected: 82567,
		},
		{
			data:     []int{98473, 41866, 71129, 65936, 42626, 9194, 46718, 96921, 45613, 47677, 8763, 54634, 47259, 71448, 9918, 22666, 32711, 21692, 40207, 2017, 23040, 86083, 77809, 15472, 30718, 39085, 87911, 54827, 41686, 28354, 37203, 6548, 74184, 3043, 43961, 95189, 1238, 22002, 93507, 63546, 32527, 42778, 31614},
			expected: -14953,
		},
	}

	iss := is.New(t)
	for i := 0; i < len(tests); i++ {
		got := Solve(tests[i].data)
		iss.Equal(got, tests[i].expected)
	}
}
