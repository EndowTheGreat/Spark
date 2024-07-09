package main

import (
	"fmt"

	"gitlab.com/EndowTheGreat/spark/cmd/cli"
)

func main() {
	if err := cli.Setup(); err != nil {
		fmt.Println(err)
	}
}
