package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetClineSetting() fyne.CanvasObject {
	clientMultiLineEntry := widget.NewMultiLineEntry()
	return widget.NewCard("", "客户端",
		container.NewHBox(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(330, 190)), clientMultiLineEntry),
		),
	)
}
