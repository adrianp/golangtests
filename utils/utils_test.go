package utils

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
    var tests = []struct {
        input []int
        want []int
    }{
        {[]int{3, 2, 1}, []int{1, 2, 3}},
        {[]int{}, []int{}},
        {[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
        {[]int{1, 2, 1, 2}, []int{1, 1, 2, 2}},
        {[]int{3, 1, 2, 4, 0, 19, -1, 17}, []int{-1, 0, 1, 2, 3, 4, 17, 19}},
        {[]int{1}, []int{1}},
    }
    for _, test := range tests {
        if got := InsertionSort(test.input); !reflect.DeepEqual(got, test.want) {
            t.Errorf("InsertionSort(%v) = %v, want %v", test.input, got, test.want)
        }
    }
}

func TestGetQuote(t *testing.T) {
    result := GetQuote()
    if len(result) < 1 {
        t.Errorf("Didn't get any quote")
    }
}
