package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/omurilo/aprenda-go/aprenda-go/ui"
	"github.com/omurilo/aprenda-go/exercises"
	"github.com/omurilo/aprenda-go/modules"
)

func PrintHint(infoFile string) {
	ClearScreen()
	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
	}
	color.Yellow(exercise.Hint)
	module, err := modules.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next modules")
	}
	color.Yellow(module.Hint)
}

func PrintList(infoFile string) {
	ClearScreen()
	exs, err := exercises.List(infoFile)
	if err != nil {
		color.Red("Failed to list exercises")
	}
	ui.PrintExercisesList(os.Stdout, exs)
	mds, err := modules.List(infoFile)
	if err != nil {
		color.Red("Failed to list exercises")
	}
	ui.PrintModulesList(os.Stdout, mds)
}

func RunNextExercise(infoFile string) {
	ClearScreen()
	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
	}

	result, err := exercise.Run()
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
			color.Red("exercise is still pending")
		}
	}
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	}
}
