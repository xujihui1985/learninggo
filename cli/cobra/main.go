package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xujihui1985/learninggo/cli/cobra/args"
)

var (
	F = struct {
		Persistent string
		Local      string
	}{}
)

func main() {
	cobra.OnInitialize()

	rootCmd := &cobra.Command{
		Use:   "cobra",
		Short: "",
		Long: `this is long description 
		another line`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// PersistentFlags can be used from sub1
	rootCmd.PersistentFlags().StringVar(&F.Persistent, "persistent", "", "the name of the flag")
	viper.BindPFlag("persistent", rootCmd.PersistentFlags().Lookup("persistent"))
	viper.BindEnv("id", "BASEMENT_ID")

	// LocalFlags can not be used from sub1
	// when cobra sub1 -
	rootCmd.Flags().StringVar(&F.Local, "local", "", "the name of the flag")

	subCmd := &cobra.Command{
		Use:   "sub1",
		Short: "s",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(F)
		},
	}

	errorCmd := &cobra.Command{
		Use: "error",
		Run: func(cmd *cobra.Command, args []string) {
			panic("error")
		},
	}

	vaildArgs := &cobra.Command{
		Use:        "validargs",
		Example:    `cobra validargs foo`,
		SuggestFor: []string{"foo", "bar"},
		Args:       args.Combine(cobra.ExactArgs(1), cobra.OnlyValidArgs, args.CustomValidator),
		ValidArgs:  []string{"foo", "bar"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	// not working
	completionCmd := &cobra.Command{
		Use:       "completion",
		Example:   `cobra completion`,
		ValidArgs: []string{"bash", "zsh"},
		Args:      args.Combine(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "bash" {
				rootCmd.GenBashCompletion(os.Stdout)
			}
			if args[0] == "zsh" {
				rootCmd.GenZshCompletion(os.Stdout)
			}
		},
	}

	viperCmd := &cobra.Command{
		Use: "viper",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.Get("id"))
		},
	}

	rootCmd.AddCommand(subCmd, errorCmd, vaildArgs, completionCmd, viperCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("we have error when execute command", err)
		os.Exit(1)
	}

}
