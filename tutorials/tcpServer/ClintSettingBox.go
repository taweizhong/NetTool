package tcpServer

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GetClineSetting() fyne.CanvasObject {
	clientMultiLineEntry := widget.NewMultiLineEntry()
	for i := 0; i < len(data.TcpSSet.LinkList); i++ {
		clientMultiLineEntry.Text += data.TcpSSet.LinkList[i] + "\n"
	}
	but1 := widget.NewButton("全选", func() {
		log.Println("全选")
	})
	but2 := widget.NewButton("断开", func() {
		log.Println("断开")
	})
	but3 := widget.NewButton("全断开", func() {
		log.Println("全断开")
	})
	return widget.NewCard("", "客户端",
		container.NewHBox(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 190)), clientMultiLineEntry),
			container.NewVBox(
				container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 32)), but1),
				layout.NewSpacer(),
				container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 32)), but2),
				layout.NewSpacer(),
				container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 32)), but3),
			),
		),
	)
}
