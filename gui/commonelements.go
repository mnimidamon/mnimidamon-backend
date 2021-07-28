package gui

import "fyne.io/fyne/v2"

type ContentGetter interface {
	GetContent() fyne.CanvasObject
}
