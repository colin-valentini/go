package main

import "github.com/alecthomas/kong"

var cli struct {
	Init InitCmd `cmd:"" help:"Initialize files for new LeetCode problem."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
