package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type toolbarLabel struct {
	*widget.Label
}

func (t toolbarLabel) ToolbarObject() fyne.CanvasObject {
	return t.Label
}

func NewToolbarLabel(label string) widget.ToolbarItem {
	l := widget.NewLabelWithStyle(label, fyne.TextAlignCenter, fyne.TextStyle{Monospace: true})
	return &toolbarLabel{l}
}