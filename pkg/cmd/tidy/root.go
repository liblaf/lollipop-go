package tidy

import (
	"github.com/liblaf/lollipop/pkg/cli"
	"github.com/spf13/cobra"
)

func CmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tidy",
	}
	cmd = cli.InitRootCmd(cmd)
	cmd.AddCommand(CmdToml())
	return cmd
}
