package main

import (
	"fmt"

	"github.com/mharbol/aoc-2023/cmd"
)

func main() {
    out, err := cmd.RunProb(&cmd.CmdParser{})

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(out)
    }
}
