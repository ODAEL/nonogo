package solver

import (
	"github.com/stretchr/testify/assert"
	"nonogram-solver/holder"
	"testing"
)

type solverTestCase struct {
	want            [][]int
	topBox, leftBox holder.Box
	width, height   int
}

var cases = map[string]solverTestCase{
	"Empty": solverTestCase{
		want: [][]int{
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
		},
		topBox: holder.Box{Numbers: [][]int{
			{},
			{},
			{},
		}},
		leftBox: holder.Box{Numbers: [][]int{
			{},
			{},
			{},
		}},
		width:  3,
		height: 3},

	"Fill": solverTestCase{
		want: [][]int{
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
		},
		topBox: holder.Box{Numbers: [][]int{
			{3},
			{3},
			{3},
		}},
		leftBox: holder.Box{Numbers: [][]int{
			{3},
			{3},
			{3},
		}},
		width:  3,
		height: 3},

	"Plus": solverTestCase{
		want: [][]int{
			{holder.StateWhite, holder.StateBlack, holder.StateWhite},
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
			{holder.StateWhite, holder.StateBlack, holder.StateWhite},
		},
		topBox: holder.Box{Numbers: [][]int{
			{1},
			{3},
			{1},
		}},
		leftBox: holder.Box{Numbers: [][]int{
			{1},
			{3},
			{1},
		}},
		width:  3,
		height: 3},

	"Hard": solverTestCase{
		want: [][]int{
			{holder.StateWhite, holder.StateBlack, holder.StateWhite},
			{holder.StateBlack, holder.StateWhite, holder.StateBlack},
		},
		topBox: holder.Box{Numbers: [][]int{
			{1},
			{1},
			{1},
		}},
		leftBox: holder.Box{Numbers: [][]int{
			{1},
			{1, 1},
		}},
		width:  3,
		height: 2},

	"Rabbit": solverTestCase{
		want: [][]int{
			{-1, -1, -1, -1, 01, 01, -1, 01, 01, -1, -1, -1, -1},
			{-1, -1, -1, 01, 01, 01, 01, 01, 01, -1, -1, -1, -1},
			{-1, -1, -1, 01, 01, 01, 01, 01, -1, -1, -1, -1, -1},
			{-1, 01, 01, 01, 01, 01, -1, -1, -1, -1, -1, -1, -1},
			{01, 01, -1, 01, 01, -1, -1, -1, -1, -1, -1, -1, -1},
			{01, 01, 01, 01, 01, 01, 01, 01, -1, -1, -1, -1, -1},
			{01, 01, 01, 01, 01, 01, 01, 01, 01, 01, -1, -1, -1},
			{-1, -1, 01, 01, 01, 01, 01, 01, 01, 01, 01, -1, -1},
			{-1, -1, 01, 01, 01, 01, 01, 01, 01, 01, 01, 01, 01},
			{-1, -1, -1, 01, 01, 01, 01, 01, 01, 01, 01, 01, 01},
			{-1, -1, -1, -1, 01, 01, 01, 01, 01, 01, 01, 01, 01},
			{-1, -1, -1, -1, 01, 01, 01, 01, 01, 01, 01, 01, -1},
			{-1, -1, -1, 01, 01, 01, -1, 01, 01, 01, 01, -1, -1},
		},
		topBox: holder.Box{Numbers: [][]int{
			{3},
			{4},
			{1, 4},
			{9, 1},
			{13},
			{4, 8},
			{2, 7},
			{3, 8},
			{2, 7},
			{7},
			{6},
			{4},
			{3},
		}},
		leftBox: holder.Box{Numbers: [][]int{
			{2, 2},
			{6},
			{5},
			{5},
			{2, 2},
			{8},
			{10},
			{9},
			{11},
			{10},
			{9},
			{8},
			{3, 4},
		}},
		width:  13,
		height: 13},
}

func TestSolve(t *testing.T) {

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			want := testCase.want
			topBox := testCase.topBox
			leftBox := testCase.leftBox

			nonogram := holder.BuildEmptyNonogram(testCase.width, testCase.height)
			nonogram.TopBox = topBox
			nonogram.LeftBox = leftBox

			Solve(&nonogram)

			field := nonogram.Field

			assert.Equal(t, want, field.GetFieldAsSlice())
		})
	}
}

func ExampleSolve() {

		testCase := cases["Rabbit"]

		topBox := testCase.topBox
		leftBox := testCase.leftBox

		nonogram := holder.BuildEmptyNonogram(testCase.width, testCase.height)
		nonogram.TopBox = topBox
		nonogram.LeftBox = leftBox

		Solve(&nonogram)

		printer := holder.CmdNonogramPrinter{Nonogram: nonogram}
		printer.PrintAll()

		// Output:
		// Top box
		// ┏
		// ┃3
		// ┃4
		// ┃1 4
		// ┃9 1
		// ┃13
		// ┃4 8
		// ┃2 7
		// ┃3 8
		// ┃2 7
		// ┃7
		// ┃6
		// ┃4
		// ┃3
		// ┗
		// Left box
		// ┏
		// ┃2 2
		// ┃6
		// ┃5
		// ┃5
		// ┃2 2
		// ┃8
		// ┃10
		// ┃9
		// ┃11
		// ┃10
		// ┃9
		// ┃8
		// ┃3 4
		// ┗
		// Field
		// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓
		// ┃><><><><████><████><><><><┃
		// ┃><><><████████████><><><><┃
		// ┃><><><██████████><><><><><┃
		// ┃><██████████><><><><><><><┃
		// ┃████><████><><><><><><><><┃
		// ┃████████████████><><><><><┃
		// ┃████████████████████><><><┃
		// ┃><><██████████████████><><┃
		// ┃><><██████████████████████┃
		// ┃><><><████████████████████┃
		// ┃><><><><██████████████████┃
		// ┃><><><><████████████████><┃
		// ┃><><><██████><████████><><┃
		// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛
}
