package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"ecb-system/internal/services"
)

func StartUI(service *services.BreakerService) {
	a := app.New()
	w := a.NewWindow("ECB System")
	statusLabel := widget.NewLabel("Status: Normal")

	btnTestLeak := widget.NewButton("Simulate Leak", func() {
		service.Sensor.Write(true)
		statusLabel.SetText("Status: ⚠️ Leakage detected!")
	})

	btnReset := widget.NewButton("Reset", func() {
		service.Sensor.Write(false)
		statusLabel.SetText("Status: ✅ Normal")
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Electric Circuit Breaker System"),
		statusLabel,
		btnTestLeak,
		btnReset,
	))

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
