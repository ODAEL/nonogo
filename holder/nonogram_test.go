package holder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateTestNonogram() Nonogram {
	return Nonogram{generateTestField(), generateTestLeftBox(), generateTestTopBox()}
}

func generateTestField() Field {
	return Field{[][]Item{
		{{StateUnknown}, {StateBlack}, {StateBlack}},
		{{StateBlack}, {StateWhite}, {StateWhite}},
	}}
}

func generateTestLeftBox() Box {
	return Box{[][]int{
		{1},
		{2},
	}}
}

func generateTestTopBox() Box {
	return Box{[][]int{
		{2},
		{1},
		{1},
	}}
}

func TestNonogram_GetHeight(t *testing.T) {
	nonogram := generateTestNonogram()

	assert.Equal(t, 2, nonogram.GetHeight())
}

func TestNonogram_GetWidth(t *testing.T) {
	nonogram := generateTestNonogram()

	assert.Equal(t, 3, nonogram.GetWidth())
}

func TestBuildEmptyNonogram(t *testing.T) {
	nonogram := BuildEmptyNonogram(2, 4)

	assert.Equal(t, 2, nonogram.GetWidth())
	assert.Equal(t, 4, nonogram.GetHeight())
	assert.Equal(t, Box{make([][]int, 4)}, nonogram.LeftBox)
	assert.Equal(t, Box{make([][]int, 2)}, nonogram.TopBox)
	assert.Equal(t, CreateEmptyField(2, 4), nonogram.Field)
}

func TestBuildNonogram(t *testing.T) {
	leftBox := generateTestLeftBox();
	topBox := generateTestTopBox();
	nonogram := BuildNonogram(3, 2, leftBox.Numbers, topBox.Numbers)

	assert.Equal(t, 3, nonogram.GetWidth())
	assert.Equal(t, 2, nonogram.GetHeight())
	assert.Equal(t, generateTestLeftBox(), nonogram.LeftBox)
	assert.Equal(t, generateTestTopBox(), nonogram.TopBox)
	assert.Equal(t, CreateEmptyField(3, 2), nonogram.Field)
}