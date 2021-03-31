package sorting

import (
	"errors"
)

/*
Subdivides input slice into smaller chunks
Accepts a slice of type string and returns n number of slices with equal number of elements
Appends the remaining number of elements to the last chunk
*/
func GetSliceChunks(data []string, n int) ([][]string, error) {
	if len(data) < n {
		return nil, errors.New("length of slice cannot be less than number of chunks")
	}
	if n == 0 {
		return nil, errors.New("number of chunks cannot be zero")
	}
	chunkSize := len(data) / n
	var chunks [][]string
	if n == 1 {
		return append(chunks, data), nil
	}
	for i := 0; i < n; i++ {
		upperbound := chunkSize * (i + 1)
		if i == n-1 {
			upperbound = len(data)
		}
		chunks = append(chunks, data[i*chunkSize:upperbound])
	}
	return chunks, nil
}

/*
Merges three sorted slices of strings into one output sorted slice of strings
*/
func MergeSlices(a, b, c []string) []string {

	i, j, k, l := 0, 0, 0, 0

	answer := make([]string, len(a)+len(b)+len(c))
	for ; i < len(a) && j < len(b) && k < len(c); l++ {
		if a[i] < b[j] && a[i] < c[k] {
			answer[l] = a[i]
			i++
		} else if b[j] < a[i] && b[j] < c[k] {
			answer[l] = b[j]
			j++
		} else {
			answer[l] = c[k]
			k++
		}
	}

	for ; i < len(a); i++ {
		answer[l] = a[i]
		l++
	}

	for ; j < len(b); j++ {
		answer[l] = b[j]
		l++
	}

	for ; k < len(c); k++ {
		answer[l] = c[k]
		l++
	}
	return answer
}
