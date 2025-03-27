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
	simpleApp.Settings().Theme().Font(fyne.TextStyle{Bold: true, Monospace: true})
	simpleWindow := simpleApp.NewWindow("Logical Operators Demo")

	checkboxA := widget.NewCheck("A", nil)
	checkboxB := widget.NewCheck("B", nil)
	checkboxC := widget.NewCheck("C", nil)
	checkboxD := widget.NewCheck("D", nil)

	result := widget.NewLabel("Result from Logical Operators:")

	resultAndLabel := widget.NewLabel("1. AND (A && B && C && D) = false")
	resultOrLabel := widget.NewLabel("2. OR (A || B || C || D) = false")
	resultAndOrLabel := widget.NewLabel("3. AND OR ((A && B) || (C && D)) = false")
	resultAndOrOrLabel := widget.NewLabel("4. AND OR OR ((A && B) || C || D) = false")

	notAResultLabel := widget.NewLabel("5. NOT A (!A) = false")
	xorResultLabel := widget.NewLabel("6. A XOR B = false")
	nandResultLabel := widget.NewLabel("7. NAND (A && B) = false")
	norResultLabel := widget.NewLabel("8. NOR (A || B) = true")

	updateResult := func() {
		a := checkboxA.Checked
		b := checkboxB.Checked
		c := checkboxC.Checked
		d := checkboxD.Checked

		andResult := a && b && c && d
		orResult := a || b || c || d
		andOrResult := (a && b) || (c && d)
		andOrOrResult := (a && b) || (c && d)

		notAResult := !a
		xorResult := a != b
		nandResult := !(a && b)
		norResult := !(a || b)

		resultAndLabel.SetText(fmt.Sprintf("1. AND (A && B && C && D) = %v", andResult))
		resultOrLabel.SetText(fmt.Sprintf("2. OR (A || B || C || D) = %v", orResult))
		resultAndOrLabel.SetText(fmt.Sprintf("3. AND OR ((A && B) || (C && D)) = %v", andOrResult))
		resultAndOrOrLabel.SetText(fmt.Sprintf("4. AND OR OR ((A && B) || C || D) = %v", andOrOrResult))

		notAResultLabel.SetText(fmt.Sprintf("5. NOT A (!A) = %v", notAResult))
		xorResultLabel.SetText(fmt.Sprintf("6. A XOR B = %v", xorResult))
		nandResultLabel.SetText(fmt.Sprintf("7. NAND (A && B) = %v", nandResult))
		norResultLabel.SetText(fmt.Sprintf("8. NOR (A || B) = %v", norResult))
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

	checkboxD.OnChanged = func(checked bool) {
		updateResult()
	}

	content := container.NewVBox(
		widget.NewLabel("Toggle the checkboxes to change values of A, B, C and D:"),
		container.NewHBox(checkboxA, checkboxB, checkboxC, checkboxD),
		result,
		resultAndLabel,
		resultOrLabel,
		resultAndOrLabel,
		resultAndOrOrLabel,
		notAResultLabel,
		xorResultLabel,
		nandResultLabel,
		norResultLabel,
	)

	simpleWindow.SetContent(content)
	simpleWindow.SetContent(container.New(layout.NewBorderLayout(nil, nil, nil, nil), content))
	simpleWindow.Resize(fyne.NewSize(400, 400))
	simpleWindow.SetFixedSize(true)
	simpleWindow.ShowAndRun()
}
