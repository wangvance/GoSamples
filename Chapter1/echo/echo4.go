package main

import (
    "fmt"
    "os"
)

func main() {
    var s string
    var index int
    for index, s = range os.Args[1:] {
        fmt.Printf("%d\t%s\n",index,s)
    }
}
