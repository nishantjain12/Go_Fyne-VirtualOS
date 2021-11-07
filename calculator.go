package main

import (
	"math"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showCalculator() {
	myWindow := myApp.NewWindow("Calculator")

	var message string = ""
	var lastChar string = ""
	var val1 string = ""
	var val2 string = ""
	var symbol string = ""

	display := widget.NewEntry()
	display.TextStyle = fyne.TextStyle{Bold: true}
	display.SetPlaceHolder("")
	display.Disable()

	input := widget.NewEntry()
	input.TextStyle = fyne.TextStyle{Bold: true}
	input.SetPlaceHolder("0")
	input.Disable()

	modules := widget.NewButton("%", func() {
		if symbol != "" {
			first, _ := strconv.ParseFloat(val1, 32)
			second, _ := strconv.ParseFloat(val2, 32)

			second = (first * (second / 100))
			val2 = strconv.FormatFloat(second, 'f', -1, 32)
			input.SetText(val2)
			message = val1 + " " + symbol + " " + val2
			display.SetText(message)
		} else {
			val1 = ""
			val2 = ""
			input.SetText(val1)
			message = val1 + " " + symbol + " " + val2
			display.SetText(message)
		}
	})

	clear := widget.NewButton("C", func() {
		message = ""
		lastChar = ""
		val1 = ""
		symbol = ""
		val2 = ""
		input.SetText(lastChar)
		display.SetText(message)
	})

	back := widget.NewButton("<", func() {
		if !strings.Contains(message, "=") {
			if symbol == "" {
				if val1 == "" {
					val1 = ""
				} else {
					val1 = val1[0 : len(val1)-1]
				}
				input.SetText(val1)
				message = val1 + " " + symbol + " " + val2
				display.SetText(message)
			} else {
				if val2 == "" {
					val2 = ""
				} else {
					val2 = val2[0 : len(val2)-1]
				}
				input.SetText(val2)
				message = val1 + " " + symbol + " " + val2
				display.SetText(message)
			}
		} else {
			display.SetText("")
		}
	})

	reciprocal := widget.NewButton("1/x", func() {
		if !strings.Contains(message, "=") {
			if symbol == "" {
				if val1 != "" {
					v1, _ := strconv.ParseFloat(val1, 32)
					v1 = 1 / v1
					val1 = "1/(" + val1 + ")"
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val1 = strconv.FormatFloat(v1, 'f', -1, 32)
					input.SetText(val1)
				}
			} else {
				if val2 != "" {
					v2, _ := strconv.ParseFloat(val2, 32)
					v2 = 1 / v2
					val2 = "1/(" + val2 + ")"
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val2 = strconv.FormatFloat(v2, 'f', -1, 32)
					input.SetText(val2)
				}
			}
		} else {
			if val1 != "" {
				firstValue, _ := strconv.ParseFloat(val1, 32)
				firstValue = 1 / firstValue
				val1 = "1/(" + val1 + ")"
				val2 = ""
				symbol = ""
				message = val1 + " " + symbol + " " + val2
				display.SetText(message)
				val1 = strconv.FormatFloat(firstValue, 'f', -1, 32)
				input.SetText(val1)
			}
		}
	})
	square := widget.NewButton("x²", func() {
		if !strings.Contains(message, "=") {
			if symbol == "" {
				if val1 != "" {
					v1, _ := strconv.ParseFloat(val1, 32)
					v1 = v1 * v1
					val1 = val1 + "²"
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val1 = strconv.FormatFloat(v1, 'f', -1, 32)
					input.SetText(val1)
				}
			} else {
				if val2 != "" {
					v2, _ := strconv.ParseFloat(val2, 32)
					v2 = v2 * v2
					val2 = val2 + "²"
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val2 = strconv.FormatFloat(v2, 'f', -1, 32)
					input.SetText(val2)
				}
			}
		} else {
			if val1 != "" {
				firstValue, _ := strconv.ParseFloat(val1, 32)
				firstValue = firstValue * firstValue
				val1 = val1 + "²"
				val2 = ""
				symbol = ""
				message = val1 + " " + symbol + " " + val2
				display.SetText(message)
				val1 = strconv.FormatFloat(firstValue, 'f', -1, 32)
				input.SetText(val1)
			}
		}
	})
	squareRoot := widget.NewButton("sqrt(x)", func() {
		if !strings.Contains(message, "=") {
			if symbol == "" {
				if val1 != "" {
					v1, _ := strconv.ParseFloat(val1, 32)
					v1 = math.Sqrt(v1)
					val1 = "√" + val1
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val1 = strconv.FormatFloat(v1, 'f', -1, 32)
					input.SetText(val1)
				}
			} else {
				if val2 != "" {
					v2, _ := strconv.ParseFloat(val2, 32)
					v2 = math.Sqrt(v2)
					val2 = "√" + val2
					message = val1 + " " + symbol + " " + val2
					display.SetText(message)
					val2 = strconv.FormatFloat(v2, 'f', -1, 32)
					input.SetText(val2)
				}
			}
		} else {
			if val1 != "" {
				firstValue, _ := strconv.ParseFloat(val1, 32)
				firstValue = math.Sqrt(firstValue)
				val1 = "sqrt(" + val1 + ")"
				val2 = ""
				symbol = ""
				message = val1 + " " + symbol + " " + val2
				display.SetText(message)
				val1 = strconv.FormatFloat(firstValue, 'f', -1, 32)
				input.SetText(val1)
			}
		}
	})
	divide := widget.NewButton("/", func() {
		val2 = ""
		if val1 == "" {
			val1 = "0"
			input.SetText(val1)
		} else if val1[len(val1)-1:] == "." {
			val1 = val1[0 : len(val1)-1]
			input.SetText(val1)
		}
		symbol = "÷"
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	seven := widget.NewButton("7", func() {
		if symbol == "" {
			val1 = val1 + "7"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "7"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	eight := widget.NewButton("8", func() {
		if symbol == "" {
			val1 = val1 + "8"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "8"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	nine := widget.NewButton("9", func() {
		if symbol == "" {
			val1 = val1 + "9"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "9"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	multiply := widget.NewButton("*", func() {
		val2 = ""
		if val1 == "" {
			val1 = "0"
			input.SetText(val1)
		} else if val1[len(val1)-1:] == "." {
			val1 = val1[0 : len(val1)-1]
			input.SetText(val1)
		}
		symbol = "*"
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	four := widget.NewButton("4", func() {
		if symbol == "" {
			val1 = val1 + "4"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "4"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	five := widget.NewButton("5", func() {
		if symbol == "" {
			val1 = val1 + "5"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "5"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	six := widget.NewButton("6", func() {
		if symbol == "" {
			val1 = val1 + "6"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "6"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	minus := widget.NewButton("-", func() {
		val2 = ""
		if val1 == "" {
			val1 = "0"
			input.SetText(val1)
		} else if val1[len(val1)-1:] == "." {
			val1 = val1[0 : len(val1)-1]
			input.SetText(val1)
		}
		symbol = "-"
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	one := widget.NewButton("1", func() {
		if symbol == "" {
			val1 = val1 + "1"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "1"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	two := widget.NewButton("2", func() {
		if symbol == "" {
			val1 = val1 + "2"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "2"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	three := widget.NewButton("3", func() {
		if symbol == "" {
			val1 = val1 + "3"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "3"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})
	add := widget.NewButton("+", func() {
		if val1 == "" {
			val1 = "0"
			input.SetText(val1)
		} else if val1[len(val1)-1:] == "." {
			val1 = val1[0 : len(val1)-1]
			input.SetText(val1)
		}
		symbol = "+"
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	mod := widget.NewButton("+/-", func() {
		if !strings.Contains(message, "=") {
			if symbol == "" {
				if val1 != "" {
					firstValue, _ := strconv.ParseFloat(val1, 32)
					if firstValue > 0 {
						val1 = "-" + val1
						input.SetText(val1)
					} else {
						val1 = val1[1:]
						input.SetText(val1)
					}
				}
			} else {
				if val2 != "" {
					secondValue, _ := strconv.ParseFloat(val2, 32)
					if secondValue > 0 {
						val2 = "-" + val2
						input.SetText(val2)
					} else {
						val2 = val2[1:]
						input.SetText(val2)
					}
				}
			}
		} else {
			if val1 != "" {
				firstValue, _ := strconv.ParseFloat(val1, 32)
				if firstValue > 0 {
					val1 = "-" + val1
					input.SetText(val1)
				} else {
					val1 = val1[1:]
					input.SetText(val1)
				}
			}
			val2 = ""
			symbol = ""
			message = val1 + " " + symbol + " " + val2
			display.SetText(message)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	zero := widget.NewButton("0", func() {
		if symbol == "" {
			val1 = val1 + "0"
			lastChar = val1
			input.SetText(lastChar)
		} else {
			val2 = val2 + "0"
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	dot := widget.NewButton(".", func() {
		if symbol == "" {
			if !strings.Contains(val1, ".") {
				if val1 == "" {
					val1 = "0."
				} else {
					val1 = val1 + "."
				}
			}
			lastChar = val1
			input.SetText(lastChar)
		} else {
			if !strings.Contains(val2, ".") {
				if val2 == "" {
					val2 = "0."
				} else {
					val2 = val2 + "."
				}
			}
			lastChar = val2
			input.SetText(lastChar)
		}
		message = val1 + " " + symbol + " " + val2
		display.SetText(message)
	})

	equal := widget.NewButton("=", func() {
		v1, _ := strconv.ParseFloat(val1, 32)
		if symbol != "" && val2 == "" {
			val2 = val1
		}
		v2, _ := strconv.ParseFloat(val2, 32)
		if symbol == "+" {
			v1 = v1 + v2
			input.SetText(strconv.FormatFloat(v1, 'f', -1, 32))
			message = val1 + " " + symbol + " " + val2 + " " + "="
			display.SetText(message)
			val1 = strconv.FormatFloat(v1, 'f', -1, 32)
		}
		if symbol == "*" {
			v1 = v1 * v2
			input.SetText(strconv.FormatFloat(v1, 'f', -1, 32))
			message = val1 + " " + symbol + " " + val2 + " " + "="
			display.SetText(message)
			val1 = strconv.FormatFloat(v1, 'f', -1, 32)
		}
		if symbol == "-" {
			v1 = v1 - v2
			input.SetText(strconv.FormatFloat(v1, 'f', -1, 32))
			message = val1 + " " + symbol + " " + val2 + " " + "="
			display.SetText(message)
			val1 = strconv.FormatFloat(v1, 'f', -1, 32)
		}
		if symbol == "÷" {
			v1 = v1 / v2
			input.SetText(strconv.FormatFloat(v1, 'f', -1, 32))
			message = val1 + " " + symbol + " " + val2 + " " + "="
			display.SetText(message)
			val1 = strconv.FormatFloat(v1, 'f', -1, 32)
		}
	})

	myWindow.SetContent(
		container.New(
			layout.NewGridLayoutWithRows(1),
			container.New(
				layout.NewGridLayout(1),
				container.New(layout.NewGridLayout(1), display),
				container.New(layout.NewGridLayout(1), input),
				container.New(layout.NewGridLayout(3), modules, clear, back),
				container.New(layout.NewGridLayout(4), reciprocal, square, squareRoot, divide),
				container.New(layout.NewGridLayout(4), seven, eight, nine, multiply),
				container.New(layout.NewGridLayout(4), four, five, six, minus),
				container.New(layout.NewGridLayout(4), one, two, three, add),
				container.New(layout.NewGridLayout(4), mod, zero, dot, equal),
			),
		),
	)
	myWindow.Resize(fyne.NewSize(320, 470))
	myWindow.CenterOnScreen()
	myWindow.Show()
}
