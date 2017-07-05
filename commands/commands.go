package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	Name string `cli:"name" usage:"tell me your name"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("Hello, %s!\n", argv.Name)
		return nil
	})
}
