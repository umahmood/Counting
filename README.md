# Counting

A small Go library which implements the [counting sort algorithm](https://en.wikipedia.org/wiki/Counting_sort).

# Installation

> go get github.com/umahmood/counting

> cd $GOPATH/src/github.com/umahmood/counting/

> go test -v ./...

# Usage

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

# Documentation

> [http://godoc.org/github.com/umahmood/counting](> http://godoc.org/github.com/umahmood/counting)

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
