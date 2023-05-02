package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/omurilo/aprenda-go/aprenda-go/ui"
	"github.com/omurilo/aprenda-go/exercises"
	"github.com/omurilo/aprenda-go/modules"
	"github.com/spf13/cobra"
)

func ListCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "list <modules|exercises>?",
		Short: "List all <modules and/or exercises>",
		Run: func(cmd *cobra.Command, args []string) {
			exs, err := exercises.List(infoFile)

			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}

			mds, err := modules.List(infoFile)

			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}

			if len(args) == 0 {
				ui.PrintExercisesList(os.Stdout, exs)
				ui.PrintModulesList(os.Stdout, mds)
			} else if args[0] == "exercises" {
				ui.PrintExercisesList(os.Stdout, exs)
			} else if args[0] == "modules" {
				ui.PrintModulesList(os.Stdout, mds)
			}
		},
	}
}
