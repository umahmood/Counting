/*
Package counting implements the counting sort algorithm for integers.

Example:

    package main

    import (
        "fmt"

        "github.com/umahmood/counting"
    )

    func main() {
        nums := []int{42, -88, 1, -3, 2, 5}

        sorted := counting.Sort(nums)

        fmt.Println(sorted)

        nums = []int{-88, -3, 1, 2, 5, 42}

        sorted = counting.Sort(nums)

        fmt.Println(sorted)
    }
*/
package counting
