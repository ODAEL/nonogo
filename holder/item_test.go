package holder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItem_GetState(t *testing.T) {

	t.Run("Unknown", func(t *testing.T) {
		item := Item{StateUnknown}

		assert.Equal(t, StateUnknown, item.GetState())
	})

	t.Run("Black", func(t *testing.T) {
		item := Item{StateBlack}

		assert.Equal(t, StateBlack, item.GetState())
	})

	t.Run("White", func(t *testing.T) {
		item := Item{StateWhite}

		assert.Equal(t, StateWhite, item.GetState())
	})
}

func TestItem_IsUnknown(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		item := Item{}
		assert.True(t, item.IsUnknown())

		item = Item{StateUnknown}
		assert.True(t, item.IsUnknown())
	})

	t.Run("False", func(t *testing.T) {
		item := Item{StateBlack}
		assert.False(t, item.IsUnknown())

		item = Item{StateWhite}
		assert.False(t, item.IsUnknown())
	})
}

func TestItem_IsBlack(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		item := Item{StateBlack}
		assert.True(t, item.IsBlack())
	})

	t.Run("False", func(t *testing.T) {
		item := Item{}
		assert.False(t, item.IsBlack())

		item = Item{StateUnknown}
		assert.False(t, item.IsBlack())

		item = Item{StateWhite}
		assert.False(t, item.IsBlack())
	})
}

func TestItem_IsWhite(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		item := Item{StateWhite}
		assert.True(t, item.IsWhite())
	})

	t.Run("False", func(t *testing.T) {
		item := Item{}
		assert.False(t, item.IsWhite())

		item = Item{StateUnknown}
		assert.False(t, item.IsWhite())

		item = Item{StateBlack}
		assert.False(t, item.IsWhite())
	})
}

func TestItem_PaintUnknown(t *testing.T) {
	item := Item{StateBlack}
	item.PaintUnknown()

	assert.True(t, item.IsUnknown())
}

func TestItem_PaintBlack(t *testing.T) {
	item := Item{StateWhite}
	item.PaintBlack()

	assert.True(t, item.IsBlack())
}

func TestItem_PaintWhite(t *testing.T) {
	item := Item{StateUnknown}
	item.PaintWhite()

	assert.True(t, item.IsWhite())
}
