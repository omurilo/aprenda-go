package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "aprenda-go",
		Short:         "Aprenda Go de forma fácil e completa: youtube.com/aprenda-go",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
	}

	rootCmd.AddCommand(HintCmd("info.toml"))
	rootCmd.AddCommand(ListCmd("info.toml"))
	rootCmd.AddCommand(RunCmd("info.toml"))
	rootCmd.AddCommand(VerifyCmd("info.toml"))
	rootCmd.AddCommand(WatchCmd("info.toml"))

	return rootCmd
}

func Execute(version string) {
	if err := NewRootCmd(version).Execute(); err != nil {
		os.Exit(1)
	}
}
