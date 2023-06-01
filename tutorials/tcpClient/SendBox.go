package tcpClient

import (
	"NetTool/data"
	"NetTool/handle"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GetBox1() fyne.CanvasObject {
	Entry := widget.NewEntry()
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), Entry),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(144, 32)), widget.NewButton("发送", func() {
				handle.TcpClientSendData(Entry.Text)
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "TcpClientSend",
					Content: "Content:" + Entry.Text,
				})
				data.TcpCSet.SendMsgList = append(data.TcpCSet.SendMsgList, Entry.Text+"\n")
				Entry.Text = ""
				RenderSend()
			})),
		),
	)
}
func GetBox2(w fyne.Window) fyne.CanvasObject {
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), widget.NewEntry()),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 32)), widget.NewButton("导入文本文件", func() {
				fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if reader == nil {
						log.Println("Cancelled")
						return
					}
					log.Println(reader)
				}, w)
				fd.Show()
			},
			),
			),
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
