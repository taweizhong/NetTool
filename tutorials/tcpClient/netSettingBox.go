package tcpClient

import (
	"NetTool/handle"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"net"
)

var AddressText string
var PortText string
var C = make(chan net.Conn)

func GetAddressSetting() fyne.CanvasObject {
	addressLabel := widget.NewLabel("服务器IP地址")

	addressSelect := widget.NewSelect([]string{"127.0.0.1", "43.143.144.180"}, func(value string) {
		AddressText = value
	})
	addressSelect.SetSelected("127.0.0.1")
	portLabel := widget.NewLabel("端口")
	portEntry := widget.NewEntry()
	EntryBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 32)), portEntry)
	listenBut := widget.NewButton("连接", func() {
		PortText = portEntry.Text
		go handle.TcpClient(AddressText+":"+PortText, C)
	})
	stopBut := widget.NewButton("断开", func() {
		log.Println("stop")
		go handle.TcpClientClose(C)
	})

	return widget.NewCard("", "网络设置",
		container.NewVBox(
			container.NewHBox(addressLabel, addressSelect, portLabel, EntryBox),
			widget.NewSeparator(),
			container.NewHBox(layout.NewSpacer(), listenBut, layout.NewSpacer(), stopBut, layout.NewSpacer()),
		),
	)
}
