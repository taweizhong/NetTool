package tcpServer

import (
	"NetTool/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

var HeartRate string

func GetHeartbeatSetting() fyne.CanvasObject {
	heartbeatCheck := widget.NewCheck("启用心跳", func(value bool) {
		data.TcpSSet.IsHeartBeat = value
		log.Println("Check set to", value)
	})
	heartbeatFps := widget.NewSelect([]string{"每30秒", "每一分钟"}, func(value string) {
		data.TcpSSet.HeartFps = value
		HeartRate = value
	})
	heartbeatFps.SetSelected("每30秒")

	return widget.NewCard("", "心跳设置",
		container.NewHBox(
			heartbeatCheck,
			widget.NewLabel("心跳频率"),
			heartbeatFps,
		),
	)

}
