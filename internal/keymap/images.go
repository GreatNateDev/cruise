package keymap

import (
	"reflect"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type ImagesMap struct {
	Remove key.Binding
	Prune  key.Binding
	Push   key.Binding
	Pull   key.Binding
	Build  key.Binding
	Layers key.Binding
	Exit   key.Binding
}

func NewImagesMap() ImagesMap {
	m := config.Cfg.Keybinds.Images
	return ImagesMap{
		Remove: key.NewBinding(
			key.WithKeys(m.Remove),
			key.WithHelp(m.Remove, "remove"),
		),
		Prune: key.NewBinding(
			key.WithKeys(m.Prune),
			key.WithHelp(m.Prune, "prune"),
		),
		Pull: key.NewBinding(
			key.WithKeys(m.Pull),
			key.WithHelp(m.Pull, "pull"),
		),
		Push: key.NewBinding(
			key.WithKeys(m.Push),
			key.WithHelp(m.Push, "push"),
		),
		Build: key.NewBinding(
			key.WithKeys(m.Build),
			key.WithHelp(m.Build, "build"),
		),
		Layers: key.NewBinding(
			key.WithKeys(m.Build),
			key.WithHelp(m.Build, "layers"),
		),
	}
}

func (m ImagesMap) Bindings() []key.Binding {
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
