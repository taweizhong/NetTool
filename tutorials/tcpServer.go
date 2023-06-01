package tutorials

import (
	"NetTool/tutorials/tcpServer"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func tcpServerScreen(_ fyne.Window) fyne.CanvasObject {

	basicSetting := widget.NewCard("", "", container.NewVBox(tcpServer.GetAddressSetting(), tcpServer.GetHeartbeatSetting()))

	headBox := container.NewGridWithColumns(3, basicSetting, tcpServer.GetSendSetting(), tcpServer.GetClineSetting())

	return container.NewVBox(
		headBox,
		tcpServer.GetBox1(), tcpServer.GetBox2(), tcpServer.GetBox3(),
		tcpServer.GetSendInfo(), tcpServer.GetAcceptInfo(),
	)
}
