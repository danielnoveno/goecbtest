package main

// import (
// 	"fmt"
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	var x uint8 = 0

// 	a := app.New()
// 	w := a.NewWindow("Traffic Lights Project")
// 	w.Resize(fyne.NewSize(400, 400))
// 	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

// 	rect1 := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
// 	rect1.Resize(fyne.NewSize(50, 50))
// 	rect2 := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
// 	rect2.Resize(fyne.NewSize(58, 58))
// 	rect3:= canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
// 	rect3.Resize(fyne.NewSize(50, 50))

// 	btnR := widget.NewButton("Red", func() {
// 		fmt.Printf("%t", true)
// 		//rect.Move(fyne.NewPos(incrementMove(x, y)))
// 		Reset(rect1, rect2, rect3)
// 		rect1.FillColor = color.NRGBA{R: incrementMove(x), G: 0, B: 0, A: 255}
// 		rect1.Refresh()
// 	})

// 	btnG := widget.NewButton("Green", func() {
// 		//rect.Move(fyne. NewPos(incrementMove(x, y)))
// 		Reset(rect1, rect2, rect3)
// 		rect2.FillColor = color.NRGBA{R: 0, G: incrementMove(x), B: 0, A: 255}
// 		rect2.Refresh()
// 	})

// 	btnB := widget.NewButton("Blue", func() {
// 		//rect.Move(fyne. NewPos(incrementMove(x, y)))
// 		Reset(rect1, rect2, rect3)
// 		rect3.FillColor = color.NRGBA{R: 0, G: 0, B: incrementMove(x), A: 255}
// 		rect3.Refresh()
// 	})

// 	btnReset := widget.NewButton("Reset", func(){
// 		rect1.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
// 		rect2.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
// 		rect3.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

// 		rect1.Refresh()
// 		rect2.Refresh()
// 		rect3.Refresh()
// 	})

// 	w.SetContent(
// 		container.NewHSplit(
// 			container.NewGridWithRows(
// 				7,
// 				layout.NewSpacer(),
// 				rect1,
// 				layout.NewSpacer(),
// 				rect2,
// 				layout.NewSpacer(),
// 				rect3,
// 				layout.NewSpacer(),
// 			),
// 			container.NewGridWithColumns(
// 				9,
// 				layout.NewSpacer(),
// 				btnR,
// 				btnG,
// 				btnB,
// 				btnReset,
// 				layout.NewSpacer(),
// 			),
// 		),
// 	)
// 	w.ShowAndRun()
// }

// func incrementMove(x uint8) (a uint8) {
// 	if x == 255 {
// 		x = 0
// 	} else if x == 0 {
// 		x = 255
// 	}
// 	return x
// }

// func Reset(rect1, rect2, rect3 *canvas.Circle) {
// 	rect1.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
// 	rect2.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
// 	rect3.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

// 	rect1.Refresh()
// 	rect2.Refresh()
// 	rect3.Refresh()
// }