package ui

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/omurilo/aprenda-go/exercises"
	"github.com/omurilo/aprenda-go/modules"
)

func PrintExercisesList(o io.Writer, exs []exercises.Exercise) {
	t := table.NewWriter()
	t.SetOutputMirror(o)
	rowConfigAutoMerge := table.RowConfig{
		AutoMerge:      true,
		AutoMergeAlign: text.AlignCenter,
	}
	t.AppendHeader(table.Row{"Exercises", "Exercises", "Exercises"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"Name", "Path", "State"})
	for _, ex := range exs {
		t.AppendRow(table.Row{ex.Name, ex.Path, ex.State()})
	}
	t.Render()
}

func PrintModulesList(o io.Writer, mds []modules.Module) {
	t := table.NewWriter()
	t.SetOutputMirror(o)
	rowConfigAutoMerge := table.RowConfig{
		AutoMerge:      true,
		AutoMergeAlign: text.AlignCenter,
	}
	t.AppendHeader(table.Row{"Modules", "Modules", "Modules"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"Name", "Path", "State"})
	for _, md := range mds {
		t.AppendRow(table.Row{md.Name, md.Path, md.State()})
	}
	t.Render()
}
