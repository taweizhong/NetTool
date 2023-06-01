package tcpClient

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

var HeartRate string

func GetHeartbeatSetting() fyne.CanvasObject {
	heartbeatCheck := widget.NewCheck("启用心跳", func(value bool) {
		log.Println("Check set to", value)
	})
	heartbeatFps := widget.NewSelect([]string{"每30秒", "每一分钟"}, func(value string) {
		HeartRate = value
	})
	heartbeatFps.SetSelected("每30秒")

	return widget.NewCard("", "心跳设置",
		container.NewHBox(
			heartbeatCheck,
			widget.NewLabel("心跳字符"),
			widget.NewEntry(),
			widget.NewLabel("心跳频率"),
			heartbeatFps,
		),
	)

}
