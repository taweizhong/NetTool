package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	image := canvas.NewImageFromFile("./data/assets/home.png")
	image.FillMode = canvas.ImageFillStretch
	return image
}
