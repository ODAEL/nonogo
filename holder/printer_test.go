package holder

func generatePrinterTestNonogram() Nonogram {
	return Nonogram{generatePrinterTestField(), generatePrinterTestLeftBox(), generatePrinterTestTopBox()}
}

func generatePrinterTestField() Field {
	return Field{[][]Item{
		{{StateUnknown}, {StateBlack}},
		{{StateBlack}, {StateWhite}},
	}}
}

func generatePrinterTestLeftBox() Box {
	return Box{[][]int{
		{1},
		{1},
	}}
}

func generatePrinterTestTopBox() Box {
	return Box{[][]int{
		{1},
		{1},
	}}
}

func ExampleCmdNonogramPrinter_Print() {
	nonogram := generatePrinterTestNonogram()

	printer := CmdNonogramPrinter{nonogram}

	printer.Print()

	// Output:
	// Top box
	// ┏
	// ┃1
	// ┃1
	// ┗
	// Left box
	// ┏
	// ┃1
	// ┃1
	// ┗
	// Field
	// ┏━━━━┓
	// ┃  ██┃
	// ┃██><┃
	// ┗━━━━┛
}