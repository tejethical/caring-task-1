package sorting

import (
	"reflect"
	"testing"
)

func TestGetSliceChunks(t *testing.T) {
	_, err := GetSliceChunks([]string{}, 3)
	if err == nil {
		t.Error("func should return an error on sending empty array")
	}
	testArr := []string{"abc", "def", "ghi", "jkl", "mno"}

	_, err = GetSliceChunks(testArr, 0)
	if err == nil {
		t.Error("func should return an error on n = 0")
	}

	_, err = GetSliceChunks(testArr, len(testArr)+1)
	if err == nil {
		t.Error("func should return an error if n is greater than length of array")
	}

	n := 4
	testChunks, err := GetSliceChunks(testArr, n)
	if err != nil {
		t.Error("Func result validation failed")
	}
	if len(testChunks) != n {
		t.Error("number of returned chunks is not equal to n")
	}
	totalLen := 0
	for i := 0; i < n; i++ {
		totalLen += len(testChunks[i])
	}
	if totalLen != len(testArr) {
		t.Error("total length of generated chunks doesn't match the input array")
	}
}

type TestMergeSlicesTable struct {
	testId int
	a []string
	b []string
	c []string
	err bool
	output []string
	outputLength int
}

func TestMergeSlices(t *testing.T) {
	testCases := []TestMergeSlicesTable{
		{testId: 1, a: []string{}, b: []string{}, c: []string{}, err: true},
		{testId: 2, a: []string {"abc", "def", "ghi"}, b: []string{}, c: []string{}, err: false, output: []string{"abc", "def", "ghi"}},
		{testId: 3, a: []string{}, b: []string {"abc", "def", "ghi"}, c: []string{}, err: false, output: []string{"abc", "def", "ghi"}},
		{testId: 4, a: []string{}, b: []string{}, c: []string {"abc", "def", "ghi"}, err: false, output: []string{"abc", "def", "ghi"}},
		{testId: 5, a: []string {"abc", "def", "ghi"}, b: []string {"jkl", "mno", "pqr"}, c: []string {"stu", "vwx", "yz"}, err: false, output: []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}},
	}

	for _, testCase := range testCases {
		res, err := MergeSlices(testCase.a, testCase.b, testCase.c)

		if testCase.err {
			if err == nil {
				t.Error("func should return an error")
			}
		} else {
			if !reflect.DeepEqual(testCase.output, res) {
				t.Errorf("Test # %d failed. Expected: %v, Returned: %v", testCase.testId, testCase.output, res)
			}
		}
	}

}
