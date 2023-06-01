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

var AcceptNumbs = binding.NewString()
var AcceptCharNumbs = binding.NewString()
var AcceptDataBoxEntry *widget.Entry

func GetAcceptInfo() fyne.CanvasObject {
	AcceptNumbs.Set("0")
	AcceptCharNumbs.Set("0")
	infoBox := container.NewHBox(
		widget.NewLabel("接受信息"),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewButton("清除日志", func() {
			data.TcpSSet.AcceptMsgList = data.TcpSSet.AcceptMsgList[:0]
			RenderAccept()
		})),
		widget.NewLabel("接受数量:"),
		widget.NewLabelWithData(AcceptNumbs),
		widget.NewLabel("接受字符数:"),
		widget.NewLabelWithData(AcceptCharNumbs),
	)

	AcceptDataBoxEntry = widget.NewMultiLineEntry()
	RenderAccept()

	return widget.NewCard(
		"",
		"",
		container.NewVBox(
			infoBox,
			container.New(layout.NewGridWrapLayout(fyne.NewSize(1040, 130)), AcceptDataBoxEntry),
		),
	)
}
func RenderAccept() {
	str := ""
	for i := 0; i < len(data.TcpSSet.AcceptMsgList); i++ {
		str += data.TcpSSet.AcceptMsgList[i]
	}
	AcceptDataBoxEntry.SetText(str)

	num := len(data.TcpSSet.AcceptMsgList)
	AcceptNumbs.Set(strconv.Itoa(num))
}
