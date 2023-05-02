package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/omurilo/aprenda-go/exercises"
	"github.com/omurilo/aprenda-go/modules"
	"github.com/spf13/cobra"
)

func HintCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "hint <modules|exercises> <name>",
		Short: "Get a hint for an exercise or module by that name",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			var exercise exercises.Exercise
			var module modules.Module
			var err error

			if args[0] == "exercises" {
				if args[1] == "next" {
					exercise, err = exercises.NextPending(infoFile)
				} else {
					exercise, err = exercises.Find(args[1], infoFile)
				}

				if err != nil {
					color.Red(err.Error())
					os.Exit(1)
				}

				color.Yellow(exercise.Hint)
			} else if args[0] == "modules" {
				if args[1] == "next" {
					module, err = modules.NextPending(infoFile)
				} else {
					module, err = modules.Find(args[1], infoFile)
				}

				if err != nil {
					color.Red(err.Error())
					os.Exit(1)
				}

				color.Yellow(module.Hint)
			}
		},
	}
}
