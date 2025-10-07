// package main

// import (
// 	"fmt"
// 	"goecbtest/internal/models"
// 	"io/ioutil"
// 	"path/filepath"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )
// func main() {
// 	fmt.Println("Running ECB Test Application...")

// 	a := app.New()
// 	w := a.NewWindow("ECB Test Application")
// 	w.Resize(fyne.NewSize(400,300))

// 	r, _ := LoadResourceFromPath("E:/Magang - Noveno/goecbtest_3/goecbtest/assets/images/cat.png")
//     w.SetIcon(r)

// 	t := models.NewComprefg("TYPE","1234567890","ACTIVE")

// 	// t, _ := fyne.LoadResourceFromURLString("https://th.bing.com/th/id/OIP.FC65_sWDK45XrGsYRV5BRAHaFk?w=233&h=180&c=7&r=0&o=7&cb=12&pid=1.7&rm=3")
// 	// img := canvas.NewImageFromResource(t)

// 	// w.SetContent(widget.NewLabel("This is a test application for ECB."))
// 	// labelX := widget.NewLabel("This is a test application for ECB.")

// 	// btn := widget.NewButton("button name", func() {
// 	// 	fmt.Println("Button clicked!")
// 	// })

// 	// check := widget.NewCheck("check box", func (b bool)  {
// 	// 	fmt.Printf("Check box value: %t\n", b)
// 	// })

// 	// url, _ := url.Parse("http://danielnoveno.com")
// 	// hyperlink := widget.NewHyperlink("Visit Daniel's Website", url)

// 	// colorX := color.NRGBA{R:0, G:255, B:0, A:255}
// 	// textX := canvas.NewText("Here my text!", colorX)

// 	// imgx := canvas.NewImageFromFile("../../public/images/test-361512_640.jpg")
// 	// imgx.FillMode = canvas.ImageFillContain

// 	// iconX := widget.NewIcon(theme.CancelIcon())

// 	// w.SetContent(labelX)
// 	// w.SetContent(img)
// 	// w.SetContent(btn)
// 	// w.SetContent(check)
// 	// w.SetContent(hyperlink)
// 	// w.SetContent(textX)
// 	// w.SetContent(imgx)
// 	// w.SetContent(iconX)

// 	newComfreCtypeTxt := widget.NewEntry()
// 	newComfreCtypeTxt.SetPlaceHolder("Enter Comprefg Ctype")
// 	addBtnCtype := widget.NewButton("Add", func() { fmt.Println("Add was clicked!") })
// 	addBtnCtype.Disable()

// 	newComfreCtypeTxt.OnChanged = func (s string) {
// 		addBtnCtype.Disable()
// 		if len(s) >= 3 {
// 			addBtnCtype.Enable()
// 		}
// 	}

// 	w.SetContent(
//         container.NewBorder(
//             nil, // TOP of the container

//             // this will be a the BOTTOM of the container
// 			container.NewBorder(
// 				nil,
// 				nil,
// 				nil,
// 				addBtnCtype,
// 				newComfreCtypeTxt,
// 			),

//             nil, // Right
//             nil, // Left

//             // the rest will take all the rest of the space
//             container.NewCenter(
//                 widget.NewLabel(t.String()),
//             ),
//         ),
//     )

// 	w.ShowAndRun()
// }
// type StaticResource struct {
//     StaticName    string
//     StaticContent []byte
// }

// func (r *StaticResource) Name() string {
//     return r.StaticName
// }

// func (r *StaticResource) Content() []byte {
//     return r.StaticContent
// }

// func NewStaticResource(name string, content []byte) *StaticResource {
//     return &StaticResource{
//         StaticName:    name,
//         StaticContent: content,
//     }
// }
// func LoadResourceFromPath(path string) (fyne.Resource, error) {
//     bytes, err := ioutil.ReadFile(filepath.Clean(path))
//     if err != nil {
//         return nil, err
//     }
//     name := filepath.Base(path)
//     return NewStaticResource(name, bytes), nil
// }

package main

import (
	"runtime"

	"goecbtest/internal/gpio"
	"goecbtest/internal/services"
	"goecbtest/internal/view"
)

func main() {
	var sensorPin gpio.Pin
	var relayPin gpio.Pin

	if runtime.GOOS == "windows" {
		sensorPin = gpio.NewMockPin("SENSOR")
		relayPin = gpio.NewMockPin("RELAY")
	} 
	// else {
	// 	sensorPin = gpio.NewRealPin("GPIO17")
	// 	relayPin = gpio.NewRealPin("GPIO27")
	// }

	breaker := services.NewBreaker(sensorPin, relayPin)

	go breaker.Monitor() // jalan di background

	view.StartUI(breaker)
}
