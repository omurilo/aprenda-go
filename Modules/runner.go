package modules

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Result struct {
	Module Module
	Out    string
	Err    string
}

func (m Module) Run() (Result, error) {
	args := BuildArgs(m)
	cmd := exec.Command("go", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return Result{Module: m, Out: stdout.String(), Err: stderr.String()}, err
}

func BuildArgs(m Module) []string {
	args := []string{}
	if m.Mode == "compile" {
		args = append(args, "run")
	} else {
		args = append(args, "test", "-v")
	}

	args = append(args, fmt.Sprintf("./%s", m.Path))
	return args
}
