package cli

import (
	"os"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func CmdDocgen() *cobra.Command {
	cmd := &cobra.Command{
		Use: "docgen",
	}
	cmd.AddCommand(CmdDocgenMan(), CmdDocgenMarkdown(), CmdDocgenReST(), CmdDocgenYaml())
	return cmd
}

func CmdDocgenMan() *cobra.Command {
	cmd := &cobra.Command{
		Use: "man",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			c := cmd.Root()
			output, err := oops.Wrap2(cmd.Flags().GetString("output"))
			if err != nil {
				return err
			}
			err = oops.Wrap(os.MkdirAll(output, 0755))
			if err != nil {
				return err
			}
			header := &doc.GenManHeader{}
			err = oops.Wrap(doc.GenManTree(c, header, output))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("output", "o", "docs/man/", "")
	return cmd
}

func CmdDocgenMarkdown() *cobra.Command {
	cmd := &cobra.Command{
		Use: "markdown",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			c := cmd.Root()
			output, err := oops.Wrap2(cmd.Flags().GetString("output"))
			if err != nil {
				return err
			}
			err = oops.Wrap(os.MkdirAll(output, 0755))
			if err != nil {
				return err
			}
			err = oops.Wrap(doc.GenMarkdownTree(c, output))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("output", "o", "docs/markdown/", "")
	return cmd
}

func CmdDocgenReST() *cobra.Command {
	cmd := &cobra.Command{
		Use: "rest",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			c := cmd.Root()
			output, err := oops.Wrap2(cmd.Flags().GetString("output"))
			if err != nil {
				return err
			}
			err = oops.Wrap(os.MkdirAll(output, 0755))
			if err != nil {
				return err
			}
			err = oops.Wrap(doc.GenReSTTree(c, output))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("output", "o", "docs/rest/", "")
	return cmd
}

func CmdDocgenYaml() *cobra.Command {
	cmd := &cobra.Command{
		Use: "yaml",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			c := cmd.Root()
			output, err := oops.Wrap2(cmd.Flags().GetString("output"))
			if err != nil {
				return err
			}
			err = oops.Wrap(os.MkdirAll(output, 0755))
			if err != nil {
				return err
			}
			err = oops.Wrap(doc.GenYamlTree(c, output))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("output", "o", "docs/yaml/", "")
	return cmd
}
