package tcpClient

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GetSendSetting() fyne.CanvasObject {
	sendSettingCheck1 := widget.NewCheck("发送信息添加结束宇符", func(value bool) {
		log.Println("Check set to", value)
	})

	Label1 := widget.NewLabel("结束字符设置：")
	sendSettingSelect1 := widget.NewSelect([]string{"0x03", "0x02", "0x01"}, func(s string) {
	})
	sendSettingSelect1.SetSelected("0x03")

	Label2 := widget.NewLabel("宇符集编码：              ")
	sendSettingSelect2 := widget.NewSelect([]string{"utf-8", "gbk"}, func(s string) {
	})
	sendSettingSelect2.SetSelected("utf-8")

	sendSettingCheck2 := widget.NewCheck("添加字符串信息后开始", func(value bool) {
		log.Println("Check set to", value)
	})

	Label3 := widget.NewLabel("长度位数:            ")
	sendSettingSelect3 := widget.NewSelect([]string{"4", "8"}, func(s string) {
	})
	sendSettingSelect3.SetSelected("4")

	LabelEX := widget.NewLabel("HEX(十六进制发送: ")
	sendSettingSelectHEX := widget.NewSelect([]string{"N", "A"}, func(s string) {
	})
	sendSettingSelectHEX.SetSelected("N")

	sendSettingCheck3 := widget.NewCheck("每次发送后断开连接", func(value bool) {
		log.Println("Check set to", value)
	})

	sendSettingCheck4 := widget.NewCheck("是否循环发送", func(value bool) {
		log.Println("Check set to", value)
	})
	Label4 := widget.NewLabel("宿环次数：                   ")
	Entry1 := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 32)), widget.NewEntry())
	Label5 := widget.NewLabel("间隔(ms)：")
	Entry2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 32)), widget.NewEntry())

	return widget.NewCard("", "收发设置",
		container.NewGridWithColumns(3,
			sendSettingCheck1, container.NewHBox(Label1, sendSettingSelect1), container.NewHBox(Label2, sendSettingSelect2),
			sendSettingCheck2, container.NewHBox(Label3, sendSettingSelect3), container.NewHBox(LabelEX, sendSettingSelectHEX),
			sendSettingCheck3,
			sendSettingCheck4, container.NewHBox(Label4, Entry1), container.NewHBox(Label5, Entry2),
		),
	)
}
