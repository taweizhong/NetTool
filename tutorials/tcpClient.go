package tutorials

import (
	"NetTool/tutorials/tcpClient"
	"NetTool/tutorials/tcpServer"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func tcpClientScreen(w fyne.Window) fyne.CanvasObject {
	basicSetting := widget.NewCard("", "", container.NewVBox(tcpClient.GetAddressSetting(), tcpClient.GetHeartbeatSetting()))

	headBox := container.NewHBox(basicSetting, tcpClient.GetSendSetting())

	return container.NewVBox(
		headBox,
		tcpClient.GetBox1(), tcpClient.GetBox2(w), tcpServer.GetBox2(), tcpClient.GetBox3(),
		tcpClient.GetSendInfo(), tcpClient.GetAcceptInfo(),
	)
}
