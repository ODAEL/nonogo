package holder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateTestItems() [][]Item {
	return [][]Item{
		{{StateUnknown}, {StateBlack}, {StateWhite}},
		{{StateBlack}, {StateWhite}, {StateUnknown}},
	}
}

func TestField_GetItems(t *testing.T) {
	field := Field{generateTestItems()}

	assert.Equal(t, generateTestItems(), field.GetItems())
}

func TestField_GetHeight(t *testing.T) {
	field := Field{generateTestItems()}

	assert.Equal(t, 2, field.GetHeight())
}

func TestField_GetWidth(t *testing.T) {
	field := Field{generateTestItems()}

	assert.Equal(t, 3, field.GetWidth())
}

func TestField_GetFieldAsSlice(t *testing.T) {
	field := Field{generateTestItems()}

	want := [][]int{
		{StateUnknown, StateBlack, StateWhite},
		{StateBlack, StateWhite, StateUnknown},
	}

	assert.Equal(t, want, field.GetFieldAsSlice())
}

func TestCreateEmptyField(t *testing.T) {
	field := CreateEmptyField(3, 3)

	assert.Equal(t, 3, field.GetHeight())
	assert.Equal(t, 3, field.GetWidth())

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t, StateUnknown, field.Items[i][j].GetState())
		}
	}
}

func TestCreateBySlice(t *testing.T) {
	slice := [][]int{
		{StateUnknown, StateBlack, StateWhite},
		{StateBlack, StateWhite, StateUnknown},
	}

	field := CreateFieldBySlice(slice)

	assert.Equal(t, generateTestItems(), field.Items)
}
