package modules

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var ErrModuleNotFound = errors.New("module not found")
var ErrNoPendingModules = errors.New("no pending modules")

type Info struct {
	Modules []Module
}

func List(infoFile string) ([]Module, error) {
	var info Info

	data, err := os.ReadFile(infoFile)
	if err != nil {
		return info.Modules, err
	}

	if err := toml.Unmarshal(data, &info); err != nil {
		return info.Modules, err
	}

	return info.Modules, nil
}

func NextPending(infoFile string) (Module, error) {
	allModules, err := List(infoFile)
	if err != nil {
		return Module{}, err
	}

	for _, module := range allModules {
		if module.State() == Pending {
			return module, nil
		}
	}

	return Module{}, ErrNoPendingModules
}

func Find(module string, infoFile string) (Module, error) {
	exs, err := List(infoFile)
	if err != nil {
		return Module{}, err
	}

	for _, ex := range exs {
		if ex.Name == module {
			return ex, nil
		}
	}

	return Module{}, ErrModuleNotFound
}
