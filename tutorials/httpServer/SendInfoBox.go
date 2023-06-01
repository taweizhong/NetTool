package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var SendNumbs = binding.NewString()
var SendCharNumbs = binding.NewString()

func GetSendInfo() fyne.CanvasObject {
	SendNumbs.Set("0")
	SendCharNumbs.Set("0")
	infoBox := container.NewHBox(
		widget.NewLabel("发送/接受信息"),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewButton("清除日志", func() {})),
		widget.NewLabel("发送数量:"),
		widget.NewLabelWithData(SendNumbs),
		widget.NewLabel("发送字符数:"),
		widget.NewLabelWithData(SendCharNumbs),
	)

	sendDataBoxEntry := widget.NewMultiLineEntry()

	return widget.NewCard(
		"",
		"",
		container.NewVBox(
			infoBox,
			container.New(layout.NewGridWrapLayout(fyne.NewSize(1180, 450)), sendDataBoxEntry),
		),
	)
}
