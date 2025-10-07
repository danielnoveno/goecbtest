package ui

import (
	"goecbtest/internal/services"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartUI(service *services.BreakerService) {
	a := app.New()
	w := a.NewWindow("ECB System")

	statusLabel := widget.NewLabel("Status: ✅ Normal")
	logsView := widget.NewMultiLineEntry()
	logsView.SetReadOnly(true)
	logsView.SetText("System started...\n")

	btnTestLeak := widget.NewButton("Simulate Leak", func() {
		service.Sensor.Write(true)
		statusLabel.SetText("Status: ⚠️ Leakage detected!")
	})

	btnReset := widget.NewButton("Reset", func() {
		service.Sensor.Write(false)
		statusLabel.SetText("Status: ✅ Normal")
	})

	layout := container.NewVBox(
		widget.NewLabel("Electric Circuit Breaker System"),
		statusLabel,
		container.NewHBox(btnTestLeak, btnReset),
		widget.NewLabel("Event Log:"),
		logsView,
	)

	w.SetContent(layout)
	w.Resize(fyne.NewSize(500, 400))

	// Goroutine buat update log dari channel
	go func() {
		for msg := range service.Logs {
			current := logsView.Text
			newLog := time.Now().Format("15:04:05") + " " + msg + "\n"
			logsView.SetText(current + newLog)
		}
	}()

	w.ShowAndRun()
}
