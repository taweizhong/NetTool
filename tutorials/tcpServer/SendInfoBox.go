package tcpServer

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var SendNumbs = binding.NewString()
var SendCharNumbs = binding.NewString()
var SendDataBoxEntry *widget.Entry

func GetSendInfo() fyne.CanvasObject {
	SendNumbs.Set("0")
	SendCharNumbs.Set("0")
	infoBox := container.NewHBox(
		widget.NewLabel("发送信息"),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewButton("清除日志", func() {
			data.TcpSSet.SendMsgList = data.TcpSSet.SendMsgList[:0]
			RenderSend()
		})),
		widget.NewLabel("发送数量:"),
		widget.NewLabelWithData(SendNumbs),
		widget.NewLabel("发送字符数:"),
		widget.NewLabelWithData(SendCharNumbs),
	)

	SendDataBoxEntry = widget.NewMultiLineEntry()
	RenderSend()

	return widget.NewCard(
		"",
		"",
		container.NewVBox(
			infoBox,
			container.New(layout.NewGridWrapLayout(fyne.NewSize(1040, 130)), SendDataBoxEntry),
		),
	)
}
func RenderSend() {
	str := ""
	for i := 0; i < len(data.TcpSSet.SendMsgList); i++ {
		str += data.TcpSSet.SendMsgList[i]
	}

	SendDataBoxEntry.SetText(str)

	num := len(data.TcpSSet.SendMsgList)
	SendNumbs.Set(strconv.Itoa(num))
}
