package cli

import (
	"github.com/rs/zerolog"
	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func InitRootCmd(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		verbose, err := oops.Wrap2(cmd.Flags().GetCount("verbose"))
		if err != nil {
			return err
		}
		quiet, err := oops.Wrap2(cmd.Flags().GetCount("quiet"))
		if err != nil {
			return err
		}
		InitLogging(zerolog.Level(verbose - quiet))
		// err = oops.Wrap(viper.ReadInConfig())
		// if err != nil {
		// 	return err
		// }
		return nil
	}
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.DisableAutoGenTag = true
	cmd.PersistentFlags().CountP("quiet", "q", "")
	cmd.PersistentFlags().CountP("verbose", "v", "")
	cmd.AddCommand(CmdDocgen())
	return cmd
}
