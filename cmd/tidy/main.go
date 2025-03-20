package main

import (
	"github.com/liblaf/lollipop/pkg/cli"
	"github.com/liblaf/lollipop/pkg/cmd/tidy"
	"github.com/rs/zerolog"
	"github.com/samber/oops"
)

func main() {
	cli.InitLogging(zerolog.InfoLevel)
	cmd := tidy.CmdRoot()
	if err := oops.Wrap(cmd.Execute()); err != nil {
		cli.PrintFatalError(err)
	}
}
