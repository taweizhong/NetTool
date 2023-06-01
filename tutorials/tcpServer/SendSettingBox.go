package tcpServer

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GetSendSetting() fyne.CanvasObject {
	sendSettingCheck1 := widget.NewCheck("当客户黹连接后发送HELLO", func(value bool) {
		data.TcpSSet.IsSendHello = value
		log.Println("Check set to", value)
	})
	sendSettingCheck2 := widget.NewCheck("关闭日志显示提高性能", func(value bool) {

		log.Println("Check set to", value)
	})
	sendSettingCheck3 := widget.NewCheck("接收到信息后中断连接", func(value bool) {
		log.Println("Check set to", value)
	})
	Label1 := widget.NewLabel("宇符集编码:")
	sendSettingSelect := widget.NewSelect([]string{"utf-8", "gbk"}, func(s string) {
		data.TcpSSet.CharacterSet = s
	})
	sendSettingSelect.SetSelected("utf-8")

	sendSettingCheck5 := widget.NewCheck("收到客户端信息自动回复", func(value bool) {
		data.TcpSSet.IsAutomaticReply = value
		log.Println("Check set to", value)
	})
	sendSettingCheck6 := widget.NewCheck("HEX(十六进制)发送", func(value bool) {
		log.Println("Check set to", value)
	})
	sendSettingCheck7 := widget.NewCheck("是否循环发送", func(value bool) {
		log.Println("Check set to", value)
	})
	Label2 := widget.NewLabel("宿环次数：")
	Entry1 := container.New(layout.NewGridWrapLayout(fyne.NewSize(40, 32)), widget.NewEntry())
	Label3 := widget.NewLabel("间隔(ms)：")
	Entry2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(40, 32)), widget.NewEntry())

	return widget.NewCard("", "收发设置",
		container.NewVBox(
			container.NewHBox(sendSettingCheck1, layout.NewSpacer(), sendSettingCheck2),
			container.NewHBox(sendSettingCheck3, layout.NewSpacer(), container.NewHBox(Label1, sendSettingSelect)),
			container.NewHBox(sendSettingCheck5, layout.NewSpacer(), sendSettingCheck6),
			container.NewHBox(sendSettingCheck7, layout.NewSpacer()),
			container.NewHBox(container.NewHBox(Label2, Entry1), container.NewHBox(Label3, Entry2)),
		),
	)
}
