package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/omurilo/aprenda-go/exercises"
	"github.com/omurilo/aprenda-go/modules"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func RunCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "run <module|exercise> next | <exercise name>",
		Short: "Run a single <module|exercise>",
		Long: `example next pending exercise : aprenda-go run exercise next
example specific exercise : aprenda-go run exercise variables1`,
		Args:          cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var exercise exercises.Exercise
			var module modules.Module
			var err error

			if args[0] == "exercise" {
				if args[1] == "next" {
					exercise, err = exercises.NextPending(infoFile)
				} else {
					exercise, err = exercises.Find(args[1], infoFile)
				}

				spinner := RunSpinner(exercise.Name)

				if errors.Is(err, exercises.ErrExerciseNotFound) {
					color.White("No exercise found for '%s'", args[1])
					return err
				}

				result, err := exercise.Run()

				spinner.Close()

				if err != nil {
					color.Cyan("Failed to compile the exercise %s\n\n", result.Exercise.Path)
					color.White("Check the output below: \n\n")
					color.Red(result.Err)
					color.Red(result.Out)
					color.Yellow("If you feel stuck, ask a hint by executing `aprenda-go hint %s`", result.Exercise.Name)
				} else {
					color.Green("Congratulations!\n\n")
					color.Green("Here is the output of your program:\n\n")
					color.Cyan(result.Out)
					if result.Exercise.State() == exercises.Pending {
						color.White("Remove the 'I AM NOT DONE' from the file to keep going\n")
						return fmt.Errorf("exercise is still pending")
					}
				}
			} else if args[0] == "module" {
				if args[1] == "next" {
					module, err = modules.NextPending(infoFile)
				} else {
					module, err = modules.Find(args[1], infoFile)
				}

				spinner := RunSpinner(module.Name)

				if errors.Is(err, modules.ErrModuleNotFound) {
					color.White("No module found for '%s'", args[1])
					return err
				}

				result, err := module.Run()

				spinner.Close()

				if err != nil {
					color.Cyan("Failed to compile the module %s\n\n", result.Module.Path)
					color.White("Check the output below: \n\n")
					color.Red(result.Err)
					color.Red(result.Out)
					color.Yellow("If you feel stuck, ask a hint by executing `aprenda-go hint %s`", result.Module.Name)
				} else {
					color.Green("Congratulations!\n\n")
					color.Green("Here is the output of your program:\n\n")
					color.Cyan(result.Out)
					if result.Module.State() == modules.Pending {
						color.White("Remove the 'I AM NOT DONE' from the file to keep going\n")
						return fmt.Errorf("module is still pending")
					}
				}
			}

			return err
		},
	}
}

func RunSpinner(name string) *progressbar.ProgressBar {
	spinner := progressbar.NewOptions(
		-1, // a negative number makes turns the progress bar into a spinner
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription(color.WhiteString("Running: %s", name)),
		progressbar.OptionOnCompletion(func() {
			color.White("\nRunning complete!\n\n")
		}),
	)
	go func() {
		for x := 0; x < 100; x++ {
			spinner.Add(1) // nolint
			time.Sleep(250 * time.Millisecond)
		}
	}()

	return spinner
}
