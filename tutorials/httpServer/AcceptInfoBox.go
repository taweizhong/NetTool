package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var AcceptNumbs = binding.NewString()
var AcceptCharNumbs = binding.NewString()

func GetAcceptInfo() fyne.CanvasObject {
	AcceptNumbs.Set("0")
	AcceptCharNumbs.Set("0")
	infoBox := container.NewHBox(
		widget.NewLabel("接受信息"),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewButton("清除日志", func() {})),
		widget.NewLabel("接受数量:"),
		widget.NewLabelWithData(AcceptNumbs),
		widget.NewLabel("接受字符数:"),
		widget.NewLabelWithData(AcceptCharNumbs),
	)

	AcceptDataBoxEntry := widget.NewMultiLineEntry()

	return widget.NewCard(
		"",
		"",
		container.NewVBox(
			infoBox,
			container.New(layout.NewGridWrapLayout(fyne.NewSize(1180, 150)), AcceptDataBoxEntry),
		),
	)
}
