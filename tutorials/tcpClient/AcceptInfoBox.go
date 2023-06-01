package tcpClient

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
			data.TcpCSet.AcceptMsgList = data.TcpCSet.AcceptMsgList[:0]
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
			container.New(layout.NewGridWrapLayout(fyne.NewSize(1180, 110)), AcceptDataBoxEntry),
		),
	)
}
func RenderAccept() {
	str := ""
	for i := 0; i < len(data.TcpCSet.AcceptMsgList); i++ {
		str += data.TcpCSet.AcceptMsgList[i]
	}
	AcceptDataBoxEntry.SetText(str)

	num := len(data.TcpCSet.AcceptMsgList)
	AcceptNumbs.Set(strconv.Itoa(num))
}
