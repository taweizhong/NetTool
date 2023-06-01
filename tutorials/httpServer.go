package tutorials

import (
	"NetTool/tutorials/httpServer"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func httpServerScreen(_ fyne.Window) fyne.CanvasObject {

	basicSetting := widget.NewCard("", "", container.NewVBox(httpServer.GetAddressSetting(), httpServer.GetHeartbeatSetting()))

	headBox := container.NewGridWithColumns(3, basicSetting, httpServer.GetSendSetting(), httpServer.GetClineSetting())

	return container.NewVBox(
		headBox,
		httpServer.GetBox3(),
		httpServer.GetSendInfo(),
	)
}
