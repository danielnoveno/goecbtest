package view

import (
	"fmt"
	"time"

	"goecbtest/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartUI(service *services.BreakerService) {
	a := app.New()
	w := a.NewWindow("ECB System")

	r, _ := fyne.LoadResourceFromPath("E:/Magang - Noveno/goecbtest_3/goecbtest/assets/images/cat.png")
    w.SetIcon(r)

	statusLabel := widget.NewLabelWithStyle(
		"Status: ✅ Normal",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Area log besar + scrollable
	logsView := widget.NewMultiLineEntry()
	// logsView.Disable()
	logsView.Wrapping = fyne.TextWrapWord
	logsView.SetText("System started...\n")

	scroll := container.NewVScroll(logsView)
	scroll.SetMinSize(fyne.NewSize(600, 150)) // lebih besar dari sebelumnya

	btnTestLeak := widget.NewButton("⚡ Simulate Leak", func() {
		service.Sensor.Write(true)
		statusLabel.SetText("Status: ⚠️ Leakage detected!")
		statusLabel.Refresh()
		service.Logs <- "[USER] Mouse clicked on Simulate Leak"
	})

	btnReset := widget.NewButton("🔁 Reset", func() {
		service.Sensor.Write(false)
		statusLabel.SetText("Status: ✅ Normal")
		statusLabel.Refresh()
		service.Logs <- "[USER] Mouse clicked on Reset"
	})

	// Area "listener" transparan untuk menangkap mouse dan keyboard event
	mouseArea := canvas.NewRectangle(nil)
	mouseArea.Resize(fyne.NewSize(600, 400))
	mouseArea.SetMinSize(fyne.NewSize(600, 400))
	mouseArea.Hide()

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		service.Logs <- "[USER] Keyboard key pressed: " + string(k.Name)
	})

	// w.Canvas().SetOnMouseDown(func(e *fyne.MouseEvent) {
	// 	service.Logs <- "[USER] Mouse clicked at (" + 
	// 		formatFloat(e.Position.X) + ", " + formatFloat(e.Position.Y) + ")"
	// })

	layout := container.NewVBox(
		widget.NewLabelWithStyle(
			"Electric Circuit Breaker (ECB) System",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		statusLabel,
		container.NewGridWithColumns(2, btnTestLeak, btnReset),
		widget.NewLabel("Event Log (latest first):"),
		scroll,
	)

	w.SetContent(container.NewStack(mouseArea, layout))
	w.Resize(fyne.NewSize(620, 520))

	// Goroutine buat update log terbaru di atas
	go func() {
		for msg := range service.Logs {
			newLog := time.Now().Format("15:04:05") + " " + msg + "\n"
			current := logsView.Text
			logsView.SetText(newLog + current)
			logsView.Refresh()
		}
	}()

	w.ShowAndRun()
}

func formatFloat(v float32) string {
	return fmt.Sprintf("%.0f", v)
}
