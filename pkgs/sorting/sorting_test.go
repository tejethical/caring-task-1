package sorting

import (
	"github.com/davecgh/go-spew/spew"
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
		spew.Dump(testChunks[i])
		totalLen += len(testChunks[i])
	}
	if totalLen != len(testArr) {
		spew.Dump(totalLen, testArr)
		t.Error("total length of generated chunks doesn't match the input array")
	}
}
