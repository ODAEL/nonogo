package solver

import (
	"github.com/stretchr/testify/assert"
	"nonogram-solver/holder"
	"testing"
)

func TestSolve(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		nonogram := holder.BuildEmptyNonogram(3, 3)

		want := [][]int{
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
			{holder.StateWhite, holder.StateWhite, holder.StateWhite},
		}

		Solve(&nonogram)

		field := nonogram.GetField()

		assert.Equal(t, want, field.GetFieldAsSlice())
	})

	t.Run("Fill", func(t *testing.T) {
		want := [][]int{
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
			{holder.StateBlack, holder.StateBlack, holder.StateBlack},
		}

		field := holder.CreateFieldBySlice(want)

		box := holder.Box{Numbers: [][]int{
			{3},
			{3},
			{3},
		}}

		nonogram := holder.Nonogram{Field: field, LeftBox: box, TopBox: box}

		Solve(&nonogram)

		field = nonogram.GetField()

		assert.Equal(t, want, field.GetFieldAsSlice())
	})
}
