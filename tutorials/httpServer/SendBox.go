package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetBox1() fyne.CanvasObject {
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), widget.NewEntry()),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(144, 32)), widget.NewButton("发送", func() {})),
		),
	)
}
func GetBox2() fyne.CanvasObject {
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), widget.NewEntry()),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 32)), widget.NewButton("循环发送", func() {})),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 32)), widget.NewButton("发送", func() {})),
		),
	)
}
func GetBox3() fyne.CanvasObject {
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("自动回复"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 32)), widget.NewEntry()),
		),
	)

}
