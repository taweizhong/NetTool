package httpServer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

var AddressText string
var PortText string

func GetAddressSetting() fyne.CanvasObject {
	addressLabel := widget.NewLabel("HTTP服务器IP:")

	addressSelect := widget.NewSelect([]string{"127.0.0.1", "43.143.144.180"}, func(value string) {
		AddressText = value
	})
	addressSelect.SetSelected("127.0.0.1")

	portLabel := widget.NewLabel("服务器端口:")
	portEntry := container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewEntry())

	linkLabel := widget.NewLabel("连接上下文:")
	linkEntry := container.New(layout.NewGridWrapLayout(fyne.NewSize(64, 32)), widget.NewEntry())

	listenBut := widget.NewButton("启动HTTP服务器", func() {

		log.Println(AddressText, PortText)
	})
	stopBut := widget.NewButton("停止", func() {
		log.Println("stop")
	})

	return widget.NewCard("", "HTTP连接信息",
		container.NewVBox(
			container.NewHBox(addressLabel, addressSelect, portLabel, portEntry),
			widget.NewSeparator(),
			container.NewHBox(linkLabel, linkEntry, listenBut, stopBut),
		),
	)
}
