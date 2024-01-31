package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	simpleApp := app.New()
	simpleApp.Settings().SetTheme(theme.LightTheme())
	simpleWindow := simpleApp.NewWindow("Logical Operators Demo")

	checkboxA := widget.NewCheck("A", nil)
	checkboxB := widget.NewCheck("B", nil)
	checkboxC := widget.NewCheck("C", nil)

	resultLabel := widget.NewLabel("Result: ")
	resultLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	updateResult := func() {
		a := checkboxA.Checked
		b := checkboxB.Checked
		c := checkboxC.Checked

		andResult := a && b && c
		orResult := a || b || c
		andOrResult := (a && b) || c

		resultLabel.SetText(fmt.Sprintf("Result: A && B && C = %v,\n A || B || C = %v,\n (A && B) || C = %v", andResult, orResult, andOrResult))
	}

	checkboxA.OnChanged = func(checked bool) {
		updateResult()
	}

	checkboxB.OnChanged = func(checked bool) {
		updateResult()
	}

	checkboxC.OnChanged = func(checked bool) {
		updateResult()
	}

	content := container.NewVBox(
		widget.NewLabel("Toggle the checkboxes to change values of A, B, and C:"),
		container.NewHBox(checkboxA, checkboxB, checkboxC),
		resultLabel,
	)

	simpleWindow.SetContent(content)
	simpleWindow.SetContent(container.New(layout.NewBorderLayout(nil, nil, nil, nil), content))
	simpleWindow.Resize(fyne.NewSize(400, 200))
	simpleWindow.SetFixedSize(true)
	simpleWindow.ShowAndRun()
}
