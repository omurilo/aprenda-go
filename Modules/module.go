package modules

import (
	"os"
	"regexp"
)

var notDoneRegex = regexp.MustCompile(`(?m)^\s*///?\s*I\s+AM\s+NOT\s+DONE`)

type Module struct {
	Name string
	Path string
	Mode string
	Hint string
}

func (m Module) State() State {
	data, err := os.ReadFile(m.Path)
	if err != nil {
		return Pending
	}

	if notDoneRegex.Match(data) {
		return Pending
	}
	return Done
}

type State int

const (
	Pending State = iota + 1
	Done
)

func (s State) String() string {
	return [...]string{"Pending", "Done"}[s-1]
}
