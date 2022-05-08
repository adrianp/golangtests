package main

import (
    "example.org/utils"
    "os"
)

func main() {
    //fmt.Println(strings.Join(os.Args[1:], ","))
    //fmt.Println(utils.GetQuote())

    //list := []int{3, 1, 2, 4, 0, 19, 17}
    //fmt.Println(utils.InsertionSort(list))

    utils.Lissajous(os.Stdout)
}
