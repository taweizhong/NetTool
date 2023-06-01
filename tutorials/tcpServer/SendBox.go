package tcpServer

import (
	"NetTool/data"
	"NetTool/handle"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

var flag = make(chan int)

func GetBox1() fyne.CanvasObject {
	SendEntry := widget.NewEntry()
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), SendEntry),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(144, 32)), widget.NewButton("发送", func() {
				handle.TcpServerSendData(SendEntry.Text)
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "TcpServerSend",
					Content: "Content:" + SendEntry.Text,
				})
				data.TcpSSet.SendMsgList = append(data.TcpSSet.SendMsgList, SendEntry.Text+"\n")
				SendEntry.Text = ""
				RenderSend()
			})),
		),
	)
}
func GetBox2() fyne.CanvasObject {
	SendEntry := widget.NewEntry()
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("发送数据"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(800, 32)), SendEntry),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 32)), widget.NewButton("循环发送", func() {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "循环发送中",
					Content: "",
				})
			})),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 32)), widget.NewButton("停止", func() {

			})),
		),
	)
}
func GetBox3() fyne.CanvasObject {
	Entry := widget.NewEntry()
	Entry.Text = "SERVER AUTO REPLY MESSAGE OK!"
	return widget.NewCard(
		"",
		"",
		container.NewHBox(
			widget.NewLabel("自动回复"),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 32)), Entry),
			widget.NewLabel("心跳信息："),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 32)), widget.NewEntry()),
		),
	)
}
func Send(msg string, bo chan int) {
	f := <-bo
	for f == 1 {
		handle.TcpServerSendData(msg)
		data.TcpSSet.SendMsgList = append(data.TcpSSet.SendMsgList, msg+"\n")
		RenderSend()
		time.Sleep(time.Second * 2)
	}
}
