package utils

import "rsc.io/quote"

func GetQuote() string {
    var quote = quote.Go()
    return quote
}

func InsertionSort(l []int) []int {
    for i := 0; i < len(l); i++ {
        j := i
        for ((j > 0) && l[j] < l[j-1]) {
            l[j-1], l[j] = l[j], l[j-1]
            j -= 1
        }
    }
    return l;
}
