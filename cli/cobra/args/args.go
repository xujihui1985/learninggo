package args

import (
	"errors"

	"github.com/spf13/cobra"
)

func Combine(funcs ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range funcs {
			err := f(cmd, args)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func CustomValidator(cmd *cobra.Command, args []string) error {
	if args[0] != "foo" {
		return errors.New("the first args must be foo")
	}
	return nil
}
