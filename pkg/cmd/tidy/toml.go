package tidy

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func CmdToml() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "toml",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fpath := args[0]
			inPlace, err := oops.Wrap2(cmd.Flags().GetBool("in-place"))
			if err != nil {
				return err
			}
			data, err := oops.Wrap2(os.ReadFile(fpath))
			if err != nil {
				return err
			}
			data, err = TomlSort(data)
			if err != nil {
				return err
			}
			data, err = TaploFormat(data)
			if err != nil {
				return err
			}
			text := FixSchemaComments(data)
			if inPlace {
				// write data to fpath
				err = os.WriteFile(fpath, []byte(text), 0644)
				if err != nil {
					return err
				}
			} else {
				print(text)
			}
			return err
		},
	}
	cmd.Flags().Bool("in-place", true, "overwrite the original input file with changes")
	return cmd
}

func TomlSort(data []byte) (result []byte, err error) {
	cmd := exec.Command("toml-sort", "--all")
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stderr = os.Stderr
	stdout, err := oops.Wrap2(cmd.StdoutPipe())
	if err != nil {
		return nil, err
	}
	defer stdout.Close()
	err = oops.Wrap(cmd.Start())
	if err != nil {
		return nil, err
	}
	defer cmd.Wait()
	result, err = oops.Wrap2(io.ReadAll(stdout))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		// TODO: remove this workaround when <https://github.com/pappasam/toml-sort/pull/83> is fixed.
		return data, nil
	}
	return result, nil
}

func TaploFormat(data []byte) (result []byte, err error) {
	cmd := exec.Command("taplo", "format", "-")
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stderr = os.Stderr
	stdout, err := oops.Wrap2(cmd.StdoutPipe())
	if err != nil {
		return nil, err
	}
	defer stdout.Close()
	err = oops.Wrap(cmd.Start())
	if err != nil {
		return nil, err
	}
	defer cmd.Wait()
	result, err = oops.Wrap2(io.ReadAll(stdout))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FixSchemaComments(data []byte) (result string) {
	text := string(data)
	for line := range strings.SplitSeq(text, "\n") {
		if strings.HasPrefix(line, "# :schema ") {
			line = strings.Replace(line, "# :schema ", "#:schema ", 1)
		}
		result += line + "\n"
	}
	return result
}
