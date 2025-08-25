package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type NetMap struct {
	Remove      key.Binding
	Prune       key.Binding
	ShowDetails key.Binding
	ExitDetails key.Binding
}

func NewNetMap() NetMap {
	m := config.Cfg.Keybinds.Network
	return NetMap{
		Remove: key.NewBinding(
			key.WithKeys(m.Remove),
			key.WithHelp(m.Remove, "remove"),
		),
		Prune: key.NewBinding(
			key.WithKeys(m.Prune),
			key.WithHelp(m.Prune, "prune"),
		),
		ShowDetails: key.NewBinding(
			key.WithKeys(m.ShowDetails),
			key.WithHelp(m.ShowDetails, "show details"),
		),
		ExitDetails: key.NewBinding(
			key.WithKeys(m.ExitDetails),
			key.WithHelp(m.ExitDetails, "exit details"),
		),
	}
}

func (m NetMap) Bindings() []key.Binding {
	var bindings []key.Binding

	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if binding, ok := field.Interface().(key.Binding); ok {
			bindings = append(bindings, binding)
		}
	}

	return bindings
}
