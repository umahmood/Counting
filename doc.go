/*
Package counting implements the counting sort algorithm for integers.

Example:

    package main

    import (
        "fmt"

        "github.com/umahmood/counting"
    )

    func main() {
        nums := []int{10, 7, 3, 1, 6, 4, 8, 2, 5, 9}

        sorted := counting.Sort(nums)

        fmt.Println(sorted)

        nums = []int{-88, -3, 1, 2, 5, 42}

        sorted = counting.Sort(nums)

        fmt.Println(sorted)
    }
*/
package counting
