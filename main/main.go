package main

import (
    "fmt"
    "example.org/utils"
)

func main() {
    fmt.Println(utils.GetQuote())

    list := []int{3, 1, 2, 4, 0, 19, 17}
    fmt.Println(utils.InsertionSort(list))
}
