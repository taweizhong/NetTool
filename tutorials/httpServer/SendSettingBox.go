package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GetSendSetting() fyne.CanvasObject {
	Label1 := widget.NewLabel("宇符集编码：")
	sendSettingSelect := widget.NewSelect([]string{"utf-8", "gbk"}, func(s string) {
	})
	sendSettingSelect.SetSelected("utf-8")

	sendSettingCheck5 := widget.NewCheck("收到客户端信息自动回复", func(value bool) {
		log.Println("Check set to", value)
	})
	sendSettingCheck6 := widget.NewCheck("是否显示接收日志", func(value bool) {
		log.Println("Check set to", value)
	})

	return widget.NewCard("", "收发设置",
		container.NewVBox(
			container.NewHBox(sendSettingCheck5, layout.NewSpacer(), container.NewHBox(Label1, sendSettingSelect)),
			container.NewHBox(sendSettingCheck6),
		),
	)
}
